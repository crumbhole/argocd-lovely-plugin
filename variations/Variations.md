# Variations

This is a generator of variations on the lovely dockerfile.

## variations.txt

Three space separated fields per line, each representing a variation. The fields are
- Name of resulting dockerfile. In Github this will be pushed to ghcr.io/crumbhole/<name>:<version>
- Base image on from which to build this image. BASE magically means the most basic, versioned, sidecar image
- Dockerfile to run to perform the transformation from source to the result

## Building

`make variations` will build all variations locally.

## Updating

`make generate` to rebuild the github actions before committing the result.
