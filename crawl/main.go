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

	var (
		pitcherStatsList []scraper.PitcherStats
	)

	for teamID := 1; teamID <= 12; teamID++ {
		pStats, err := s.GetTeamPitchers(teamID)
		if err != nil {
			return 1
		}

		pitcherStatsList = append(pitcherStatsList, pStats...)
	}

	// dbに存在するplayerを取得

	// statsListとdbに存在するplayerから、存在しないplayerがわかるので、存在しないplayerをscrape

	// dbに存在しないplayerを追加

	// dbにstatsを追加

	return 0
}
