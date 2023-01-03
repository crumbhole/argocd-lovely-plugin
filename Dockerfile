FROM golang:1.19.4 as builder
ARG YQ_VERSION=4.28.2
ARG KUSTOMIZE_VERSION=4.5.7 #https://github.com/kubernetes-sigs/kustomize/releases
ARG HELM_VERSION=3.10.3

RUN apt update && apt install -y curl wget unzip git golint && rm -rf /var/lib/apt/lists/*

# Install Helm
RUN curl -SL https://get.helm.sh/helm-v${HELM_VERSION}-linux-amd64.tar.gz | tar -xz linux-amd64/helm && mv linux-amd64/helm /usr/local/bin/

# Install Kustomize
RUN curl -SL https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2Fv${KUSTOMIZE_VERSION}/kustomize_v${KUSTOMIZE_VERSION}_linux_amd64.tar.gz | tar -xzC /usr/local/bin

# Install yq
RUN curl -L -s "https://github.com/mikefarah/yq/releases/download/v${YQ_VERSION}/yq_linux_amd64" -o /usr/local/bin/yq && chmod +x /usr/local/bin/yq

ADD . /build
WORKDIR /build
RUN make -j4

FROM alpine:3.17.0 as putter
COPY --from=builder /build/build/argocd-lovely-plugin .
USER 999
ENTRYPOINT [ "cp", "argocd-lovely-plugin", "/custom-tools/" ]
