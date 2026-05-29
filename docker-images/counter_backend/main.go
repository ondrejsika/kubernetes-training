package main

import (
	"fmt"
	"os"

	"github.com/ondrejsika/counter/pkg/server"
	"github.com/ondrejsika/counter/version"
)

var BUILD_ID string = "0"

func main() {
	setDefault("API_ONLY", "1")
	version.Version = fmt.Sprintf("be-%s", BUILD_ID)
	server.Server(server.ServerOptions{})
}

func setDefault(key, value string) {
	if os.Getenv(key) == "" {
		os.Setenv(key, value)
	}
}
