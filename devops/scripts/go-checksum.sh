#!/bin/bash
set -euo pipefail

SERVICE=$1

CHECKSUM=""

CHECKSUM=${CHECKSUM}`cat ./backend/go.mod && cat ./backend/go.sum`

# Dockefile
CHECKSUM=${CHECKSUM}`cat ./devops/dockerfiles/${SERVICE}.Dockerfile`
# Helm 
CHECKSUM=${CHECKSUM}`find ./devops/helmchart/templates/${SERVICE} -type f -exec md5sum "{}" +`
# package main
CHECKSUM=${CHECKSUM}`find ./backend/cmd/${SERVICE} -type f -exec md5sum "{}" +`

# dependencies
for i in $(cd backend && go list -json -f '{{.Imports}}' ./cmd/hats | jq -r '.Deps[]' | grep "^backend/"); do
    CHECKSUM=${CHECKSUM}`find ./${i} -type f -exec md5sum "{}" +`
done

md5sum <<< $CHECKSUM | awk '{ print $1 }'
