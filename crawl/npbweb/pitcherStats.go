package npbweb

import (
	"math"
	"strconv"
)

const (
	pitcherEraCol = iota + 2
	pitcherGameCol
	pitcherGameStartCol
	pitcherCompleteCol
	pitcherShutOutCol
	pitcherQualityStartCol
	pitcherWinCol
	pitcherLoseCol
	pitcherHoldCol
	pitcherHoldPointCol
	pitcherSaveCol
	pitcherWinPercentCol
	pitcherInningCol
	pitcherHitCol
	pitcherHomeRunCol
	pitcherStrikeOutCol
	pitcherStrikeOutPercentCol
	pitcherWalkCol
	pitcherHitByPitchCol
	pitcherWildPitchCol
	pitcherBalkCol
	pitcherRunCol
	pitcherEarnedRunCol
	pitcherAverageCol
	pitcherKbbCol
	pitcherWhipCol
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

	inning := parseFloatCol(row, pitcherInningCol)
	if inning != nil {
		flr := math.Floor(*inning)
		*inning = flr + (*inning-flr)*(10.0/3.0)
	}

	return PitcherStats{
		PlayerID:         id,
		Era:              parseFloatCol(row, pitcherEraCol),
		Game:             parseIntCol(row, pitcherGameCol),
		GameStart:        parseIntCol(row, pitcherGameStartCol),
		Complete:         parseIntCol(row, pitcherCompleteCol),
		ShutOut:          parseIntCol(row, pitcherShutOutCol),
		QualityStart:     parseIntCol(row, pitcherQualityStartCol),
		Win:              parseIntCol(row, pitcherWinCol),
		Lose:             parseIntCol(row, pitcherLoseCol),
		Hold:             parseIntCol(row, pitcherHoldCol),
		HoldPoint:        parseIntCol(row, pitcherHoldPointCol),
		Save:             parseIntCol(row, pitcherSaveCol),
		WinPercent:       parseFloatCol(row, pitcherWinPercentCol),
		Inning:           inning,
		Hit:              parseIntCol(row, pitcherHitCol),
		HomeRun:          parseIntCol(row, pitcherHomeRunCol),
		StrikeOut:        parseIntCol(row, pitcherStrikeOutCol),
		StrikeOutPercent: parseFloatCol(row, pitcherStrikeOutPercentCol),
		Walk:             parseIntCol(row, pitcherWalkCol),
		HitByPitch:       parseIntCol(row, pitcherHitByPitchCol),
		WildPitch:        parseIntCol(row, pitcherWildPitchCol),
		Balk:             parseIntCol(row, pitcherBalkCol),
		Run:              parseIntCol(row, pitcherRunCol),
		EarnedRun:        parseIntCol(row, pitcherEarnedRunCol),
		Average:          parseFloatCol(row, pitcherAverageCol),
		Kbb:              parseFloatCol(row, pitcherKbbCol),
		Whip:             parseFloatCol(row, pitcherWhipCol),
	}, nil
}
