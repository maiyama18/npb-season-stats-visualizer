package cli

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mui87/npb-season-stats-visualizer/domain"

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

	logModeStr := getEnv("GORM_LOG_MODE", "off")
	logMode := false
	if logModeStr == "on" {
		logMode = true
	}

	dbClient, err := db.NewClient(dbUser, dbPassword, dbHost, dbPort, dbSchema, logMode)
	if err != nil {
		return nil, fmt.Errorf("failed to create db client: %scraper", err)
	}

	logger.Println("db client initialized")

	if err := dbClient.CreateTables(); err != nil {
		return nil, err
	}

	logger.Println("initialization complete")

	return &CLI{
		scraper:  scraper,
		dbClient: dbClient,
		logger:   logger,
	}, nil
}

func (c *CLI) Run() error {
	defer c.dbClient.CloseDB()

	pitcherStatsList, batterStatsList, err := c.fetchStatsLists()
	if err != nil {
		return err
	}

	unsavedPitcherIDs, unsavedBatterIDs, err := c.getUnsavedPlayerIDs(pitcherStatsList, batterStatsList)
	if err != nil {
		return err
	}

	unsavedPitchers, unsavedBatters, err := c.fetchUnsavedPlayers(unsavedPitcherIDs, unsavedBatterIDs)
	if err != nil {
		return err
	}

	dbUnsavedPitchers, dbUnsavedBatters := convertPlayers(unsavedPitchers, unsavedBatters)
	if err := c.saveUnsavedPlayers(dbUnsavedPitchers, dbUnsavedBatters); err != nil {
		return err
	}

	dbPitcherStatsList, dbBatterStatsList := convertStatsLists(pitcherStatsList, batterStatsList)
	if err := c.saveStatsLists(dbPitcherStatsList, dbBatterStatsList); err != nil {
		return err
	}

	return nil
}

func (c *CLI) fetchStatsLists() ([]npbweb.PitcherStats, []npbweb.BatterStats, error) {
	c.logger.Println("fetching stats...")

	var (
		pitcherStatsList []npbweb.PitcherStats
		batterStatsList  []npbweb.BatterStats
	)
	for teamID := 1; teamID <= 12; teamID++ {
		pStatsList, err := c.scraper.GetTeamPitcherStatsList(teamID)
		if err != nil {
			return nil, nil, err
		}
		bStatsList, err := c.scraper.GetTeamBatterStatsList(teamID)
		if err != nil {
			return nil, nil, err
		}

		pitcherStatsList = append(pitcherStatsList, pStatsList...)
		batterStatsList = append(batterStatsList, bStatsList...)
	}

	c.logger.Printf("complete! fetched %d pitching stats and %d batting stats", len(pitcherStatsList), len(batterStatsList))

	return pitcherStatsList, batterStatsList, nil
}

func (c *CLI) getUnsavedPlayerIDs(
	pitcherStatsList []npbweb.PitcherStats, batterStatsList []npbweb.BatterStats,
) ([]int, []int, error) {
	savedPitcherIDs, savedBatterIDs, err := c.dbClient.GetPlayerIDs()
	if err != nil {
		return nil, nil, err
	}

	var (
		remotePitcherIDs []int
		remoteBatterIDs  []int
	)
	for _, pStats := range pitcherStatsList {
		remotePitcherIDs = append(remotePitcherIDs, pStats.PlayerID)
	}
	for _, bStats := range batterStatsList {
		remoteBatterIDs = append(remoteBatterIDs, bStats.PlayerID)
	}

	unsavedPitcherIDs := selectUnsavedPlayerIDs(remotePitcherIDs, savedPitcherIDs)
	unsavedBatterIDs := selectUnsavedPlayerIDs(remoteBatterIDs, savedBatterIDs)

	c.logger.Printf("there are %d pitchers and %d batters not saved in db", len(unsavedPitcherIDs), len(unsavedBatterIDs))

	return unsavedPitcherIDs, unsavedBatterIDs, nil
}

func (c *CLI) fetchUnsavedPlayers(pitcherIDs, batterIDs []int) ([]npbweb.Player, []npbweb.Player, error) {
	c.logger.Println("fetching players' profiles...")

	pitchers, err := c.scraper.GetPlayers(pitcherIDs)
	if err != nil {
		return nil, nil, err
	}

	batters, err := c.scraper.GetPlayers(batterIDs)
	if err != nil {
		return nil, nil, err
	}

	c.logger.Printf("complete! fetched %d pitcher profiles and %d batter profiles", len(pitchers), len(batters))

	return pitchers, batters, nil
}

func (c *CLI) saveUnsavedPlayers(pitchers []domain.Pitcher, batters []domain.Batter) error {
	if err := c.dbClient.CreatePlayers(pitchers, batters); err != nil {
		return err
	}

	c.logger.Printf("saved %d pitcher profiles and %d batter profiles", len(pitchers), len(batters))

	return nil
}

func (c *CLI) saveStatsLists(pitcherStatsList []domain.PitcherStats, batterStatsList []domain.BatterStats) error {
	savedPCount, savedBCount, err := c.dbClient.CreateStatsList(pitcherStatsList, batterStatsList)
	if err != nil {
		return err
	}

	c.logger.Printf("saved %d pitcher stats and %d batter stats", savedPCount, savedBCount)

	return nil
}

