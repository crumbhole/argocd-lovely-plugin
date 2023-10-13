#!/bin/bash

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

# Install Helm
curl -SL https://get.helm.sh/helm-${HELM_VERSION}-linux-${ARCH}.tar.gz | tar -xz linux-${ARCH}/helm && mv linux-${ARCH}/helm /usr/local/bin/

# Install Kustomize
curl -SL https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2F$(echo ${KUSTOMIZE_VERSION}|tr -d kustomize/)/kustomize_$(echo ${KUSTOMIZE_VERSION}|tr -d kustomize/)_linux_${ARCH}.tar.gz | tar -xzC /usr/local/bin

# Install yq
curl -L -s "https://github.com/mikefarah/yq/releases/download/${YQ_VERSION}/yq_linux_${ARCH}" -o /usr/local/bin/yq && chmod +x /usr/local/bin/yq

# Install Helmfile
curl -SL https://github.com/helmfile/helmfile/releases/download/${HELMFILE_VERSION}/helmfile_${HELMFILE_VERSION:1}_linux_${ARCH}.tar.gz | tar -xzC /usr/local/bin
