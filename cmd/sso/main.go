package main

import (
	"fmt"
	"sso/internal/config"
)

func main() {

	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger (slog)

	// TODO: init app

	// TODO: run gRPC server

}
