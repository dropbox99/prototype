package main

import (
	"prototype/config"
	"prototype/lib/env"
)

func main() {
	cfg := config.NewConfig()

	cfg.Router.Run(env.String("MainSetup.ServerHost", "3000"))
}
