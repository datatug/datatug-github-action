FROM golang:1.15 as builder

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 go build -v -o datatug-github-action .

FROM alpine:latest

COPY --from=builder /app/datatug-github-action /datatug-github-action

ENTRYPOINT ["/datatug-github-action"]
