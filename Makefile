GOCMD = go
GOBUILD = $(GOCMD) build
GOGET = $(GOCMD) get -v
GOCLEAN = $(GOCMD) clean
GOINSTALL = $(GOCMD) install
GOTEST = $(GOCMD) test

.PHONY: all

all: build

test:
	$(GOTEST) -v -cover ./...

build:
	$(GOBUILD) -v -o stock-exchange

clean:
	$(GOCLEAN) -n -i -x
	rm -f $(GOPATH)/bin/stock-exchange
	rm -rf bin/stock-exchange

install: 
	$(GOINSTALL)
