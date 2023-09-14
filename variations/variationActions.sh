#!/bin/bash

set -e

{
	echo "name: Variations"
	echo "description: Build variations of lovely"
	echo "inputs:"
	echo "  version:"
    echo "    description: 'Version to build'"
    echo "    required: true"
	echo "runs:"
	echo "  using: \"composite\""
	echo "  steps:"
	while IFS= read -r line; do
		linesplit=(${line// / })
		target="${linesplit[0]}"
		source="${linesplit[1]}"
		if [ "$source" = "BASE" ]
		then
			source="${BASE_LOVELY_IMAGE}"
		fi
		dockerfile="${linesplit[2]}"

		echo "    - name: Build and Push ${target}"
        echo "      uses: docker/build-push-action@v5"
        echo "      with:"
        echo "        context: ."
        echo "        file: variations/${dockerfile}"
        echo "        push: true"
        echo "        platforms: \${{ env.PLATFORMS }}"
        echo "        tags: ghcr.io/crumbhole/${target}:\${{ inputs.version }}"
        echo "        build-args: |"
		echo "          VERSION=\${{ inputs.version }}"
		echo "          PARENT=ghcr.io/crumbhole/${source}"
		echo "          NAME=${target}"
#		echo "  variation: ${target} from ${source} using ${dockerfile}"
	done < variations/variations.txt
} > ".github/actions/variations/action.yaml"
