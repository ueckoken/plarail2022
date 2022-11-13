package main

import (
	"log"
	"ueckoken/plarail2022-external/internal"

	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalln("failed to initialize zap")
	}
	internal.Run(logger)
}
