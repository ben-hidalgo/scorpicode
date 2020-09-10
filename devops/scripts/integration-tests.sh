#!/bin/bash
set -exuo pipefail

NAMESPACE=$1

(cd backend && \
MONGO_URI=mongodb://localhost:27017 \
AMQP_DSN=amqp://guest:guest@localhost:5672/ \
MONGO_DB=scdata \
go test ./... -v -count=1)
