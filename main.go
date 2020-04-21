package main

import (
	"duitku-integration-api/internal/pkg/config"
	"duitku-integration-api/internal/pkg/httpwrap"
)

func main() {
	config.Init()

	h := httpwrap.New()

	// Init HTTP Route
	h.InitRoute()

	// Start API
	h.Start()
}
