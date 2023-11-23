FROM golang:1.21.4 as builder
 # https://github.com/mikefarah/yq/releases
 # renovate: datasource=github-releases depName=mikefarah/yq
ARG YQ_VERSION=v4.40.3
 # https://github.com/kubernetes-sigs/kustomize/releases
 # renovate: datasource=github-releases depName=kubernetes-sigs/kustomize
ARG KUSTOMIZE_VERSION=v5.0.3
 # https://github.com/helm/helm/releases
 # renovate: datasource=github-releases depName=helm/helm
ARG HELM_VERSION=v3.13.2
 # https://github.com/helmfile/helmfile/releases
 # renovate: datasource=github-releases depName=helmfile/helmfile
ARG HELMFILE_VERSION=v0.158.1

ARG LOVELY_VERSION

RUN apt update && apt install -y curl wget unzip git && rm -rf /var/lib/apt/lists/*

ADD . /build
WORKDIR /build
# Install Dependencies
RUN /build/scripts/deps.sh

RUN make plugin_versioned.yaml all -j4

FROM alpine:3.18.4
ENV LOVELY_HELM_PATH=/usr/local/bin/helm
ENV LOVELY_HELMFILE_PATH=/usr/local/bin/helmfile
ENV LOVELY_KUSTOMIZE_PATH=/usr/local/bin/kustomize
ENV LOVELY_PLUGINS=
ENV LOVELY_PREPROCESSORS=
ENV LOVELY_SIDECAR=true
COPY --from=builder /usr/local/bin/yq /usr/local/bin/yq
COPY --from=builder /usr/local/bin/helm /usr/local/bin/helm
COPY --from=builder /usr/local/bin/helmfile /usr/local/bin/helmfile
COPY --from=builder /usr/local/bin/kustomize /usr/local/bin/kustomize
COPY --from=builder /build/build/argocd-lovely-plugin /usr/local/bin/argocd-lovely-plugin
RUN apk add git bash --no-cache

USER 999
COPY --from=builder --chown=999 /build/plugin_versioned.yaml /home/argocd/cmp-server/config/plugin.yaml
COPY entrypoint.sh /entrypoint.sh
# /var/run/argocd/argocd-cmp-server does NOT exist inside the image, must be mounted from argocd
ENTRYPOINT [ "/entrypoint.sh" ]
