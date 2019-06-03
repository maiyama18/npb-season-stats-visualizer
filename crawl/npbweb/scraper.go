package npbweb

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

const idCol = 1

type Scraper struct {
	baseURL       string
	maxRetryCount int

	visitedURLs map[string]bool

	collector *colly.Collector
	logger    *log.Logger
}

func NewScraper(baseURL string, randomDelay time.Duration, maxRetryCount int, logger *log.Logger) (*Scraper, error) {
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
		baseURL:       baseURL,
		maxRetryCount: maxRetryCount,
		visitedURLs:   make(map[string]bool),
		collector:     collector,
		logger:        logger,
	}, nil
}

func (c *Scraper) GetTeamPitcherStatsList(teamID int) ([]PitcherStats, error) {
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

func (c *Scraper) GetTeamBatterStatsList(teamID int) ([]BatterStats, error) {
	url := fmt.Sprintf("%s/teams/%d/memberlist?type=b", c.baseURL, teamID)

	rows, err := c.scrapeRows(url)
	if err != nil {
		return nil, err
	}

	filteredRows := c.filterRows(rows)

	statsList, err := c.constructBatterStatsList(filteredRows)
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

	reqErr := c.visitWithRetry(url)

	c.collector.OnHTMLDetach(`table.NpbPlSt > tbody > tr`)

	if reqErr != nil {
		return nil, reqErr
	} else {
		return rows, nil
	}
}

func (c *Scraper) visitWithRetry(url string) error {
	var reqErr error

	retryCount := 0
	c.collector.OnError(func(r *colly.Response, err error) {
		if _, exist := c.visitedURLs[url]; exist {
			return
		}

		retryCount++
		if retryCount <= 3 {
			c.logger.Printf("request to %s failed: %s. retry: %d\n", url, err, retryCount)
			_ = r.Request.Retry()
		} else {
			reqErr = err
		}
	})

	_ = c.collector.Visit(url)

	c.visitedURLs[url] = true

	return reqErr
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

func parseIntCol(row []string, col int) *int {
	if col < 0 || col >= len(row) {
		return nil
	}

	i, err := strconv.Atoi(row[col])
	if err != nil {
		return nil
	}
	return &i
}

func parseFloatCol(row []string, col int) *float64 {
	if col < 0 || col >= len(row) {
		return nil
	}

	f, err := strconv.ParseFloat(row[col], 64)
	if err != nil {
		return nil
	}
	return &f
}
