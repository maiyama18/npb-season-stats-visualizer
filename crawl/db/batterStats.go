package db

import "time"

type BatterStats struct {
	Game                       *int `gorm:"primary_key" sql:"type:int"`
	BatterID                   int  `gorm:"primary_key" sql:"type:int"`
	Batter                     Batter
	Date                       time.Time `sql:"not null;type:date"`
	Average                    *float64
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
	TimeStamps
}

func (BatterStats) TableName() string {
	return "batter_stats_list"
}
