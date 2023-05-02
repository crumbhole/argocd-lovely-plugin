FROM golang:1.20.3 as builder
 # https://github.com/mikefarah/yq/releases
 # renovate: datasource=github-releases depName=mikefarah/yq
ARG YQ_VERSION=v4.33.3
 # https://github.com/kubernetes-sigs/kustomize/releases
 # renovate: datasource=github-releases depName=kubernetes-sigs/kustomize
ARG KUSTOMIZE_VERSION=v5.0.0
 # https://github.com/helm/helm/releases
 # renovate: datasource=github-releases depName=helm/helm
ARG HELM_VERSION=v3.11.3
 # https://github.com/helmfile/helmfile/releases
 # renovate: datasource=github-releases depName=helmfile/helmfile
ARG HELMFILE_VERSION=v0.153.1

RUN apt update && apt install -y curl wget unzip git golint && rm -rf /var/lib/apt/lists/*

ADD . /build
WORKDIR /build
# Install Dependencies
RUN /build/deps.sh

RUN make -j4

FROM alpine:3.17.3
ENV LOVELY_HELM_PATH=/usr/local/bin/helm
ENV LOVELY_KUSTOMIZE_PATH=/usr/local/bin/kustomize
COPY --from=builder /usr/local/bin/helm /usr/local/bin/helm
COPY --from=builder /usr/local/bin/helmfile /usr/local/bin/helmfile
COPY --from=builder /usr/local/bin/kustomize /usr/local/bin/kustomize
COPY --from=builder /build/build/argocd-lovely-plugin /usr/local/bin/argocd-lovely-plugin
COPY ./plugin.yaml /home/argocd/cmp-server/config/plugin.yaml
USER 999
# This does NOT exist inside the image, must be mounted from argocd
ENTRYPOINT [ "/var/run/argocd/argocd-cmp-server" ]
