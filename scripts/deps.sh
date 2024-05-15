#!/usr/bin/env bash

case $(uname -m) in
	x86_64)
		ARCH=amd64;;
	aarch64)
		ARCH=arm64;;
	*)
		echo Unknown architecture $(uname -m)
		exit 1;;
esac

echo Using architecture ${ARCH}

DEST=/usr/local/bin

# Install Helm
curl -SL https://get.helm.sh/helm-${HELM_VERSION}-linux-${ARCH}.tar.gz | tar -xz linux-${ARCH}/helm && mv linux-${ARCH}/helm ${DEST}

# Install Kustomize
curl -SL https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv$(echo ${KUSTOMIZE_VERSION}|tr -d kustomize/)/kustomize_v$(echo ${KUSTOMIZE_VERSION}|tr -d kustomize/)_linux_${ARCH}.tar.gz | tar -xzC ${DEST}

# Install yq
curl -L -s "https://github.com/mikefarah/yq/releases/download/${YQ_VERSION}/yq_linux_${ARCH}" -o ${DEST}/yq && chmod +x ${DEST}/yq

# Install Helmfile
curl -SL https://github.com/helmfile/helmfile/releases/download/${HELMFILE_VERSION}/helmfile_${HELMFILE_VERSION:1}_linux_${ARCH}.tar.gz | tar -xzC ${DEST}

# Install Helmwave
curl -SL https://github.com/helmwave/helmwave/releases/download/${HELMWAVE_VERSION}/helmwave_${HELMWAVE_VERSION:1}_linux_${ARCH}.tar.gz | tar -xzC ${DEST}
