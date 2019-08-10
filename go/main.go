package main

import (
	"com.morrisoncole/katana/ingest"
	"com.morrisoncole/katana/server"
)

func main() {
	ingest.Ingest()

	server.Connect()
	server.Serve()
}
