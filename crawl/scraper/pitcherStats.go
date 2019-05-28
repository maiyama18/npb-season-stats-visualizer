package scraper

import (
	"math"
	"strconv"
)

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

func (c *Scraper) constructPitcherStatsList(rows [][]string) ([]PitcherStats, error) {
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
