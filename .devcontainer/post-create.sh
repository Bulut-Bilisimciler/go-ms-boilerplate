#!/bin/sh

go mod download

# install golang task runner
go get github.com/go-task/task/cmd/task 2>&1
# install golang task, mage and gopls
go install github.com/magefile/mage 2>&1
# go get -u -v golang.org/x/tools/cmd/gopls 2>&1

