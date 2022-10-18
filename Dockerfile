FROM golang:1.19
WORKDIR /go/src/github.com/warp-contracts/sequencer/

COPY . .
RUN CGO_ENABLED=1 go build

FROM debian:stable

WORKDIR /app

COPY --from=0 /go/src/github.com/warp-contracts/sequencer/sequencer ./
COPY config.yaml ./

ENV LOG_FORMAT=json
ENV SEQUENCER_MODE=release
ENV POSTGRES_SSLMODE=require

CMD ["/app/sequencer"]
