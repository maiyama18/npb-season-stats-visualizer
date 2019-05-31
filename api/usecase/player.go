package usecase

import "github.com/mui87/npb-season-stats-visualizer/domain"

type PlayerRepository interface {
	SearchPitchers(query string) ([]domain.Pitcher, error)
	SearchBatters(query string) ([]domain.Batter, error)
	GetPitcher(playerID int) (domain.Pitcher, error)
	GetBatter(playerID int) (domain.Batter, error)
}

type PlayerService struct {
	playerRepository PlayerRepository
}

func NewPlayerService(playerRepository PlayerRepository) *PlayerService {
	return &PlayerService{playerRepository: playerRepository}
}

func (ps *PlayerService) SearchPitchers(query string) ([]domain.Pitcher, error) {
	if query == "" {
		return []domain.Pitcher{}, nil
	}

	return ps.playerRepository.SearchPitchers(query)
}

func (ps *PlayerService) SearchBatters(query string) ([]domain.Batter, error) {
	if query == "" {
		return []domain.Batter{}, nil
	}

	return ps.playerRepository.SearchBatters(query)
}

func (ps *PlayerService) GetPitcher(playerID int) (domain.Pitcher, error) {
	return ps.playerRepository.GetPitcher(playerID)
}

func (ps *PlayerService) GetBatter(playerID int) (domain.Batter, error) {
	return ps.playerRepository.GetBatter(playerID)
}
