package npbweb

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
	PlayerID         int
	Era              *float64
	Game             *int
	GameStart        *int
	Complete         *int
	ShutOut          *int
	QualityStart     *int
	Win              *int
	Lose             *int
	Hold             *int
	HoldPoint        *int
	Save             *int
	WinPercent       *float64
	Inning           *float64
	Hit              *int
	HomeRun          *int
	StrikeOut        *int
	StrikeOutPercent *float64
	Walk             *int
	HitByPitch       *int
	WildPitch        *int
	Balk             *int
	Run              *int
	EarnedRun        *int
	Average          *float64
	Kbb              *float64
	Whip             *float64
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
		PlayerID:         id,
		Era:              parseFloatCol(row[eraCol]),
		Game:             parseIntCol(row[gameCol]),
		GameStart:        parseIntCol(row[gameStartCol]),
		Complete:         parseIntCol(row[completeCol]),
		ShutOut:          parseIntCol(row[shutOutCol]),
		QualityStart:     parseIntCol(row[qualityStartCol]),
		Win:              parseIntCol(row[winCol]),
		Lose:             parseIntCol(row[loseCol]),
		Hold:             parseIntCol(row[holdCol]),
		HoldPoint:        parseIntCol(row[holdPointCol]),
		Save:             parseIntCol(row[saveCol]),
		WinPercent:       parseFloatCol(row[winPercentCol]),
		Inning:           inning,
		Hit:              parseIntCol(row[hitCol]),
		HomeRun:          parseIntCol(row[homeRunCol]),
		StrikeOut:        parseIntCol(row[strikeOutCol]),
		StrikeOutPercent: parseFloatCol(row[strikeOutPercentCol]),
		Walk:             parseIntCol(row[walkCol]),
		HitByPitch:       parseIntCol(row[hitByPitchCol]),
		WildPitch:        parseIntCol(row[wildPitchCol]),
		Balk:             parseIntCol(row[balkCol]),
		Run:              parseIntCol(row[runCol]),
		EarnedRun:        parseIntCol(row[earnedRunCol]),
		Average:          parseFloatCol(row[averageCol]),
		Kbb:              parseFloatCol(row[kbbCol]),
		Whip:             parseFloatCol(row[whipCol]),
	}, nil
}
