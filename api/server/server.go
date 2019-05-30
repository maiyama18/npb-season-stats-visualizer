package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/mui87/npb-season-stats-visualizer/db"

	"github.com/gin-gonic/gin"
)

type Server struct {
	dbClient *db.Client
	port     string
}

func New() (*Server, error) {
	port := getEnv("SERVER_PORT", "8080")

	var emptyEnvVars []string
	dbUser := getEnv("DB_USER", "")
	if dbUser == "" {
		emptyEnvVars = append(emptyEnvVars, "DB_USER")
	}
	dbPassword := getEnv("DB_PASSWORD", "")
	if dbPassword == "" {
		emptyEnvVars = append(emptyEnvVars, "DB_PASSWORD")
	}
	dbHost := getEnv("DB_HOST", "")
	if dbHost == "" {
		emptyEnvVars = append(emptyEnvVars, "DB_HOST")
	}
	dbPort := getEnv("DB_PORT", "")
	if dbPort == "" {
		emptyEnvVars = append(emptyEnvVars, "DB_PORT")
	}
	dbSchema := getEnv("DB_SCHEMA", "")
	if dbSchema == "" {
		emptyEnvVars = append(emptyEnvVars, "DB_SCHEMA")
	}
	if len(emptyEnvVars) > 0 {
		return nil, fmt.Errorf("the following environment variables should be set: %s", strings.Join(emptyEnvVars, ", "))
	}

	dbClient, err := db.NewClient(dbUser, dbPassword, dbHost, dbPort, dbSchema, false)
	if err != nil {
		return nil, err
	}

	return &Server{
		dbClient: dbClient,
		port:     port,
	}, nil
}

func (s *Server) Run() error {
	defer s.dbClient.CloseDB()

	r := gin.Default()

	// 検索用API: query(名前の一部と野手/投手)を受け取って、該当する選手一覧を返す
	// 成績を返す用API: 選手のidを受け取って成績一覧を返す
	api := r.Group("/api")
	{
		api.GET("/search/pitchers", s.searchPitchers)
		api.GET("/search/batters", s.searchBatters)
	}

	if err := r.Run(":" + s.port); err != nil {
		return err
	}

	return nil
}

func (s *Server) searchPitchers(c *gin.Context) {
	query, ok := c.Request.URL.Query()["query"]
	if !ok || query[0] == "" {
		c.JSON(http.StatusOK, []db.Player{})
		return
	}

	players, err := s.dbClient.SearchPitchers(query[0])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to search pitchers: %s", err)})
		return
	}

	c.JSON(http.StatusOK, players)
}

func (s *Server) searchBatters(c *gin.Context) {
	query, ok := c.Request.URL.Query()["query"]
	if !ok || query[0] == "" {
		c.JSON(http.StatusOK, []db.Player{})
		return
	}

	players, err := s.dbClient.SearchBatters(query[0])
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to search batters: %s", err)})
		return
	}

	c.JSON(http.StatusOK, players)
}

func getEnv(envName, defaultValue string) string {
	value := os.Getenv(envName)
	if value == "" {
		return defaultValue
	}
	return value
}
