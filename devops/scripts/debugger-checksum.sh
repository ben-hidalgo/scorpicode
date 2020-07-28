#!/bin/bash
set -euo pipefail

CHECKSUM=""

CHECKSUM=${CHECKSUM}`cat ./backend/go.mod && cat ./backend/go.sum`

# Dockefile
CHECKSUM=${CHECKSUM}`cat ./devops/dockerfiles/debugger.Dockerfile`

# Deployment
CHECKSUM=${CHECKSUM}`cat ./devops/helmchart/templates/debugger/deployment.yaml`

md5sum <<< $CHECKSUM | awk '{ print $1 }'
