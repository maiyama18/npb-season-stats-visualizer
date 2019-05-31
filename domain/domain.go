package domain

import "time"

type TimeStamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Player struct {
	ID   int `gorm:"primary_key" sql:"type:int"`
	Name string
	Kana string
	TimeStamps
}

type Pitcher struct {
	Player
}

type Batter struct {
	Player
}

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
