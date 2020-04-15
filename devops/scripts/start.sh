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
go run main.go) &

(cd backend/cmd/hats/ && \
LOG_LEVEL=trace \
MONGO_URI=mongodb://hats:hats@`minikube ip`:`kubectl get svc scorpicode-mongodb -n dev -o json | jq '.spec.ports[0].nodePort'`/hats \
go run main.go) &

cat
