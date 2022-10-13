FROM golang:1.19
WORKDIR /go/src/github.com/warp-contracts/sequencer/


#COPY go.mod go.sum ./
#RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=1 go build

#FROM alpine:3.16.2
FROM debian

WORKDIR /app

COPY --from=0 /go/src/github.com/warp-contracts/sequencer/sequencer ./
COPY config.yaml ./
RUN pwd
RUN ls -la

CMD ["/app/sequencer"]
