# 刀 (ka·ta·na) [![Build Status](https://travis-ci.org/MorrisonCole/katana.svg?branch=master)](https://travis-ci.org/MorrisonCole/katana)
A modern, open-source, English ⇄ Japanese dictionary with a native Android client. Written in Go & Kotlin.

## Structure

```
.
├── go (Go services)
├── android (Android client)
```

## Update Workspace Dependencies
```
bazel run //:gazelle -- update-repos -from_file=go.mod
```

## Build Go
```
bazel build go
```

## Architecture
![Architecture diagram](/architecture.png?raw=true "Architecture")
