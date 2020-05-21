#!/bin/bash
set -euo pipefail

SERVICE=$1

CHECKSUM=""

# Dockefile
CHECKSUM=${CHECKSUM}`cat ./devops/dockerfiles/${SERVICE}.Dockerfile`
# Helm 
CHECKSUM=${CHECKSUM}`find ./devops/helmchart/templates/${SERVICE} -type f -exec md5sum "{}" +`
# 
CHECKSUM=${CHECKSUM}`find ./${SERVICE} -path node_modules -prune  -type f -exec md5sum "{}" +`

md5sum <<< $CHECKSUM | awk '{ print $1 }'
