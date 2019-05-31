package router

import "github.com/mui87/npb-season-stats-visualizer/domain"

type ErrorResponse struct {
	Error string
}

func ConstructErrorResponse(errMsg string) ErrorResponse {
	return ErrorResponse{Error: errMsg}
}

type PlayerSearchResponse struct {
	Query      string
	PlayerType string
	Players    []Player
}

func ConstructPlayerSearchResponseFromPitchers(query string, pitchers []domain.Pitcher) PlayerSearchResponse {
	var players []Player
	for _, p := range pitchers {
		players = append(players, Player{ID: p.ID, Name: p.Name})
	}

	return PlayerSearchResponse{Query: query, PlayerType: "pitcher", Players: players}
}

func ConstructPlayerSearchResponseFromBatters(query string, batters []domain.Batter) PlayerSearchResponse {
	var players []Player
	for _, p := range batters {
		players = append(players, Player{ID: p.ID, Name: p.Name})
	}

	return PlayerSearchResponse{Query: query, PlayerType: "batter", Players: players}
}

type Player struct {
	ID   int
	Name string
}
