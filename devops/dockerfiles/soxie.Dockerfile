FROM golang:1.13.6-alpine3.11

# setup
WORKDIR /app

# cache
COPY ./backend/go.mod ./
COPY ./backend/go.sum ./
RUN go mod download

# TODO: prevent pem from COPYing
# build
ADD ./backend/ /app/
RUN go build cmd/soxie/main.go

# run
RUN adduser -S -D -H -h /app mainapp
USER mainapp
CMD ["./main"]
