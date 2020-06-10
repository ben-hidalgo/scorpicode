#!/bin/bash
set -exuo pipefail

# var MongoURI = "mongodb://scuser:scpass@localhost:27017/scdata"
# var AmqpDsn = "amqp://rabbit:rabbit@localhost:5672/"

NAMESPACE=$1

(cd backend && \
AMQP_DSN=amqp://scuser:scpass@`minikube ip`:`kubectl get svc scorpicode-rabbitmq-ha -n ${NAMESPACE} -o json | jq '.spec.ports[1].nodePort'`/schost \
MONGO_URI=mongodb://scuser:scpass@`minikube ip`:`kubectl get svc scorpicode-mongodb -n ${NAMESPACE} -o json | jq '.spec.ports[0].nodePort'`/scdata \
MONGO_DB=scdata \
go test ./... -v -count=1)
