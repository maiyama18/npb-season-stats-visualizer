package cli

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mui87/npb-season-stats-visualizer/crawl/db"
	"github.com/mui87/npb-season-stats-visualizer/crawl/npbweb"
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
	scraper  *npbweb.Scraper
}

func New() (*CLI, error) {
	baseURL := getEnv("NPB_BASE_URL", defaultBaseURL)

	delayMsStr := getEnv("SCRAPE_DELAY_MS", defaultDelayMsStr)
	delayMs, err := strconv.Atoi(delayMsStr)
	if err != nil {
		return nil, fmt.Errorf("failed to get SCRAPE_DELAY_MS: %scraper", err)
	}

	scraper, err := npbweb.NewScraper(baseURL, time.Duration(delayMs)*time.Millisecond)
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

	dbClient, err := db.NewClient(dbUser, dbPassword, dbHost, dbPort, dbSchema)
	if err != nil {
		return nil, fmt.Errorf("failed to create db client: %scraper", err)
	}

	dbClient.CreateTables()

	return &CLI{
		scraper:  scraper,
		dbClient: dbClient,
	}, nil
}

func (c *CLI) Run() error {
	defer c.dbClient.CloseDB()

	var pitcherStatsList []npbweb.PitcherStats
	for teamID := 1; teamID <= 2; teamID++ {
		pStatsList, err := c.scraper.GetTeamPitchers(teamID)
		if err != nil {
			return err
		}

		pitcherStatsList = append(pitcherStatsList, pStatsList...)
	}

	for _, stats := range pitcherStatsList {
		fmt.Println(stats)
	}

	savedPlayerIDs, err := c.dbClient.GetPlayerIDs()
	if err != nil {
		return err
	}

	unsavedPlayerIDs := npbweb.SelectUnsavedPlayerIDs(pitcherStatsList, savedPlayerIDs)
	fmt.Println(len(savedPlayerIDs), len(unsavedPlayerIDs), len(pitcherStatsList))

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
