package cli

import (
	"fmt"
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
	if dbPort == "" {
		emptyEnvVars = append(emptyEnvVars, "DB_PORT")
	}
	if len(emptyEnvVars) > 0 {
		return nil, fmt.Errorf("the following environment variables should be set: %s", strings.Join(emptyEnvVars, ", "))
	}

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

	unsavedPlayers, err := c.scraper.GetPlayers(unsavedPlayerIDs)
	if err != nil {
		return err
	}
	fmt.Println(unsavedPlayers)

	dbUnsavedPlayers := convertPlayers(unsavedPlayers)
	if err := c.dbClient.CreatePlayers(dbUnsavedPlayers); err != nil {
		return err
	}

	// dbにstatsを追加
	return nil
}

func convertPlayers(inputPlayers []npbweb.Player) []db.Player {
	var convertedPlayers []db.Player
	for _, player := range inputPlayers {
		convertedPlayers = append(convertedPlayers, convertPlayer(player))
	}
	return convertedPlayers
}

func convertPlayer(inputPlayer npbweb.Player) db.Player {
	return db.Player{
		ID:   inputPlayer.ID,
		Name: inputPlayer.Name,
		Kana: inputPlayer.Kana,
	}
}

func getEnv(envName, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		return defaultValue
	}
	return value
}
