# Go parameters
MAIN_PATH=apimusic.go
BINARY_PATH=bin
BINARY_NAME=$(BINARY_PATH)/apimusic

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORM=rm -rf

all: test build
build:
	$(GOBUILD) -o $(BINARY_NAME) -race $(MAIN_PATH)

run:
	$(GOBUILD) -o $(BINARY_NAME) -race $(MAIN_PATH)
	./$(BINARY_NAME) $(arg1)
clean:
	$(GOCLEAN) $(MAIN_PATH)
	$(GORM) $(BINARY_PATH)
