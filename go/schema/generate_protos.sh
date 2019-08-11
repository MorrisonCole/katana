#!/bin/sh

protoc -I ../../android/app/src/main/proto/ ../../android/app/src/main/proto/dictionary.proto --go_out=plugins=grpc:.
