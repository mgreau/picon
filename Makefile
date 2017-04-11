# Makefile for picon
# -----------------------------------------------------------------
#
#        ENV VARIABLE
#
# -----------------------------------------------------------------

# go env vars
GO=$(firstword $(subst :, ,$(GOPATH)))
# list of pkgs for the project without vendor
PKGS=$(shell go list ./... | grep -v /vendor/)
DOCKER_IP=$(shell if [ -z "$(DOCKER_MACHINE_NAME)" ]; then echo 'localhost'; else docker-machine ip $(DOCKER_MACHINE_NAME); fi)
export GO15VENDOREXPERIMENT=1
 
# -----------------------------------------------------------------
#        Version
# -----------------------------------------------------------------

# version
VERSION=0.0.1
BUILDDATE=$(shell date -u '+%s')
BUILDHASH=$(shell git rev-parse --short HEAD)
VERSION_FLAG=-ldflags "-X main.Version=$(VERSION) -X main.GitHash=$(BUILDHASH) -X main.BuildStmp=$(BUILDDATE)"

# -----------------------------------------------------------------
#        Main targets
# -----------------------------------------------------------------

all: clean build

help:
	@echo $(GO)
	@echo "----- BUILD ------------------------------------------------------------------------------"
	@echo "all                  clean and build the project"
	@echo "clean                clean the project"
	@echo "build                build all libraries and binaries"
	@echo "----- TESTS && LINT ----------------------------------------------------------------------"
	@echo "format               format all packages"
	@echo "----- OTHERS -----------------------------------------------------------------------------"
	@echo "help                 print this message"

clean:
	@go clean
	@rm -Rf .tmp
	@rm -Rf .DS_Store
	@rm -Rf *.log
	@rm -Rf *.out
	@rm -Rf *.mem
	@rm -Rf *.test
	@rm -Rf build

build: format
	@go build -v $(VERSION_FLAG) -o $(GO)/bin/picon picon.go

format:
	@go fmt $(PKGS)

lint:
	@golint ./.