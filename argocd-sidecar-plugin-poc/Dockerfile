FROM ghcr.io/crumbhole/argocd-lovely-plugin:stable as lovely

FROM quay.io/argoproj/argocd:v2.5.1 as argo
COPY --from=lovely argocd-lovely-plugin /usr/local/bin
