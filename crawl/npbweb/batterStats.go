package npbweb

import (
	"strconv"
)

const (
	batterAverageCol = iota + 2
	batterGameCol
	batterPlateAppearanceCol
	batterAtBatCol
	batterHitCol
	batterDoubleCol
	batterTripleCol
	batterHomeRunCol
	batterTotalBaseCol
	batterRunBattedInCol
	batterRunCol
	batterStrikeOutCol
	batterWalkCol
	batterHitByPitchCol
	batterSacrificeCol
	batterSacrificeFlyCol
	batterStolenBaseCol
	batterCaughtStealingCol
	batterDoublePlayCol
	batterOnBasePercentCol
	batterSluggingPercentCol
	batterOpsCol
	batterAverageWithScoringPositionCol
	batterErrorCol
)

type BatterStats struct {
	PlayerID                   int
	Average                    *float64
	Game                       *int
	PlateAppearance            *int
	AtBat                      *int
	Hit                        *int
	Double                     *int
	Triple                     *int
	HomeRun                    *int
	TotalBase                  *int
	RunBattedIn                *int
	Run                        *int
	StrikeOut                  *int
	Walk                       *int
	HitByPitch                 *int
	Sacrifice                  *int
	SacrificeFly               *int
	StolenBase                 *int
	CaughtStealing             *int
	DoublePlay                 *int
	OnBasePercent              *float64
	SluggingPercent            *float64
	Ops                        *float64
	AverageWithScoringPosition *float64
	Error                      *int
}

func (c *Scraper) constructBatterStatsList(rows [][]string) ([]BatterStats, error) {
	var statsList []BatterStats
	for _, row := range rows {
		stats, err := constructBatterStats(row)
		if err != nil {
			return nil, err
		}

		statsList = append(statsList, stats)
	}

	return statsList, nil
}

func constructBatterStats(row []string) (BatterStats, error) {
	id, err := strconv.Atoi(row[idCol])
	if err != nil {
		return BatterStats{}, err
	}

	return BatterStats{
		PlayerID:                   id,
		Average:                    parseFloatCol(row, batterAverageCol),
		Game:                       parseIntCol(row, batterGameCol),
		PlateAppearance:            parseIntCol(row, batterPlateAppearanceCol),
		AtBat:                      parseIntCol(row, batterAtBatCol),
		Hit:                        parseIntCol(row, batterHitCol),
		Double:                     parseIntCol(row, batterDoubleCol),
		Triple:                     parseIntCol(row, batterTripleCol),
		HomeRun:                    parseIntCol(row, batterHomeRunCol),
		TotalBase:                  parseIntCol(row, batterTotalBaseCol),
		RunBattedIn:                parseIntCol(row, batterRunBattedInCol),
		Run:                        parseIntCol(row, batterRunCol),
		StrikeOut:                  parseIntCol(row, batterStrikeOutCol),
		Walk:                       parseIntCol(row, batterWalkCol),
		HitByPitch:                 parseIntCol(row, batterHitByPitchCol),
		Sacrifice:                  parseIntCol(row, batterSacrificeCol),
		SacrificeFly:               parseIntCol(row, batterSacrificeFlyCol),
		StolenBase:                 parseIntCol(row, batterStolenBaseCol),
		CaughtStealing:             parseIntCol(row, batterCaughtStealingCol),
		DoublePlay:                 parseIntCol(row, batterDoublePlayCol),
		OnBasePercent:              parseFloatCol(row, batterOnBasePercentCol),
		SluggingPercent:            parseFloatCol(row, batterSluggingPercentCol),
		Ops:                        parseFloatCol(row, batterOpsCol),
		AverageWithScoringPosition: parseFloatCol(row, batterAverageWithScoringPositionCol),
		Error:                      parseIntCol(row, batterErrorCol),
	}, nil
}
