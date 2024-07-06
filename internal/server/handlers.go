package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"

	openfga "github.com/ashok-an/openfga-wrapper/internal/openfga"
)

func (s *Server) rootHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = os.Getenv("APP_VERSION")

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	status := s.db.Health()
	status["openfga"] = strconv.FormatBool(openfga.IsHealthy())
	c.JSON(http.StatusOK, status)
}

func (s *Server) getStores(c *gin.Context) {
	stores := openfga.GetStores()
	c.JSON(http.StatusOK, stores)
}

func (s *Server) getStore(c *gin.Context) {
	storeID := c.Param("storeID")
	store := openfga.GetStore(storeID)
	if store.ID == "" {
		resp := make(map[string]string)
		resp["message"] = fmt.Sprintf("store with ID:%s not found", storeID)
		c.JSON(http.StatusNotFound, resp)
		return
	}
	c.JSON(http.StatusOK, store)
}

func (s *Server) getModels(c *gin.Context) {
	storeID := c.Param("storeID")
	models := openfga.GetModels(storeID)
	c.JSON(http.StatusOK, models)
}

func (s *Server) getModel(c *gin.Context) {
	storeID := c.Param("storeID")
	modelID := c.Param("modelID")
	model := openfga.GetModel(storeID, modelID)
	if model.Authorization == nil {
		resp := make(map[string]string)
		resp["message"] = fmt.Sprintf("model with ID:%s not found under storeID:%s", modelID, storeID)
		c.JSON(http.StatusNotFound, resp)
		return
	}
	c.JSON(http.StatusOK, model)
}
