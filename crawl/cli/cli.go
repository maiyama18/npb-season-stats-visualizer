package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mui87/npb-season-stats-visualizer/crawl/db"
	"github.com/mui87/npb-season-stats-visualizer/crawl/scraper"
)

const (
	defaultBaseURL    = "https://baseball.yahoo.co.jp/npb"
	defaultDelayMsStr = "2000"

	defaultDbUser   = "root"
	defaultDbHost   = "localhost"
	defaultDbPort   = "3306"
	defaultDbSchema = "npb_season_stats"
)

type CLI struct {
	dbClient *db.Client
	scraper  *scraper.Scraper
}

func New() (*CLI, error) {
	baseURL := getEnv("NPB_BASE_URL", defaultBaseURL)

	delayMsStr := getEnv("SCRAPE_DELAY_MS", defaultDelayMsStr)
	delayMs, err := strconv.Atoi(delayMsStr)
	if err != nil {
		return nil, fmt.Errorf("failed to get SCRAPE_DELAY_MS: %s", err)
	}

	s, err := scraper.New(baseURL, time.Duration(delayMs)*time.Millisecond)
	if err != nil {
		return nil, err
	}

	dbUser := getEnv("DB_USER", defaultDbUser)
	dbPassword := getEnv("DB_PASSWORD", "")
	if dbPassword == "" {
		return nil, errors.New("DB_PASSWORD not set")
	}
	dbHost := getEnv("DB_HOST", defaultDbHost)
	dbPort := getEnv("DB_PORT", defaultDbPort)
	dbSchema := getEnv("DB_SCHEMA", defaultDbSchema)

	c, err := db.NewClient(dbUser, dbPassword, dbHost, dbPort, dbSchema)
	if err != nil {
		return nil, fmt.Errorf("failed to create db client: %s", err)
	}

	return &CLI{
		scraper:  s,
		dbClient: c,
	}, nil
}

func (c *CLI) Run() error {
	defer c.dbClient.CloseDB()

	//for teamID := 1; teamID <= 12; teamID++ {
	//	pStatsList, err := s.GetTeamPitchers(teamID)
	//	if err != nil {
	//		return 1
	//	}
	//
	//	pitcherStatsList = append(pitcherStatsList, pStatsList...)
	//}

	// dbに存在するplayerを取得

	// statsListとdbに存在するplayerから、存在しないplayerがわかるので、存在しないplayerをscrape

	// dbに存在しないplayerを追加

	// dbにstatsを追加
	return nil
}

func getEnv(envName, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		return defaultValue
	}
	return value
}
