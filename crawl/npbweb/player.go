package npbweb

import (
	"fmt"

	"github.com/gobwas/glob/util/runes"
	"github.com/gocolly/colly"
)

type Player struct {
	ID   int
	Name string
	Kana string
}

func (c *Scraper) GetPlayer(id int) (Player, error) {
	u := fmt.Sprintf("%s/player/%d/", c.baseURL, id)

	player, err := c.scrapePlayer(u)
	if err != nil {
		return Player{}, err
	}
	player.ID = id

	return player, nil
}

func (c *Scraper) scrapePlayer(url string) (Player, error) {
	var player Player
	c.collector.OnHTML(`div.PlayerAdBox h1`, func(e *colly.HTMLElement) {
		name, kana := extractNames(e.Text)
		player = Player{
			Name: name,
			Kana: kana,
		}
	})

	if err := c.collector.Visit(url); err != nil {
		return Player{}, err
	}

	c.collector.OnHTMLDetach(`div.PlayerAdBox h1`)

	return player, nil
}

func extractNames(text string) (string, string) {
	textRunes := []rune(text)
	iOpen := runes.IndexRune(textRunes, '（')
	iClose := runes.IndexRune(textRunes, '）')

	name := textRunes[:iOpen]
	kana := textRunes[iOpen+1 : iClose]

	var cleanedName, cleanedKana []rune
	for _, c := range name {
		if c != '　' && c != ' ' && c != '・' {
			cleanedName = append(cleanedName, c)
		}
	}
	for _, c := range kana {
		if c != '　' && c != ' ' && c != '・' {
			cleanedKana = append(cleanedKana, c)
		}
	}

	return string(cleanedName), string(cleanedKana)
}
