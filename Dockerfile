FROM golang:1.17.1 as builder

RUN apt update && apt install -y curl wget unzip

# Install Helm
RUN curl -s https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash

# Install Kustomize
RUN curl -s "https://raw.githubusercontent.com/kubernetes-sigs/kustomize/master/hack/install_kustomize.sh" | bash

ADD . /build
WORKDIR /build
RUN go vet ./...
RUN go test ./...
RUN go build -o build/argocd-lovely-plugin

FROM alpine as putter
COPY --from=builder /build/build/argocd-lovely-plugin .
ENTRYPOINT [ "mv", "argocd-lovely-plugin", "/custom-tools/" ]
