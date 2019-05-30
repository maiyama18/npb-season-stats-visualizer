package db

import "time"

type PitcherStats struct {
	Game             *int `gorm:"primary_key sql:"type:int"`
	PitcherID        int  `gorm:"primary_key sql:"type:int"`
	Pitcher          Pitcher
	Date             time.Time `sql:"not null;type:date"`
	Era              *float64
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
	TimeStamps
}

func (PitcherStats) TableName() string {
	return "pitcher_stats_list"
}
