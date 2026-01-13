# Dependency Audit Report: argocd-lovely-plugin

**Date:** January 2026
**Auditor:** Automated Analysis

## Executive Summary

The project has a well-maintained dependency structure with Renovate automation in place. However, there are several areas of concern regarding security, outdated packages, and potential bloat.

---

## 1. Security Concerns

### 1.1 High Priority

| Dependency | Version | Issue | Recommendation |
|------------|---------|-------|----------------|
| `github.com/gogo/protobuf` | v1.3.2 | This library is in maintenance mode and has had historical CVEs (CVE-2021-3121). While v1.3.2 includes fixes, it's recommended to migrate away. | Consider migrating to `google.golang.org/protobuf` if possible (transitive dep from k8s.io) |
| `k8s.io/client-go` | v11.0.0+incompatible | Pinned to a very old version via replace directive | This is a known Kubernetes ecosystem issue; the replace is intentional to avoid pulling newer versions |

### 1.2 Medium Priority

| Dependency | Version | Issue |
|------------|---------|-------|
| `jaytaylor.com/html2text` | 2023-03-21 | Unmaintained (last commit 2023). Has dependency issues requiring a `replace` directive for tablewriter. |
| `github.com/ssor/bom` | 2017 | Very old dependency, transitive from html2text |

### 1.3 Positive Findings

- The commented-out CVE replace directives in go.mod are no longer needed as the current versions include fixes
- `gopkg.in/yaml.v3 v3.0.1` is patched for CVE-2022-28948
- The project uses `#nosec` annotations appropriately for intentional file operations

---

## 2. Outdated Packages

### 2.1 Direct Dependencies

| Package | Current | Notes |
|---------|---------|-------|
| `github.com/evanphx/json-patch` | v5.9.11+incompatible | Using +incompatible tag; consider `github.com/evanphx/json-patch/v5` module path |
| `github.com/gomarkdown/markdown` | 2025-08-10 (pseudo-version) | Using unreleased commit; acceptable but watch for tagged releases |
| `github.com/hexops/gotextdiff` | v1.0.3 | Last release 2022; project appears dormant |

### 2.2 Dockerfile Tool Versions

| Tool | Current Version | Status |
|------|----------------|--------|
| Go | 1.25.5 | Current |
| Alpine | 3.22.2 | Current |
| yq | v4.49.2 | Current |
| kustomize | 5.8.0 | Current |
| Helm | v3.18.6 | Marked `donotrenovatefornow` - intentionally held |
| helmfile | v1.2.2 | Current |

---

## 3. Unnecessary Bloat Analysis

### 3.1 Heavy Transitive Dependencies

The project pulls in significant Kubernetes ecosystem dependencies despite limited actual usage:

**Kustomize ecosystem (`sigs.k8s.io/kustomize/...`):**
- `api v0.21.0` - Only uses `types.Kustomization` struct
- `kyaml v0.21.0` - Only uses `merge2.MergeStrings()` and `yaml.MergeOptions`

This brings in ~15+ transitive dependencies including:
- `sigs.k8s.io/structured-merge-diff/v4` and `/v6`
- `go.yaml.in/yaml/v2` and `/v3`
- Various go-openapi packages

**k8s.io/apimachinery:**
- Only uses `metav1.ObjectMeta` for plugin.yaml generation
- Brings in protobuf, json-iterator, and other heavy deps

### 3.2 Replace Directive Bloat

The go.mod has 22+ replace directives for Kubernetes packages that aren't directly used. These exist to prevent dependency resolution pulling conflicting versions. While necessary for compatibility, they add maintenance burden.

### 3.3 Dual YAML Libraries

The project uses multiple YAML libraries:
- `gopkg.in/yaml.v3` - Direct parsing
- `sigs.k8s.io/yaml` - JSON/YAML conversion
- `go.yaml.in/yaml/v2` and `/v3` - Transitive from kustomize
- `sigs.k8s.io/kustomize/kyaml/yaml` - Merge operations

---

## 4. Recommendations

### 4.1 High Priority Actions

1. **Replace `jaytaylor.com/html2text`** with an actively maintained alternative:
   - Consider `github.com/JohannesKaufmann/html-to-markdown` (active, well-maintained)
   - This would eliminate the tablewriter replace hack and the unmaintained `ssor/bom` dependency

2. **Update json-patch import path**:
   ```go
   // From:
   "github.com/evanphx/json-patch"
   // To:
   "github.com/evanphx/json-patch/v5"
   ```
   This properly uses Go modules semantics.

### 4.2 Medium Priority Actions

3. **Consider lightweight alternatives to kustomize libraries**:
   - If only using `merge2.MergeStrings()`, evaluate if a simpler YAML merge library would suffice
   - The `types.Kustomization` struct could potentially be defined locally (it's a simple struct)

4. **Consolidate YAML libraries**:
   - Standardize on one YAML library where possible
   - `sigs.k8s.io/yaml` is well-maintained and provides both parsing and JSON conversion

5. **Remove commented code in go.mod**:
   - Lines 69-80 contain commented replace directives for already-fixed CVEs
   - Clean these up for clarity

### 4.3 Low Priority / Watch Items

6. **Monitor `gotextdiff`**: Consider switching to an alternative if needed (only used in tests)

7. **Helm version hold**: Document why `HELM_VERSION` has `donotrenovatefornow` - appears intentional but reason unclear

8. **Clean up k8s.io replaces**: Periodically test if all 22 replace directives are still needed as upstream stabilizes

---

## 5. Dependency Graph Summary

```
Direct Dependencies: 12
├── Production: 8
│   ├── json-patch (YAML/JSON patching)
│   ├── splitter (string parsing)
│   ├── yaml.v3 (YAML parsing)
│   ├── kustomize/api (types only)
│   ├── kustomize/kyaml (merge2 only)
│   ├── sigs.k8s.io/yaml (conversion)
│   ├── k8s.io/apimachinery (metav1 only)
│   └── gomarkdown + html2text (doc generation)
└── Test-only: 4
    ├── testify (assertions)
    ├── gotextdiff (diff display)
    ├── otiai10/copy (file operations)
    └── stretchr/testify

Transitive Dependencies: ~50+
└── Majority from k8s.io/kustomize ecosystem
```

---

## 6. Security Scan Recommendations

Add continuous vulnerability monitoring to CI pipeline:

```yaml
# Add to .github/workflows/
- name: Run govulncheck
  run: |
    go install golang.org/x/vuln/cmd/govulncheck@latest
    govulncheck ./...
```

---

## 7. Actionable Changes Summary

| Priority | Action | Effort | Impact |
|----------|--------|--------|--------|
| High | Replace html2text library | Medium | Removes unmaintained dep + replace hack |
| High | Update json-patch to v5 module path | Low | Proper Go modules usage |
| Medium | Clean commented CVE replaces in go.mod | Low | Code cleanliness |
| Medium | Add govulncheck to CI | Low | Ongoing security monitoring |
| Low | Evaluate kustomize dependency reduction | High | Reduced binary size |
| Low | Document Helm version hold reason | Low | Maintainability |
