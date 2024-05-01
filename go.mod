module github.com/crumbhole/argocd-lovely-plugin

go 1.22.2

require (
	github.com/evanphx/json-patch v5.9.0+incompatible
	github.com/go-andiamo/splitter v1.2.5
	github.com/gomarkdown/markdown v0.0.0-20240419095408-642f0ee99ae2
	github.com/hexops/gotextdiff v1.0.3
	github.com/otiai10/copy v1.14.0
	github.com/stretchr/testify v1.9.0
	gopkg.in/yaml.v3 v3.0.1
	jaytaylor.com/html2text v0.0.0-20230321000545-74c2419ad056
	k8s.io/apimachinery v0.30.0
	sigs.k8s.io/kustomize/api v0.17.1
	sigs.k8s.io/kustomize/kyaml v0.17.0
	sigs.k8s.io/yaml v1.4.0
)

require (
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/go-errors/errors v1.5.1 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/gnostic-models v0.6.8 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/ssor/bom v0.0.0-20170718123548-6386211fdfcf // indirect
	golang.org/x/net v0.23.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/protobuf v1.33.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	k8s.io/klog/v2 v2.120.1 // indirect
	k8s.io/kube-openapi v0.0.0-20240228011516-70dd3763d340 // indirect
	k8s.io/utils v0.0.0-20240310230437-4693a0247e57 // indirect
	sigs.k8s.io/json v0.0.0-20221116044647-bc3834ca7abd // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.1 // indirect
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
	k8s.io/api => k8s.io/api v0.30.0
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.30.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.30.0
	k8s.io/apiserver => k8s.io/apiserver v0.30.0
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.30.0
	k8s.io/client-go => k8s.io/client-go v11.0.0+incompatible
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.30.0
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.30.0
	k8s.io/code-generator => k8s.io/code-generator v0.30.0
	k8s.io/component-base => k8s.io/component-base v0.30.0
	k8s.io/component-helpers => k8s.io/component-helpers v0.30.0
	k8s.io/controller-manager => k8s.io/controller-manager v0.30.0
	k8s.io/cri-api => k8s.io/cri-api v0.30.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.30.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.30.0
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.30.0
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.30.0
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.30.0
	k8s.io/kubectl => k8s.io/kubectl v0.30.0
	k8s.io/kubelet => k8s.io/kubelet v0.30.0
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.30.0
	k8s.io/metrics => k8s.io/metrics v0.30.0
	k8s.io/mount-utils => k8s.io/mount-utils v0.30.0
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.30.0
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.30.0
)
