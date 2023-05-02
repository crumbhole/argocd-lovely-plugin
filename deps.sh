#!/bin/bash

# Install Helm
curl -SL https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz | tar -xz linux-amd64/helm && mv linux-amd64/helm /usr/local/bin/

# Install Kustomize
curl -SL https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2F${KUSTOMIZE_VERSION}/kustomize_${KUSTOMIZE_VERSION}_linux_amd64.tar.gz | tar -xzC /usr/local/bin

# Install yq
curl -L -s "https://github.com/mikefarah/yq/releases/download/${YQ_VERSION}/yq_linux_amd64" -o /usr/local/bin/yq && chmod +x /usr/local/bin/yq

# Install Helmfile
curl -SL https://github.com/helmfile/helmfile/releases/download/${HELMFILE_VERSION}/helmfile_${HELMFILE_VERSION:1}_linux_amd64.tar.gz | tar -xzC /usr/local/bin
