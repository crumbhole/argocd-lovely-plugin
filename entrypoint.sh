#!/usr/bin/env bash
set -e
if [[ ! -z "${PLUGIN_NAME}" ]]; then
   /usr/local/bin/yq -i e ".metadata.name |= \"${PLUGIN_NAME}\"" /home/argocd/cmp-server/config/plugin.yaml
fi
/var/run/argocd/argocd-cmp-server
