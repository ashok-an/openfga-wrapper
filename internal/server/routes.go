package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.rootHandler)
	r.GET("/health", s.healthHandler)
	r.GET("/healthz", s.healthHandler)
	r.GET("/stores", s.getStores)
	r.GET("/stores/:storeID", s.getStore)
	r.GET("/stores/:storeID/models", s.getModels)
	r.GET("/stores/:storeID/models/:modelID", s.getModel)

	return r
}
