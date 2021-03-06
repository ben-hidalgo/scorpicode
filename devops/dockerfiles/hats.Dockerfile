FROM golang:1.13.6-alpine3.11

# setup
WORKDIR /app
RUN apk add curl
RUN apk add jq

# cache
COPY ./backend/go.mod ./
COPY ./backend/go.sum ./
RUN go mod download

# build
ADD ./backend/ /app/
RUN go build cmd/hats/main.go

# run
RUN adduser -S -D -H -h /app mainapp
USER mainapp
CMD ["./main"]
