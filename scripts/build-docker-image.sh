#!/usr/bin/env sh

if [ "$1" != "" ]; then 
    echo Usage: build-docker-image.sh
    echo ""
    echo Builds an production Docker image in the local system and tags it with
    echo the ID of the current git HEAD
    exit 1
fi

IMAGE_REPO="sourcenetwork/sourcehub"

commit=$(git rev-parse HEAD)

docker image build -t $IMAGE_REPO:$commit .
docker image tag $IMAGE_REPO:$commit $IMAGE_REPO:latest
