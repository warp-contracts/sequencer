#!/bin/bash

set -e

usage() {
    cat <<EOF
Creates image of sequencer with the configuration for the given environment.
Image handles updates between specified versions.

Mandatory args:
  -e,--env        environment (devnet|testnet|mainnet) Default: devnet
  -v,--version    version to release
  -f,--from       version to update from

Optional args:
  --nobuild       do not build, do not modify git
  -h,--help       print this help

Examples: 

  $1 --env devnet --from 0.0.1 --version 0.0.2 

EOF
    return 0
}

# Ensure that all required commands are installed
check_commands() {
    for cmd in "docker" "awk"; do
        if ! [ -x "$(command -v $cmd)" ]; then
            echo "Error: $cmd is not installed" >&2
            exit 1
        fi
    done
}

is_semver() {
    if [[ $1 =~ ^(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)\.(0|[1-9][0-9]*)(-((0|[1-9][0-9]*|[0-9]*[a-zA-Z-][0-9a-zA-Z-]*)(\.(0|[1-9][0-9]*|[0-9]*[a-zA-Z-][0-9a-zA-Z-]*))*))?(\+([0-9a-zA-Z-]+(\.[0-9a-zA-Z-]+)*))?$ ]]; then
        echo "$1"
    else
        echo ""
    fi
}

parse_command_line_options() {
    # Defaults
    ENV=devnet
    VERSION=""
    FROM_VERSION=""
    NO_BUILD="false"

    optspec="he:v:f:-:"
    while getopts "$optspec" optchar; do
        case "${optchar}" in
        -)
            case "${OPTARG}" in
            env)
                ENV="${!OPTIND}"
                OPTIND=$(($OPTIND + 1))
                ;;
            version)
                VERSION="${!OPTIND}"
                OPTIND=$(($OPTIND + 1))
                ;;
            from)
                FROM_VERSION="${!OPTIND}"
                OPTIND=$(($OPTIND + 1))
                ;;
            nobuild)
                NO_BUILD="true"
                ;;
            help)
                usage
                ;;
            *)
                if [ "$OPTERR" = 1 ] && [ "${optspec:0:1}" != ":" ]; then
                    echo "Unknown option --${OPTARG}"
                    exit 1
                fi
                ;;
            esac
            ;;
        h)
            usage
            ;;
        e)
            ENV=$OPTARG
            ;;
        v)
            VERSION=$OPTARG
            ;;
        f)
            FROM_VERSION=$OPTARG
            ;;
        esac
    done

    if [[ -z "$VERSION" ]]; then
        echo "Missing release version"
        exit 1
    fi

    if [[ -z "$(is_semver $VERSION)" ]]; then
        echo "Invalid release version: $VERSION"
        exit 1
    fi

    if [[ -z "$FROM_VERSION" ]]; then
        echo "Missing from release version"
        exit 1
    fi

    if [[ -z "$(is_semver $VERSION)" ]]; then
        echo "Invalid from release version: $VERSION"
        exit 1
    fi

}

# Ensure that user has access before doing anything
configure() {
    case ${ENV} in
    devnet)
        FOLDER="./network/devnet/genessis"
        ;;
    testnet)
        FOLDER="./network/testnet/genessis"
        NO_BUILD=true
        ;;
    mainnet)
        FOLDER="./network/mainnet/genessis"
        NO_BUILD=true
        ;;
    *)
        echo "Unknown environment: $ENV"
        echo "Possible values: devnet, testnet, mainnet"
        exit 1
        ;;
    esac
}

run() {
    if [[ "$NO_BUILD" == "false" ]]; then
        echo "Building and releasing in git: $VERSION"
        pushd $WARP_SEQUENCER_REPO_PATH
        make
        git release v$VERSION
        popd
    fi

    export DOCKER_BUILDKIT=0 
    docker build --build-arg="ENV=${ENV}" -t "warpredstone/sequencer:v${VERSION}-${ENV}" .


    
}

parse_command_line_options "$@"
check_commands
configure
run
