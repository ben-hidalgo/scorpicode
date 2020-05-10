#!/bin/bash
trap 'killall' INT

# https://unix.stackexchange.com/questions/55558/how-can-i-kill-and-wait-for-background-processes-to-finish-in-a-shell-script-whe

killall() {
    trap '' INT TERM # ignore INT and TERM while shutting down
    echo "shutting down..."
    kill -TERM 0
    wait
    echo "done"
}

(cd backend/cmd/roxie/ && \
LOG_LEVEL=trace \
LOCAL_HEADERS_PATH=../../../ignored/.headers \
go run main.go) &

# TODO: use kubectl output jsonpath
(cd backend/cmd/hats/ && \
LOG_LEVEL=trace \
MONGO_URI=mongodb://scuser:scpass@`minikube ip`:`kubectl get svc scorpicode-mongodb -n dev -o json | jq '.spec.ports[0].nodePort'`/scdata \
AMQP_DSN=amqp://scuser:scpass@`minikube ip`:`kubectl get svc scorpicode-rabbitmq-ha -n dev -o json | jq '.spec.ports[1].nodePort'`/schost \
go run main.go) &

(cd backend/cmd/soxie/ && \
LOG_LEVEL=trace \
go run main.go) &

cat
