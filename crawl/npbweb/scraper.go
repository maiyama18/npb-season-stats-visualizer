package npbweb

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

const idCol = 1

type Scraper struct {
	baseURL   string
	collector *colly.Collector
}

func NewScraper(baseURL string, randomDelay time.Duration) (*Scraper, error) {
	collector := colly.NewCollector()
	err := collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,
		RandomDelay: randomDelay,
	})
	if err != nil {
		return nil, err
	}

	return &Scraper{
		baseURL:   baseURL,
		collector: collector,
	}, nil
}

func (c *Scraper) GetTeamPitchers(teamID int) ([]PitcherStats, error) {
	url := fmt.Sprintf("%s/teams/%d/memberlist?type=p", c.baseURL, teamID)

	rows, err := c.scrapeRows(url)
	if err != nil {
		return nil, err
	}

	filteredRows := c.filterRows(rows)

	statsList, err := c.constructPitcherStatsList(filteredRows)
	if err != nil {
		return nil, err
	}

	return statsList, nil
}

func (c *Scraper) scrapeRows(url string) ([][]string, error) {
	var rows [][]string
	c.collector.OnHTML(`table.NpbPlSt > tbody > tr`, func(e *colly.HTMLElement) {
		var row []string
		e.ForEach(`td`, func(i int, tde *colly.HTMLElement) {
			// save player's ID instead of Name
			if i == idCol {
				u := strings.Trim(tde.ChildAttr(`a`, `href`), "/")
				su := strings.Split(u, "/")
				row = append(row, su[len(su)-1])
			} else {
				row = append(row, tde.Text)
			}
		})
		rows = append(rows, row)
	})

	if err := c.collector.Visit(url); err != nil {
		return nil, err
	}

	c.collector.OnHTMLDetach(`table.NpbPlSt > tbody > tr`)

	return rows, nil
}

func (c *Scraper) filterRows(rows [][]string) [][]string {
	var filtered [][]string
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		validCols := 0
		for _, col := range row[idCol+1:] {
			if col != "-" {
				validCols++
			}
		}
		if validCols == 0 {
			continue
		}

		filtered = append(filtered, row)
	}

	return filtered
}

func SelectUnsavedPlayerIDs(pitcherStatsList []PitcherStats, savedIDs []int) []int {
	savedIDSet := make(map[int]bool)
	for _, id := range savedIDs {
		savedIDSet[id] = true
	}

	var unsavedIDs []int
	for _, stats := range pitcherStatsList {
		if _, exists := savedIDSet[stats.playerID]; !exists {
			unsavedIDs = append(unsavedIDs, stats.playerID)
		}
	}

	return unsavedIDs
}

func parseIntCol(text string) *int {
	i, err := strconv.Atoi(text)
	if err != nil {
		return nil
	}
	return &i
}

func parseFloatCol(text string) *float64 {
	f, err := strconv.ParseFloat(text, 64)
	if err != nil {
		return nil
	}
	return &f
}
