package db

import (
	"github.com/jinzhu/gorm"
	"github.com/mui87/npb-season-stats-visualizer/domain"
)

type PlayerRepository struct {
	db *gorm.DB
}

func NewPlayerRepository(db *gorm.DB) *PlayerRepository {
	return &PlayerRepository{db: db}
}

func (pr *PlayerRepository) SearchPitchers(query string) ([]domain.Pitcher, error) {
	var pitchers []domain.Pitcher
	likeQuery := "%" + query + "%"
	if err := pr.db.Where("name LIKE ? OR kana LIKE ?", likeQuery, likeQuery).Find(&pitchers).Error; err != nil {
		return nil, err
	}

	return pitchers, nil
}

func (pr *PlayerRepository) SearchBatters(query string) ([]domain.Batter, error) {
	var batters []domain.Batter
	likeQuery := "%" + query + "%"
	if err := pr.db.Where("name LIKE ? OR kana LIKE ?", likeQuery, likeQuery).Find(&batters).Error; err != nil {
		return nil, err
	}

	return batters, nil
}
