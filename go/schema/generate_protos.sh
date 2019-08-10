#!/bin/sh

protoc -I ./ ./dictionary.proto --go_out=plugins=grpc:.
