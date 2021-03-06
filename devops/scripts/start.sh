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
MONGO_URI=mongodb://localhost:27017 \
AMQP_DSN=amqp://guest:guest@localhost:5672/ \
go run main.go) &

(cd backend/cmd/soxie/ && \
LOG_LEVEL=trace \
HOME_PATH=/ \
AMQP_DSN=amqp://guest:guest@localhost:5672/ \
go run main.go) &

cat
