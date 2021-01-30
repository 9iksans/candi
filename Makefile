.PHONY : test

# unit test & calculate code coverage
test:
	@if [ -f coverage.txt ]; then rm coverage.txt; fi;
	@echo ">> running unit test and calculate coverage"
	@go test -race ./... -cover -coverprofile=coverage.txt -covermode=atomic \
		-coverpkg=$$(go list ./... | grep -v -e mocks -e codebase | tr '\n' ',')
	@go tool cover -func=coverage.txt
