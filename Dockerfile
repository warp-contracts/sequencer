FROM golang:1.19 as builder
WORKDIR /go/src/github.com/warp-contracts/sequencer/

COPY . .
RUN CGO_ENABLED=1 go build

FROM debian:stable as runner

RUN apt update && \
    apt install -y ca-certificates && \
    # curl needs for healthcheck in ECS
    apt install -y curl && \
    apt-get clean autoclean && \
    apt-get autoremove --yes && \
    rm -rf /var/lib/{apt,dpkg,cache,log}/

WORKDIR /app

COPY --from=builder /go/src/github.com/warp-contracts/sequencer/sequencer ./
COPY config.yaml ./

ENV LOG_FORMAT=json
ENV SEQUENCER_MODE=release
ENV POSTGRES_SSLMODE=require

CMD ["/app/sequencer"]
