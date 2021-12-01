GO=go
GOCOVER=$(GO) tool cover
GOTEST=$(GO) test

.PHONY: tcover
tcover:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCOVER) -func=coverage.out
	$(GOCOVER) -html=coverage.out

