PACKAGE_NAME          := github.com/cloudquery/cloudquery/plugins/destination/elasticsearch

.PHONY: test
test:
# we clean the cache to avoid scenarios when we change something in the db and we want to retest without noticing nothing run
	go clean -testcache
# we run the tests 5 times to avoid flaky tests
	go test -race -timeout 3m ./... || go test -race -timeout 3m ./... || go test -race -timeout 3m ./... || go test -race -timeout 3m ./... || go test -race -timeout 3m ./...

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml
