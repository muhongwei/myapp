
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
        $(GOBUILD) -o $(BINARY_MYAPP) main.go  && $(GOBUILD) -o $(BINARY_MYAPPCLI) myapp-cli.go \
        && $(DOCKERBUILD) -t myapp:2.0 .
