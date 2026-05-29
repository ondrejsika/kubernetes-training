package main

import (
	"fmt"

	"github.com/ondrejsika/counter-frontend-go/pkg/server"
	"github.com/ondrejsika/counter-frontend-go/version"
)

var BUILD_ID string = "0"

func main() {
	version.Version = fmt.Sprintf("fe-%s", BUILD_ID)
	server.Server()
}
