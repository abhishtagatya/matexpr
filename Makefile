# Define variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run

# Binary name
BINARY_NAME=matexpr

# Directories
SRC=main.go
OUT=./bin

all: clean brun

build:
	$(GOBUILD) -o $(OUT)/$(BINARY_NAME) $(SRC)

clean:
	$(GOCLEAN)
	rm -f $(OUT)/$(BINARY_NAME)

brun:
	$(GOBUILD) -o $(OUT)/$(BINARY_NAME) $(SRC)
	$(OUT)/$(BINARY_NAME) "$(EXPR)"

run:
	$(GORUN) $(SRC) "$(EXPR)"

.PHONY: all build clean run
