#!/bin/bash
set -exuo pipefail

NAMESPACE=$1

(cd backend && \
MONGO_DB=scdata \
go test ./... -v -count=1)
