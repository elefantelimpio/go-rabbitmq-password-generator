#!/bin/bash

runDevEnv() {
## -race and coverage use it for local testing
 go test ./... -short -race -coverprofile coverreport.out
 go tool cover -html coverreport.out
 go tool cover -func coverreport.out | grep total:
}

if [ $# -eq 1 ]
then
    if [ "$1" == "prod" ]
    then
    ## use it for production
    go test ./... -race -coverprofile coverage.out -json > report.json
    else
        runDevEnv
    fi
else
    runDevEnv
fi

