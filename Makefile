help:
	@echo "lint             run lint"

.PHONY: lint
lint:
	gofmt -s -w . cmd/parse_metrics
	golint .
	golint cmd/parse_metrics
	go vet
