
#Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build

#Docker parameters
DOCKERCMD=docker
DOCKERBUILD=$(DOCKERCMD) build

#BINARY_NAME
BINARY_MYAPP=myapp
BINARY_MYAPPCLI=myappcli

all:package
package:
        GO15VENDOREXPERIMENT=1 $(GOBUILD) -o $(BINARY_MYAPP) main.go  \
        && $(DOCKERBUILD) -t myapp:1.1.2.20141115_beta .
