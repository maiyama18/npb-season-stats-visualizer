package scraper

import (
	"fmt"
	"github.com/gocolly/colly"
	"math"
	"strconv"
	"strings"
	"time"
)

const idCol = 1

const (
	eraCol = iota + 2
	gameCol
	gameStartCol
	completeCol
	shutOutCol
	qualityStartCol
	winCol
	loseCol
	holdCol
	holdPointCol
	saveCol
	winPercentCol
	inningCol
	hitCol
	homeRunCol
	strikeOutCol
	strikeOutPercentCol
	walkCol
	hitByPitchCol
	wildPitchCol
	balkCol
	runCol
	earnedRunCol
	averageCol
	kbbCol
	whipCol
)

type Player struct {
	id   int
	name string
	kana string
}

type BatterStats struct {
	player Player
}

type PitcherStats struct {
	playerID         int
	era              *float64
	game             *int
	gameStart        *int
	complete         *int
	shutOut          *int
	qualityStart     *int
	win              *int
	lose             *int
	hold             *int
	holdPoint        *int
	save             *int
	winPercent       *float64
	inning           *float64
	hit              *int
	homeRun          *int
	strikeOut        *int
	strikeOutPercent *float64
	walk             *int
	hitByPitch       *int
	wildPitch        *int
	balk             *int
	run              *int
	earnedRun        *int
	average          *float64
	kbb              *float64
	whip             *float64
}

type Crawler struct {
	baseURL   string
	collector *colly.Collector
}

func New(baseURL string, randomDelay time.Duration) (*Crawler, error) {
	collector := colly.NewCollector()
	err := collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 1,
		RandomDelay: randomDelay,
	})
	if err != nil {
		return nil, err
	}

	return &Crawler{
		baseURL:   baseURL,
		collector: collector,
	}, nil
}

func (c *Crawler) TeamPitchers(teamID int) ([]PitcherStats, error) {
	url := fmt.Sprintf("%s/teams/%d/memberlist?type=p", c.baseURL, teamID)

	rows, err := c.scrapePitcherRows(url)
	if err != nil {
		return nil, err
	}

	filteredRows := c.filterRows(rows)

	statsList, err := c.constructPitcherStatsList(filteredRows)
	if err != nil {
		return nil, err
	}

	fmt.Println(len(statsList), "players")
	for _, stats := range statsList {
		fmt.Println("id", stats.playerID)
		fmt.Println("era", *stats.era)
		fmt.Println("win", *stats.win)
		fmt.Println("so", *stats.strikeOut)
		fmt.Println("inning", *stats.inning)
		fmt.Println()
	}
	return statsList, nil
}

func (c *Crawler) scrapePitcherRows(url string) ([][]string, error) {
	var rows [][]string
	c.collector.OnHTML(`table.NpbPlSt > tbody > tr`, func(e *colly.HTMLElement) {
		var row []string
		e.ForEach(`td`, func(i int, tde *colly.HTMLElement) {
			// save player's id instead of name
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

func (c *Crawler) filterRows(rows [][]string) [][]string {
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

//func (c *Crawler) getPlayers(rows [][]string) ([]Player, error) {
//	var players []Player
//	for _, row := range rows {
//		id := row[idCol]
//		u := fmt.Sprintf("%s/player/%s/", c.baseURL, id)
//
//		player, err := c.scrapePlayer(u, id)
//		if err != nil {
//			return nil, err
//		}
//		players = append(players, player)
//	}
//
//	return players, nil
//}
//
//func (c *Crawler) scrapePlayer(url, idStr string) (Player, error) {
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		return Player{}, err
//	}
//
//	var player Player
//	c.collector.OnHTML(`div.PlayerAdBox h1`, func(e *colly.HTMLElement) {
//		name, kana := c.extractNames(e.Text)
//		player = Player{
//			id: id,
//			name: name,
//			kana: kana,
//		}
//	})
//
//	if err := c.collector.Visit(url); err != nil {
//		return Player{}, err
//	}
//
//	c.collector.OnHTMLDetach(`div.PlayerAdBox h1`)
//
//	fmt.Println(player)
//	return player, nil
//}
//
//func (c *Crawler) extractNames(text string) (string, string) {
//	textRunes := []rune(text)
//	iOpen := runes.IndexRune(textRunes, '（')
//	iClose := runes.IndexRune(textRunes, '）')
//
//	name := textRunes[:iOpen]
//	kana := textRunes[iOpen+1:iClose]
//
//	var cleanedName, cleanedKana []rune
//	for _, c := range name {
//		if c != '　' && c != ' ' && c != '・' {
//			cleanedName = append(cleanedName, c)
//		}
//	}
//	for _, c := range kana {
//		if c != '　' && c != ' ' && c != '・' {
//			cleanedKana = append(cleanedKana, c)
//		}
//	}
//
//	return string(cleanedName), string(cleanedKana)
//}

func (c *Crawler) constructPitcherStatsList(rows [][]string) ([]PitcherStats, error) {
	var statsList []PitcherStats
	for _, row := range rows {
		stats, err := constructPitcherStats(row)
		if err != nil {
			return nil, err
		}

		statsList = append(statsList, stats)
	}

	return statsList, nil
}

func constructPitcherStats(row []string) (PitcherStats, error) {
	id, err := strconv.Atoi(row[idCol])
	if err != nil {
		return PitcherStats{}, err
	}

	inning := parseFloatCol(row[inningCol])
	if inning != nil {
		flr := math.Floor(*inning)
		*inning = flr + (*inning-flr)*(10.0/3.0)
	}

	return PitcherStats{
		playerID:         id,
		era:              parseFloatCol(row[eraCol]),
		game:             parseIntCol(row[gameCol]),
		gameStart:        parseIntCol(row[gameStartCol]),
		complete:         parseIntCol(row[completeCol]),
		shutOut:          parseIntCol(row[shutOutCol]),
		qualityStart:     parseIntCol(row[qualityStartCol]),
		win:              parseIntCol(row[winCol]),
		lose:             parseIntCol(row[loseCol]),
		hold:             parseIntCol(row[holdCol]),
		holdPoint:        parseIntCol(row[holdPointCol]),
		save:             parseIntCol(row[saveCol]),
		winPercent:       parseFloatCol(row[winPercentCol]),
		inning:           inning,
		hit:              parseIntCol(row[hitCol]),
		homeRun:          parseIntCol(row[homeRunCol]),
		strikeOut:        parseIntCol(row[strikeOutCol]),
		strikeOutPercent: parseFloatCol(row[strikeOutPercentCol]),
		walk:             parseIntCol(row[walkCol]),
		hitByPitch:       parseIntCol(row[hitByPitchCol]),
		wildPitch:        parseIntCol(row[wildPitchCol]),
		balk:             parseIntCol(row[balkCol]),
		run:              parseIntCol(row[runCol]),
		earnedRun:        parseIntCol(row[earnedRunCol]),
		average:          parseFloatCol(row[averageCol]),
		kbb:              parseFloatCol(row[kbbCol]),
		whip:             parseFloatCol(row[whipCol]),
	}, nil
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
