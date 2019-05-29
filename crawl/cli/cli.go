package cli

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mui87/npb-season-stats-visualizer/crawl/db"
	"github.com/mui87/npb-season-stats-visualizer/crawl/npbweb"
)

const (
	defaultBaseURL    = "https://baseball.yahoo.co.jp/npb"
	defaultDelayMsStr = "2000"
)

type CLI struct {
	dbClient *db.Client
	scraper  *npbweb.Scraper

	logger *log.Logger
}

func New(logger *log.Logger) (*CLI, error) {
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

	logger.Println("web scraper initialized")

	var emptyEnvVars []string
	dbUser := getEnv("DB_USER", "")
	if dbUser == "" {
		emptyEnvVars = append(emptyEnvVars, "DB_USER")
	}
	dbPassword := getEnv("DB_PASSWORD", "")
	if dbPassword == "" {
		emptyEnvVars = append(emptyEnvVars, "DB_PASSWORD")
	}
	dbHost := getEnv("DB_HOST", "")
	if dbHost == "" {
		emptyEnvVars = append(emptyEnvVars, "DB_HOST")
	}
	dbPort := getEnv("DB_PORT", "")
	if dbPort == "" {
		emptyEnvVars = append(emptyEnvVars, "DB_PORT")
	}
	dbSchema := getEnv("DB_SCHEMA", "")
	if dbSchema == "" {
		emptyEnvVars = append(emptyEnvVars, "DB_SCHEMA")
	}
	if len(emptyEnvVars) > 0 {
		return nil, fmt.Errorf("the following environment variables should be set: %s", strings.Join(emptyEnvVars, ", "))
	}

	dbClient, err := db.NewClient(dbUser, dbPassword, dbHost, dbPort, dbSchema)
	if err != nil {
		return nil, fmt.Errorf("failed to create db client: %scraper", err)
	}

	logger.Println("db client initialized")

	dbClient.CreateTables()

	logger.Println("initialization complete")

	return &CLI{
		scraper:  scraper,
		dbClient: dbClient,
		logger:   logger,
	}, nil
}

func (c *CLI) Run() error {
	defer c.dbClient.CloseDB()

	c.logger.Println("fetching stats...")
	var pitcherStatsList []npbweb.PitcherStats
	for teamID := 1; teamID <= 2; teamID++ {
		pStatsList, err := c.scraper.GetTeamPitchers(teamID)
		if err != nil {
			return err
		}

		pitcherStatsList = append(pitcherStatsList, pStatsList...)
	}
	c.logger.Printf("complete! fetched %d stats", len(pitcherStatsList))

	savedPitcherIDs, err := c.dbClient.GetPitcherIDs()
	if err != nil {
		return err
	}

	unsavedPitcherIDs := npbweb.SelectUnsavedPlayerIDs(pitcherStatsList, savedPitcherIDs)
	c.logger.Printf("there is %d players not saved in db", len(unsavedPitcherIDs))

	if len(unsavedPitcherIDs) > 0 {
		c.logger.Println("fetching player's info...")
		unsavedPitchers, err := c.scraper.GetPlayers(unsavedPitcherIDs)
		if err != nil {
			return err
		}
		c.logger.Printf("complete! fetched %d players' info", len(unsavedPitchers))

		dbUnsavedPitchers := convertPitchers(unsavedPitchers)
		if err := c.dbClient.CreatePitchers(dbUnsavedPitchers); err != nil {
			return err
		}
		c.logger.Printf("saved %d players to db", len(dbUnsavedPitchers))
	}

	// dbにstatsを追加
	return nil
}

func convertPitchers(inputPlayers []npbweb.Player) []db.Pitcher {
	var convertedPlayers []db.Pitcher
	for _, player := range inputPlayers {
		convertedPlayers = append(convertedPlayers, convertPitcher(player))
	}
	return convertedPlayers
}

func convertPitcher(inputPlayer npbweb.Player) db.Pitcher {
	return db.Pitcher{
		Player: db.Player{
			ID:   inputPlayer.ID,
			Name: inputPlayer.Name,
			Kana: inputPlayer.Kana,
		},
	}
}

func getEnv(envName, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		return defaultValue
	}
	return value
}
