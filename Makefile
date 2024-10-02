# Formats the code
fmt:
	go fmt github.com/danielronalds/...

GOTESTSUM_VERSION := $(shell gotestsum --version 2>/dev/null)

# Runs all tests
test:
ifdef GOTESTSUM_VERSION
	gotestsum --format testname github.com/danielronalds/...
else
	go test github.com/danielronalds/...
endif
