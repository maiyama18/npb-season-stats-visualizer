package main

import (
	"fmt"
	"github.com/mui87/npb-season-stats-visualizer/crawl/scraper"
	"os"
	"time"
)

const baseURL = "https://baseball.yahoo.co.jp/npb"

func main() {
	os.Exit(run())
}

func run() int {
	s, err := scraper.New(baseURL, 500*time.Millisecond)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	_, _ = s.GetTeamPitchers(6)
	//var (
	//	pitcherStatsList []crawler.PitcherStats
	//	batterStatsList []crawler.PitcherStats
	//)
	//for i := 1; i <= 12; i++ {
	//	_, _ = c.GetTeamPitchers(6)
	//}

	return 0
}
