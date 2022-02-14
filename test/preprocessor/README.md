This example:
- Uses the pre-processor to change Chart.yaml

This is example is deliberately contrived, but demonstrates argocd-lovely-plugin's ability to make use of preprocessor plugins to perform changes prior to helm (or kustomize).
Here we use sed to modify the Chart.yaml, and then yq to do the same but for a different part of it. Using this we can perhaps have staging and production use identical git paths, but a different chart version by adding the appropriate environment variable from [application sets](https://argocd-applicationset.readthedocs.io/en/stable/).
