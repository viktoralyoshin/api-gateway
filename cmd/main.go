package main

import (
	"api-gateway/internal/app"

	"github.com/viktoralyoshin/utils/pkg/logger"
)

func main() {
	logger.Setup()

	app.Start()
}
