#!/bin/bash

set -e

inerr=0
if [[ -z "${LOVELY_VERSION}" ]]; then
	echo Must set LOVELY_VERSION
	inerr=1
fi
if [[ -z "${IMAGE_REPO}" ]]; then
	echo Must set IMAGE_REPO
	inerr=1
fi
if [[ -z "${BASE_LOVELY_IMAGE}" ]]; then
	echo Must set BASE_LOVELY_IMAGE
	inerr=1
fi

if [ "$inerr" -eq "1" ]; then
	exit 1
fi

images=()

function imageName() {
	NAME="$1"
	echo "${IMAGE_REPO}/${NAME}:${LOVELY_VERSION}"
}

function buildImage() {
	NAME="$1"
	PARENT="$2"
	DOCKERFILE="$3"
	IMAGE="$(imageName "${NAME}" )"
	docker build -f "variations/${DOCKERFILE}" --build-arg="VERSION=${LOVELY_VERSION}" --build-arg="NAME=${NAME}" --build-arg="PARENT=${IMAGE_REPO}/${PARENT}" variations -t "${IMAGE}"
	images+=( "${IMAGE}" )
}

function pushImage() {
	IMAGE="$1"
	echo Would push "${IMAGE}"
	#docker push "${IMAGE}"
}

echo --- Building variations
while IFS= read -r line; do
	linesplit=(${line// / })
	target="${linesplit[0]}"
	source="${linesplit[1]}"
	if [ "$source" = "BASE" ]
	then
		source="${BASE_LOVELY_IMAGE}"
	fi
	dockerfile="${linesplit[2]}"
	echo "variation: ${target} from ${source} using ${dockerfile}"
	buildImage "${target}" "${source}" "${dockerfile}"
done < variations/variations.txt
echo --- Pushing variations

for IMG in "${images[@]}"
do
	pushImage "${IMG}"
done

echo --- Done variations
