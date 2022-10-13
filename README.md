# Configuration
There are two ways of configuring supported - yaml files and environment variables.

For environment variable will rewrite value from the yaml file.

Environment variable should be in the upper case and dots should be replaced with underscore.
For example, if we want to rewrite parameter "postgres.host", the environment variable should be named POSTGRES_HOST.


## Environments
config.yaml file will be included for all environments.

If you running tests, config_test.yaml will be included additionally.

# Local run and development

## Run the application

It requires to set walletJwk to run the sequencer locally:
```sh
export ARWEAVE_WALLETJWK="some arconnect key"
```
You can run postgres locally:
```sh
docker-compose -f _tests/docker-compose/postgres/docker-compose.yml up -d
```

You can run application using command line:
```sh
go run .
```
Or from you IDE you can run main.go file

## Tests

Running all tests:
```sh
go run ./...
```

Tests should be written to run in parallel (t.Parallel()).
