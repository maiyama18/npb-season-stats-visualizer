package main

import (
	"fmt"
	"os"

	"github.com/mui87/npb-season-stats-visualizer/api/infrastructure/router"
)

const (
	exitCodeOK = iota
	exitCodeErr
)

func main() {
	os.Exit(run())
}

func run() int {
	if err := router.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, err.Error())
		return exitCodeErr
	}
	return exitCodeOK
}
