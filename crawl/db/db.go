package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

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

type Client struct {
	db *gorm.DB
}

func NewClient(dbUser, dbPassword, dbHost, dbPort, dbSchema string, logMode bool) (*Client, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbSchema)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	db.LogMode(logMode)

	return &Client{
		db: db,
	}, nil
}

func (c *Client) CreateTables() error {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	objects := []interface{}{&Pitcher{}, &Batter{}, &PitcherStats{}, &BatterStats{}}
	for _, obj := range objects {
		if !c.db.HasTable(obj) {
			if err := tx.CreateTable(obj).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}

func (c *Client) GetPlayerIDs() ([]int, []int, error) {
	pitcherIDs, err := c.getPitcherIDs()
	if err != nil {
		return nil, nil, err
	}
	batterIDs, err := c.getBatterIDs()
	if err != nil {
		return nil, nil, err
	}

	return pitcherIDs, batterIDs, nil
}

func (c *Client) CreatePlayers(pitchers []Pitcher, batters []Batter) error {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, pitcher := range pitchers {
		if err := tx.Create(&pitcher).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, batter := range batters {
		if err := tx.Create(&batter).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

func (c *Client) CreateStatsList(pitcherStatsList []PitcherStats, batterStatsList []BatterStats) (int, int, error) {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return 0, 0, err
	}

	savedPitcherStatsCount := 0
	savedBatterStatsCount := 0

	for _, pStats := range pitcherStatsList {
		exists, err := c.doesPitcherStatsExists(tx, pStats.PitcherID, pStats.Game)
		if err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if exists {
			continue
		}

		// here, Save is used instead of Create because there may be the existing record which was softly deleted
		if err := tx.Save(&pStats).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		savedPitcherStatsCount++
	}

	for _, bStats := range batterStatsList {
		exists, err := c.doesBatterStatsExists(tx, bStats.BatterID, bStats.Game)
		if err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		if exists {
			continue
		}

		if err := tx.Save(&bStats).Error; err != nil {
			tx.Rollback()
			return 0, 0, err
		}
		savedBatterStatsCount++
	}

	return savedPitcherStatsCount, savedBatterStatsCount, tx.Commit().Error
}

func (c *Client) CloseDB() {
	_ = c.db.Close()
}

func (c *Client) getPitcherIDs() ([]int, error) {
	var pitchers []Pitcher
	if err := c.db.Find(&pitchers).Error; err != nil {
		return nil, err
	}

	var ids []int
	for _, pitcher := range pitchers {
		ids = append(ids, pitcher.ID)
	}

	return ids, nil
}

func (c *Client) getBatterIDs() ([]int, error) {
	var batters []Batter
	if err := c.db.Find(&batters).Error; err != nil {
		return nil, err
	}

	var ids []int
	for _, batter := range batters {
		ids = append(ids, batter.ID)
	}

	return ids, nil
}

func (c *Client) doesPitcherStatsExists(tx *gorm.DB, pitcherID int, game *int) (bool, error) {
	var existings []PitcherStats
	if err := tx.Where(&PitcherStats{PitcherID: pitcherID, Game: game}).Find(&existings).Error; err != nil {
		return false, err
	}

	return len(existings) > 0, nil
}

func (c *Client) doesBatterStatsExists(tx *gorm.DB, batterID int, game *int) (bool, error) {
	var existings []BatterStats
	if err := tx.Where(&BatterStats{BatterID: batterID, Game: game}).Find(&existings).Error; err != nil {
		return false, err
	}

	return len(existings) > 0, nil
}
