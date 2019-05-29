package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Player struct {
	ID   int    `gorm:"primary_key"`
	Name string `gorm:"not null"`
	Kana string `gorm:"not null"`
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

func (c *Client) GetPlayerIDs() ([]int, error) {
	var players []Player
	if err := c.db.Find(&players).Error; err != nil {
		return nil, err
	}

	var ids []int
	for _, player := range players {
		ids = append(ids, player.ID)
	}

	return ids, nil
}
