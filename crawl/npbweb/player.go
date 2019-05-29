package npbweb

type Player struct {
	id   int
	name string
	kana string
}

//func (c *Scraper) getPlayers(rows [][]string) ([]Player, error) {
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
//func (c *Scraper) scrapePlayer(url, idStr string) (Player, error) {
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
//func (c *Scraper) extractNames(text string) (string, string) {
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
