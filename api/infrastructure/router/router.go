package router

import (
	"os"

	"github.com/mui87/npb-season-stats-visualizer/api/infrastructure/db"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gin-gonic/gin"
)

func Run() error {
	r := gin.Default()

	gdb, err := db.NewGormDB()
	if err != nil {
		return err
	}

	playerController := NewPlayerController(gdb)
	statsController := NewStatsController(gdb)

	api := r.Group("/api")
	{
		api.GET("/search/pitchers", func(c *gin.Context) { playerController.SearchPitchers(c) })
		api.GET("/search/batters", func(c *gin.Context) { playerController.SearchBatters(c) })
		api.GET("/stats/pitcher/:id", func(c *gin.Context) { statsController.GetPitcherStats(c) })
	}

	port := getEnv("SERVER_PORT", "8080")

	if err := r.Run(":" + port); err != nil {
		return err
	}

	return nil
}

func getEnv(envName, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		return defaultValue
	}
	return value
}
