package db

import (
	"github.com/jinzhu/gorm"
	"github.com/miiton/kanaconv"
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
	likeQuery := "%" + kanaconv.HiraganaToKatakana(query) + "%"
	if err := pr.db.Where("name LIKE ? OR kana LIKE ?", likeQuery, likeQuery).Find(&pitchers).Error; err != nil {
		return nil, err
	}

	return pitchers, nil
}

func (pr *PlayerRepository) SearchBatters(query string) ([]domain.Batter, error) {
	var batters []domain.Batter
	likeQuery := "%" + kanaconv.HiraganaToKatakana(query) + "%"
	if err := pr.db.Where("name LIKE ? OR kana LIKE ?", likeQuery, likeQuery).Find(&batters).Error; err != nil {
		return nil, err
	}

	return batters, nil
}

func (pr *PlayerRepository) GetPitcher(playerID int) (domain.Pitcher, error) {
	var pitcher domain.Pitcher
	if err := pr.db.Preload("PitcherStatsList").Find(&pitcher, playerID).Error; err != nil {
		return domain.Pitcher{}, err
	}

	return pitcher, nil
}

func (pr *PlayerRepository) GetBatter(playerID int) (domain.Batter, error) {
	var batter domain.Batter
	if err := pr.db.Preload("BatterStatsList").Find(&batter, playerID).Error; err != nil {
		return domain.Batter{}, err
	}

	return batter, nil
}
