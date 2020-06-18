#!/bin/bash
set -exuo pipefail

eval $(minikube docker-env)

docker build . -f devops/dockerfiles/hats.Dockerfile     -t hats:`./devops/scripts/go-checksum.sh hats`
docker build . -f devops/dockerfiles/website.Dockerfile  -t website:`./devops/scripts/js-checksum.sh website`
docker build . -f devops/dockerfiles/frontend.Dockerfile -t frontend:`./devops/scripts/js-checksum.sh frontend`
docker build . -f devops/dockerfiles/roxie.Dockerfile    -t roxie:`./devops/scripts/go-checksum.sh roxie`
docker build . -f devops/dockerfiles/soxie.Dockerfile    -t soxie:`./devops/scripts/go-checksum.sh soxie`
docker build . -f devops/dockerfiles/debugger.Dockerfile -t debugger:latest

docker images
