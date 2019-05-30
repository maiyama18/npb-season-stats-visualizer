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
	DeletedAt *time.Time
}

type Player struct {
	ID   int    `gorm:"primary_key"`
	Name string `sql:"not null"`
	Kana string `sql:"not null"`
	TimeStamps
}

type Pitcher struct {
	Player
}

type Batter struct {
	Player
}

type PitcherStats struct {
	Pitcher          Pitcher `gorm:"foreignkey:PitcherID"`
	PitcherID        int
	Date             time.Time `sql:"not null;type:date"`
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
	TimeStamps
}

func (PitcherStats) TableName() string {
	return "pitcher_stats_list"
}

type BatterStats struct {
	Batter                     Batter `gorm:"foreignkey:BatterID"`
	BatterID                   int
	Date                       time.Time `sql:"not null;type:date"`
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
	TimeStamps
}

func (BatterStats) TableName() string {
	return "batter_stats_list"
}

type Client struct {
	db *gorm.DB
}

func NewClient(dbUser, dbPassword, dbHost, dbPort, dbSchema string) (*Client, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbSchema)
	db, err := gorm.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	return &Client{
		db: db,
	}, nil
}

func (c *Client) CreateTables() {
	objects := []interface{}{&Pitcher{}, &Batter{}, &PitcherStats{}, &BatterStats{}}

	for _, obj := range objects {
		if !c.db.HasTable(obj) {
			c.db.CreateTable(obj)
		}
	}
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

func (c *Client) CreateStatsList(pitcherStatsList []PitcherStats, batterStatsList []BatterStats) error {
	tx := c.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, pStats := range pitcherStatsList {
		if err := tx.Create(&pStats).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	for _, bStats := range batterStatsList {
		if err := tx.Create(&bStats).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
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
