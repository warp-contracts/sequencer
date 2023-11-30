ARG VERSION=testversion
ARG FROM_VERSION=testfromversion
ARG ENV=dev

# Use binaries from previous image
FROM warpredstone/sequencer:${FROM_VERSION}-${ENV} as previous

# Build the sequencer binary
FROM golang:1.21-alpine3.18 as current
LABEL stage=sequencer-builder
RUN apk add --update make build-base curl git upx

RUN go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@latest

WORKDIR /app
COPY .gopath~ .gopath~
COPY Makefile .
COPY go.mod .
COPY go.sum .
COPY app app
COPY crypto crypto
COPY docs docs
COPY tools tools
COPY x x
COPY cmd cmd
COPY testutil testutil
COPY .git .git

RUN make build-optimized
 
# Minimal output image
FROM alpine:3.18
ARG VERSION=testversion
ARG ENV=dev

RUN apk add --update jq

# Cosmovisor setup
RUN mkdir -p /root/cosmovisor/genesis/bin

# Cosmos setup
RUN mkdir -p /root/.sequencer/data
RUN echo '{"height":"0","round":0,"step":0}' > /root/.sequencer/data/priv_validator_state.json

# Genesis setup
COPY network/${ENV}/genesis/prev_sort_keys.json /root/.sequencer/genesis/prev_sort_keys.json
COPY network/${ENV}/genesis/arweave_block.json /root/.sequencer/genesis/arweave_block.json

# Executables
COPY --from=previous /root/cosmovisor/ /root/cosmovisor/
COPY --from=current  /go/bin/cosmovisor /usr/local/bin/cosmovisor
COPY --from=current  /app/bin/sequencer /root/cosmovisor/upgrades/${VERSION}/bin/sequencer

COPY utils/docker-entrypoint.sh /app/docker-entrypoint.sh

ENTRYPOINT [ "/app/docker-entrypoint.sh" ]