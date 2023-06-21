#!/usr/bin/sh
docker run --rm -ti --volume $PWD:/apps ignitehq/cli:latest $@