func selectUnsavedPlayerIDs(remoteIDs, savedIDs []int) []int {
	savedIDSet := make(map[int]bool)
	for _, id := range savedIDs {
		savedIDSet[id] = true
	}

	var unsavedIDs []int
	for _, remoteID := range remoteIDs {
		if _, exists := savedIDSet[remoteID]; !exists {
			unsavedIDs = append(unsavedIDs, remoteID)
		}
	}

	return unsavedIDs
}

func convertPlayers(inputPitchers, inputBatters []npbweb.Player) ([]domain.Pitcher, []domain.Batter) {
	var (
		convertedPitchers []domain.Pitcher
		convertedBatters  []domain.Batter
	)

	for _, p := range inputPitchers {
		convertedPitchers = append(convertedPitchers, convertPitcher(p))
	}
	for _, b := range inputBatters {
		convertedBatters = append(convertedBatters, convertBatter(b))
	}

	return convertedPitchers, convertedBatters
}

func convertPitcher(inputPlayer npbweb.Player) domain.Pitcher {
	return domain.Pitcher{
		Player: domain.Player{
			ID:   inputPlayer.ID,
			Name: inputPlayer.Name,
			Kana: inputPlayer.Kana,
		},
	}
}

func convertBatter(inputPlayer npbweb.Player) domain.Batter {
	return domain.Batter{
		Player: domain.Player{
			ID:   inputPlayer.ID,
			Name: inputPlayer.Name,
			Kana: inputPlayer.Kana,
		},
	}
}

func convertStatsLists(
	inputPitcherStatsList []npbweb.PitcherStats, inputBatterStatsList []npbweb.BatterStats,
) ([]domain.PitcherStats, []domain.BatterStats) {
	var (
		convertedPitcherStatsList []domain.PitcherStats
		convertedBatterStatsList  []domain.BatterStats
	)

	for _, p := range inputPitcherStatsList {
		convertedPitcherStatsList = append(convertedPitcherStatsList, convertPitcherStats(p))
	}
	for _, b := range inputBatterStatsList {
		convertedBatterStatsList = append(convertedBatterStatsList, convertBatterStats(b))
	}

	return convertedPitcherStatsList, convertedBatterStatsList
}

func convertPitcherStats(inputStats npbweb.PitcherStats) domain.PitcherStats {
	// date is the previous day of the scraping
	date := time.Now().Add(-24 * time.Hour)
	return domain.PitcherStats{
		PitcherID:        inputStats.PlayerID,
		Date:             date,
		Era:              inputStats.Era,
		Game:             inputStats.Game,
		GameStart:        inputStats.GameStart,
		Complete:         inputStats.Complete,
		ShutOut:          inputStats.ShutOut,
		QualityStart:     inputStats.QualityStart,
		Win:              inputStats.Win,
		Lose:             inputStats.Lose,
		Hold:             inputStats.Hold,
		HoldPoint:        inputStats.HoldPoint,
		Save:             inputStats.Save,
		WinPercent:       inputStats.WinPercent,
		Inning:           inputStats.Inning,
		Hit:              inputStats.Hit,
		HomeRun:          inputStats.HomeRun,
		StrikeOut:        inputStats.StrikeOut,
		StrikeOutPercent: inputStats.StrikeOutPercent,
		Walk:             inputStats.Walk,
		HitByPitch:       inputStats.HitByPitch,
		WildPitch:        inputStats.WildPitch,
		Balk:             inputStats.Balk,
		Run:              inputStats.Run,
		EarnedRun:        inputStats.EarnedRun,
		Average:          inputStats.Average,
		Kbb:              inputStats.Kbb,
		Whip:             inputStats.Whip,
	}
}

func convertBatterStats(inputStats npbweb.BatterStats) domain.BatterStats {
	// date is the previous day of the scraping
	date := time.Now().Add(-24 * time.Hour)
	return domain.BatterStats{
		BatterID:                   inputStats.PlayerID,
		Date:                       date,
		Average:                    inputStats.Average,
		Game:                       inputStats.Game,
		PlateAppearance:            inputStats.PlateAppearance,
		AtBat:                      inputStats.AtBat,
		Hit:                        inputStats.Hit,
		Double:                     inputStats.Double,
		Triple:                     inputStats.Triple,
		HomeRun:                    inputStats.HomeRun,
		TotalBase:                  inputStats.TotalBase,
		RunBattedIn:                inputStats.RunBattedIn,
		Run:                        inputStats.Run,
		StrikeOut:                  inputStats.StrikeOut,
		Walk:                       inputStats.Walk,
		HitByPitch:                 inputStats.HitByPitch,
		Sacrifice:                  inputStats.Sacrifice,
		SacrificeFly:               inputStats.SacrificeFly,
		StolenBase:                 inputStats.StolenBase,
		CaughtStealing:             inputStats.CaughtStealing,
		DoublePlay:                 inputStats.DoublePlay,
		OnBasePercent:              inputStats.OnBasePercent,
		SluggingPercent:            inputStats.SluggingPercent,
		Ops:                        inputStats.Ops,
		AverageWithScoringPosition: inputStats.AverageWithScoringPosition,
		Error:                      inputStats.Error,
	}
}

func getEnv(envName, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		return defaultValue
	}
	return value
}
