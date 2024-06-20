FROM golang:1-alpine3.20

COPY . .

ENTRYPOINT ["go", "run", "cmd/main.go"]