# `ALLOW_GITCHECKOUT`

Normally lovely makes a copy of the application yaml before rendering it to allow it to safely modify it with Pre-Processing, merge/patch operations.

If you set `LOVELY_ALLOW_GITCHECKOUT` we will instead use git operations to ensure the repository is clean before and after processing. This is **only** safe if you've told ArgoCD that you're going to do this by enabling [repository locking](https://argo-cd.readthedocs.io/en/stable/user-guide/config-management-plugins/).

If you can't be bothered to read the page on it, you basically must configure your plugin with `lockRepo: true`. Even if you can be bothered to read the page you also get to keep all the pieces of your cluster and deployments after this has broken it. Insert disclaimer here.

This will probably make ArgoCD run slower especially if you have a lot of repositories in a single repository.

However, it should allow kustomize base paths outside of the application directory to work as discussed in Issue #66.

You should be able safely create two plugins at the ArgoCD config map level one with this enabled and one without it and mix applications within the same repository.
