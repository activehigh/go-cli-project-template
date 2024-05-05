#!/usr/bin/env just --justfile

jd := justfile_directory()

vet:
    go vet ./...

fmt:
    go fmt ./...

tidy:
    go mod tidy

generate:
    go generate ./...

docker-generate: tidy vet fmt
    docker compose up --build go-generate

docker-build: vet fmt tidy
    docker compose up --build service

docker-test: generate
    docker compose up --build service-test

test *path:
    ginkgo {{path}}/...

start-dd-agent:
    docker compose up --build dd-agent
