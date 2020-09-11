#!/bin/bash
set -exuo pipefail

NAMESPACE=$1

(cd backend && \
MONGO_URI=mongodb://127.0.0.1:27017 \
AMQP_DSN=amqp://guest:guest@127.0.0.1:5672/ \
MONGO_DB=scdata \
go test ./... -v -count=1)
