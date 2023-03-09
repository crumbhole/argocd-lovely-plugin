FROM golang:1.20.2 as builder
 # https://github.com/mikefarah/yq/releases
 # renovate: datasource=github-releases depName=mikefarah/yq
ARG YQ_VERSION=v4.30.8
 # https://github.com/kubernetes-sigs/kustomize/releases
 # renovate: datasource=github-releases depName=kubernetes-sigs/kustomize
ARG KUSTOMIZE_VERSION=v5.0.0
 # https://github.com/helm/helm/releases
 # renovate: datasource=github-releases depName=helm/helm
ARG HELM_VERSION=v3.11.1

RUN apt update && apt install -y curl wget unzip git golint && rm -rf /var/lib/apt/lists/*

# Install Helm
RUN curl -SL https://get.helm.sh/helm-${HELM_VERSION}-linux-amd64.tar.gz | tar -xz linux-amd64/helm && mv linux-amd64/helm /usr/local/bin/

# Install Kustomize
RUN curl -SL https://github.com/kubernetes-sigs/kustomize/releases/download/kustomize%2F${KUSTOMIZE_VERSION}/kustomize_${KUSTOMIZE_VERSION}_linux_amd64.tar.gz | tar -xzC /usr/local/bin

# Install yq
RUN curl -L -s "https://github.com/mikefarah/yq/releases/download/${YQ_VERSION}/yq_linux_amd64" -o /usr/local/bin/yq && chmod +x /usr/local/bin/yq

ADD . /build
WORKDIR /build
RUN make -j4

FROM alpine:3.17.2
ENV LOVELY_HELM_PATH=/usr/local/bin/helm
ENV LOVELY_KUSTOMIZE_PATH=/usr/local/bin/kustomize
COPY --from=builder /usr/local/bin/helm /usr/local/bin/helm
COPY --from=builder /usr/local/bin/kustomize /usr/local/bin/kustomize
COPY --from=builder /build/build/argocd-lovely-plugin /usr/local/bin/argocd-lovely-plugin
COPY ./plugin.yaml /home/argocd/cmp-server/config/plugin.yaml
USER 999
# This does NOT exist inside the image, must be mounted from argocd
ENTRYPOINT [ "/var/run/argocd/argocd-cmp-server" ]
