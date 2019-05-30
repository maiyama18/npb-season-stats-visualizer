package main

import (
	"fmt"
	"os"

	"github.com/mui87/npb-season-stats-visualizer/api/server"
)

const (
	exitCodeOK = iota
	exitCodeErr
)

func main() {
	os.Exit(run())
}

func run() int {
	s, err := server.New()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		return exitCodeErr
	}

	if err := s.Run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		return exitCodeErr
	}

	return exitCodeOK
}
