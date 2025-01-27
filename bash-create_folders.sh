#!/usr/bin/bash

folders=(
    "01-basics/01-variables-and-types"
    "01-basics/02-functions"
    "01-basics/03-control-structures"
    "01-basics/04-arrays-and-slices"
    "01-basics/05-maps"
    "01-basics/06-structs-and-methods"
    "02-intermediate/01-pointers"
    "02-intermediate/02-interfaces"
    "02-intermediate/03-concurrency/goroutines"
    "02-intermediate/04-concurrency/channels"
    "02-intermediate/05-concurrency/sync"
    "02-intermediate/06-error-handling"
    "03-advanced/01-file-handling"
    "03-advanced/02-testing"
    "03-advanced/03-generics"
    "03-advanced/04-reflection"
    "04-building-cli-tools/01-cobra"
    "04-building-cli-tools/02-flag"
    "04-building-cli-tools/03-custom-tools"
    "05-web-development/01-http-server"
    "05-web-development/02-rest-api"
    "05-web-development/03-gorilla-mux"
    "06-database/01-sql"
    "06-database/02-orm"
    "06-database/03-nosql"
    "07-deployment/01-building-binaries"
    "07-deployment/02-docker"
    "07-deployment/03-ci-cd"
    "07-deployment/04-versioning")

# Create all directories
for str in ${folders[@]}; do
    mkdir -p $str
done
