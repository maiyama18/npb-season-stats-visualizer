package router

import (
	"fmt"
	"net/http"

	"github.com/mui87/npb-season-stats-visualizer/api/infrastructure/db"
	"github.com/mui87/npb-season-stats-visualizer/api/usecase"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
)

type PlayerController struct {
	playerService *usecase.PlayerService
}

func NewPlayerController(gdb *gorm.DB) *PlayerController {
	playerRepository := db.NewPlayerRepository(gdb)
	playerService := usecase.NewPlayerService(playerRepository)
	return &PlayerController{playerService: playerService}
}

func (pc *PlayerController) SearchPitchers(c *gin.Context) {
	queries, ok := c.Request.URL.Query()["query"]
	if !ok {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "query parameter should be specified"})
		return
	}
	query := queries[0]

	pitchers, err := pc.playerService.SearchPitchers(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConstructErrorResponse(fmt.Sprintf("search failed: %s", err)))
		return
	}

	c.JSON(http.StatusOK, ConstructPlayerSearchResponseFromPitchers(query, pitchers))
}

func (pc *PlayerController) SearchBatters(c *gin.Context) {
	queries, ok := c.Request.URL.Query()["query"]
	if !ok {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "query parameter should be specified"})
		return
	}
	query := queries[0]

	batters, err := pc.playerService.SearchBatters(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, ConstructErrorResponse(fmt.Sprintf("search failed: %s", err)))
		return
	}

	c.JSON(http.StatusOK, ConstructPlayerSearchResponseFromBatters(query, batters))
}
