package main

import (
	"github.com/ueckoken/plarail2022/backend/external/internal"
	"log"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln("failed to initialize zap")
	}
	internal.Run(logger)
}
