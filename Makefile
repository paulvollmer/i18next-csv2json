all: lint build

fmt:
	@go fmt

lint:
	@golint

build: fmt
	@go build

test: fmt
	@go test

test-cli: clean build
	./i18next-csv2json -i fixtures/sample.csv -o test
	cat test/en/sample.json
	cat test/de/sample.json

clean:
	@rm -rf test i18next-csv2json
