package main

import (
	"log"
	"os"

	"github.com/mui87/npb-season-stats-visualizer/crawl/cli"
)

const (
	exitCodeOK = iota
	exitCodeErr
)

func main() {
	os.Exit(run())
}

func run() int {
	logger := log.New(os.Stderr, "[LOG]", log.LstdFlags)
	c, err := cli.New(logger)
	if err != nil {
		logger.SetPrefix("[ERROR]")
		logger.Println(err.Error())
		return exitCodeErr
	}

	if err := c.Run(); err != nil {
		logger.SetPrefix("[ERROR]")
		logger.Println(err.Error())
		return exitCodeErr
	}

	return exitCodeOK
}
