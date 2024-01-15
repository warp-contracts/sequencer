PACKAGE  = sequencer
GOROOT   = $(CURDIR)/.gopath~
GOPATH   = $(CURDIR)/.gopath~
BIN      = $(GOPATH)/bin
BASE     = $(GOPATH)/cmd/sequencerd
PATH    := bin:$(PATH)
GO       = go
VERSION ?= $(shell git describe --tags --always --match=v* 2> /dev/null || cat $(CURDIR)/.version 2> /dev/null || echo v0)
COMMIT  ?= $(shell git rev-parse HEAD | tr -d '\n')
DATE    ?= $(shell date +%FT%T%z)


# Build variables
ENV	         ?= devnet
FROM_VERSION ?= v0.0.70

# Tools
VERSION_REGEX     := ^v(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)$
IS_VERSION_SEMVER := $(shell echo ${VERSION} | egrep "${tag_regex}")

export GOPATH

# Display utils
V = 0 # Verbose output, change to 1 to print commands
Q = $(if $(filter 1,$V),,@) # Conditionally print output
M = $(shell printf "\033[34;1m▶\033[0m")

# Default target
.PHONY: all
all:  build lint | $(BASE); $(info $(M) built and lint everything!) @

# Setup
$(BASE): ; $(info $(M) setting GOPATH…)
	@mkdir -p $(dir $@)
	@ln -sf $(CURDIR) $@

# External tools 
$(BIN):
	@mkdir -p $@
$(BIN)/%: | $(BIN) ; $(info $(M) installing $(REPOSITORY)…)
	$Q tmp=$$(mktemp -d); \
	   env GO111MODULE=on GOPATH=$$tmp GOBIN=$(BIN) $(GO) install $(REPOSITORY) \
		|| ret=$$?; \
	   exit $$ret

GOLANGCILINT = $(BIN)/golangci-lint
$(BIN)/golangci-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.53.3

GENTOOL = $(BIN)/gentool
$(BIN)/gentool: REPOSITORY=gorm.io/gen/tools/gentool@latest

.PHONY: lint
lint: $(GOLANGCILINT) | $(BASE) ; $(info $(M) running golangci-lint) @
	$Q $(GOLANGCILINT) run 

.PHONY: version
version:
	$Q echo -n $(VERSION) > .version

.PHONY: clean
clean:
	rm -rf bin/$(PACKAGE) .gopath~

# Build targets
.PHONY: build
build:  | $(BASE); $(info $(M) building executable…) @
	$Q cd $(BASE) && $(GO) build \
		-tags release \
		-ldflags "-s -w"  \
		-o bin/$(PACKAGE) cmd/sequencerd/main.go

.PHONY: build-race
build-race:  | $(BASE); $(info $(M) building executable…) @
	$Q cd $(BASE) && $(GO) build -race \
		-tags release \
		-o bin/$(PACKAGE) main.go

.PHONY: build-optimized
build-optimized:  | $(BASE); $(info $(M) building executable…) @
	cd $(BASE)
	$(GO) build -tags release -ldflags "-s -w \
	-X github.com/cosmos/cosmos-sdk/version.Name=WarpSequencer \
	-X github.com/cosmos/cosmos-sdk/version.AppName=sequencerd \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	-X github.com/cosmos/cosmos-sdk/version.BuildTags=env=$(ENV),build_time=$(DATE)" \
	-o bin/$(PACKAGE) cmd/sequencerd/main.go && \
 	upx -q --best --lzma bin/$(PACKAGE)

.PHONY: test
test:
	$(GO) test ./...

.PHONY: docker-build
docker-build: all | ; $(info $(M) building docker container) @ 
ifeq (IS_VERSION_SEMVER,)
	$(error Version has to be semver, have you released a new version? Version is $(VERSION))
endif
	DOCKER_BUILDKIT=0 VERSION=$(VERSION) FROM_VERSION=$(FROM_VERSION) ENV=$(ENV) \
	docker buildx bake 	-f docker-bake.hcl 	--load


.PHONY: docker-push
docker-push: all | ; $(info $(M) pushing docker container) @ 
ifeq (IS_VERSION_SEMVER,)
	$(error Version has to be semver, have you released a new version? Version is $(VERSION))
endif
	docker login
	DOCKER_BUILDKIT=0 VERSION=$(VERSION) FROM_VERSION=$(FROM_VERSION) ENV=$(ENV) \
	docker buildx bake -f docker-bake.hcl --progress=plain --push

.PHONY: docker-genesis
docker-genesis: all | ; $(info $(M) pushing docker container as the genesis image) @ 
ifeq (IS_VERSION_SEMVER,)
	$(error Version has to be semver, have you released a new version? Version is $(VERSION))
endif
	docker login
	DOCKER_BUILDKIT=0 VERSION=$(VERSION) ENV=$(ENV) \
	docker buildx bake genesis -f docker-bake.hcl --progress=plain --push

.PHONY: docker-build-genesis
docker-build-genesis: all | ; $(info $(M) build docker container as the genesis image) @ 
ifeq (IS_VERSION_SEMVER,)
	$(error Version has to be semver, have you released a new version? Version is $(VERSION))
endif
	docker login
	DOCKER_BUILDKIT=0 VERSION=$(VERSION) ENV=$(ENV) \
	docker buildx bake genesis -f docker-bake.hcl --progress=plain --load

.PHONY: docker-run
docker-run:  | ; $(info $(M) running docker container) @ 
	docker compose --profile sequencer build
	docker compose --profile sequencer up 

B = git checkout $1 -b _build-$1 && \
 cd $(BASE) && \
 $(GO) build -tags release -ldflags "-s -w" -o bin/upgrades/$1/bin/$(PACKAGE) cmd/sequencerd/main.go && \
 upx -q --best --lzma bin/upgrades/$1/bin/$(PACKAGE) && \
 git checkout main ; \
 git branch -d _build-$1

.PHONY: build-all-updates
build-all-updates: | ; $(info $(M) build every major update) @ 
	$(call B,v0.0.70)
	@mv bin/upgrades/v0.0.70 bin/genesis

