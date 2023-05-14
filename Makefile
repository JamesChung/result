.PHONY: test
test:
	go test -v

.PHONY: testcov
testcov:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
