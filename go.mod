module github.com/crumbhole/argocd-lovely-plugin

go 1.24.0

toolchain go1.24.3

require (
	github.com/evanphx/json-patch v5.9.11+incompatible
	github.com/go-andiamo/splitter v1.2.5
	github.com/gomarkdown/markdown v0.0.0-20250311123330-531bef5e742b
	github.com/hexops/gotextdiff v1.0.3
	github.com/otiai10/copy v1.14.1
	github.com/stretchr/testify v1.10.0
	gopkg.in/yaml.v3 v3.0.1
	jaytaylor.com/html2text v0.0.0-20230321000545-74c2419ad056
	k8s.io/apimachinery v0.33.1
	sigs.k8s.io/kustomize/api v0.19.0
	sigs.k8s.io/kustomize/kyaml v0.19.0
	sigs.k8s.io/yaml v1.4.0
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fxamacker/cbor/v2 v2.8.0 // indirect
	github.com/go-errors/errors v1.5.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-openapi/jsonpointer v0.21.1 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/gnostic-models v0.6.9 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.9.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/olekukonko/errors v1.1.0 // indirect
	github.com/olekukonko/ll v0.0.8 // indirect
	github.com/olekukonko/tablewriter v1.0.6 // indirect
	github.com/otiai10/mint v1.6.3 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/ssor/bom v0.0.0-20170718123548-6386211fdfcf // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/net v0.40.0 // indirect
	golang.org/x/sync v0.14.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/kube-openapi v0.0.0-20250318190949-c8a335a9a2ff // indirect
	k8s.io/utils v0.0.0-20250502105355-0f33e8f1c979 // indirect
	sigs.k8s.io/json v0.0.0-20241014173422-cfa47c3a1cc8 // indirect
	sigs.k8s.io/randfill v1.0.0 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.7.0 // indirect
)

replace (
	// // https://github.com/golang/go/issues/33546#issuecomment-519656923
	// github.com/go-check/check => github.com/go-check/check v0.0.0-20180628173108-788fd7840127

	// github.com/golang/protobuf => github.com/golang/protobuf v1.4.2
	// github.com/grpc-ecosystem/grpc-gateway => github.com/grpc-ecosystem/grpc-gateway v1.16.0
	// github.com/improbable-eng/grpc-web => github.com/improbable-eng/grpc-web v0.0.0-20181111100011-16092bd1d58a

	// // Avoid CVE-2022-3064
	// gopkg.in/yaml.v2 => gopkg.in/yaml.v2 v2.4.0

	// // Avoid CVE-2022-28948
	// gopkg.in/yaml.v3 => gopkg.in/yaml.v3 v3.0.1

	// https://github.com/kubernetes/kubernetes/issues/79384#issuecomment-505627280
	k8s.io/api => k8s.io/api v0.33.1
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.33.1
	k8s.io/apimachinery => k8s.io/apimachinery v0.33.1
	k8s.io/apiserver => k8s.io/apiserver v0.33.1
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.33.1
	k8s.io/client-go => k8s.io/client-go v11.0.0+incompatible
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.33.1
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.33.1
	k8s.io/code-generator => k8s.io/code-generator v0.33.1
	k8s.io/component-base => k8s.io/component-base v0.33.1
	k8s.io/component-helpers => k8s.io/component-helpers v0.33.1
	k8s.io/controller-manager => k8s.io/controller-manager v0.33.1
	k8s.io/cri-api => k8s.io/cri-api v0.33.1
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.33.1
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.33.1
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.33.1
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.33.1
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.33.1
	k8s.io/kubectl => k8s.io/kubectl v0.33.1
	k8s.io/kubelet => k8s.io/kubelet v0.33.1
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.30.13
	k8s.io/metrics => k8s.io/metrics v0.33.1
	k8s.io/mount-utils => k8s.io/mount-utils v0.33.1
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.33.1
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.33.1

	// workaround for https://github.com/jaytaylor/html2text/issues/67
	github.com/olekukonko/tablewriter => github.com/olekukonko/tablewriter v0.0.5
)
