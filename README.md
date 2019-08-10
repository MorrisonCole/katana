# 刀 (ka·ta·na) [![Build Status](https://travis-ci.org/MorrisonCole/katana.svg?branch=master)](https://travis-ci.org/MorrisonCole/katana)
A modern, open-source, English ⇄ Japanese dictionary with a native Android client. Written in Go & Kotlin.

## Structure

```
.
├── go (Go services)
├── android (Android client)
```

## Developer Setup

1. Follow the [Go quick start guide](https://grpc.io/docs/quickstart/go/).
2. Follow the [Google Cloud SDK quick start guide](https://cloud.google.com/sdk/docs/quickstarts).

## Go Dependencies

Run `go mod vendor` to reset the module's `vendor` directory (see: [make vendored copy of dependencies](https://tip.golang.org/cmd/go/#hdr-Make_vendored_copy_of_dependencies)).
