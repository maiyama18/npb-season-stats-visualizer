package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/mui87/npb-season-stats-visualizer/api/infrastructure/db"
	"github.com/mui87/npb-season-stats-visualizer/api/usecase"
)

type StatsController struct {
	playerService *usecase.PlayerService
}

func NewStatsController(gdb *gorm.DB) *StatsController {
	playerRepository := db.NewPlayerRepository(gdb)
	playerService := usecase.NewPlayerService(playerRepository)
	return &StatsController{playerService: playerService}
}

func (sc *StatsController) GetPitcherStats(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, constructErrorResponse("id should be an integer. got: %s", idStr))
		return
	}

	pitcher, err := sc.playerService.GetPitcher(id)
	if err != nil {
		c.JSON(http.StatusNotFound, constructErrorResponse("could not find pitcher with id %s", id))
		return
	}

	c.JSON(http.StatusOK, constructPitcherStatsResponse(pitcher))
}
