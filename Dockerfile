FROM golang:1.19.3 as builder

RUN apt update && apt install -y curl wget unzip git && rm -rf /var/lib/apt/lists/*

# Install Helm
RUN curl -s https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

# Install Kustomize
RUN curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh" | bash && mv /go/kustomize /usr/local/bin

# Install yq
RUN curl -L -s "https://github.com/mikefarah/yq/releases/download/v4.27.5/yq_linux_amd64" -o /usr/local/bin/yq && chmod +x /usr/local/bin/yq

ADD . /build
WORKDIR /build
RUN make -j4

FROM alpine:3.17.0 as putter
COPY --from=builder /build/build/argocd-lovely-plugin .
USER 999
ENTRYPOINT [ "cp", "argocd-lovely-plugin", "/custom-tools/" ]
