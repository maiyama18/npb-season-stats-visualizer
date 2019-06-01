package main

import (
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mui87/npb-season-stats-visualizer/crawl/cli"
)

func Run() error {
	logger := log.New(os.Stderr, "[LOG]", log.LstdFlags)
	c, err := cli.New(logger)
	if err != nil {
		logger.SetPrefix("[ERROR]")
		logger.Println(err.Error())
		return err
	}

	if err := c.Run(); err != nil {
		logger.SetPrefix("[ERROR]")
		logger.Println(err.Error())
		return err
	}

	return nil
}

func main() {
	lambda.Start(Run)
}
