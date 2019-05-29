package main

import (
	"fmt"
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
	c, err := cli.New()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] %s\n", err.Error())
		return exitCodeErr
	}

	if err := c.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "[ERROR] %s\n", err.Error())
		return exitCodeErr
	}

	return exitCodeOK
}
