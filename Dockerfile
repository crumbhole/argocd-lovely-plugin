FROM golang:1.15.8 as builder
ADD . /build
WORKDIR /build
RUN go vet ./...
RUN go test ./...
RUN go build -o build/argocd-lovely-plugin

FROM alpine as putter
COPY --from=builder /build/build/argocd-lovely-plugin .
ENTRYPOINT [ "mv", "argocd-lovely-plugin", "/custom-tools/" ]