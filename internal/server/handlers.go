package server

import (
	"fmt"
	"io"
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

func (s *Server) CreateModel(c *gin.Context) {
	storeID := c.Param("storeID")
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}

	output := openfga.CreateModel(storeID, string(body))
	fmt.Println(output)
	if output.ModelID == "" {
		resp := make(map[string]string)
		resp["message"] = fmt.Sprintf("model creation failed for storeID:%s", storeID)
		c.JSON(http.StatusUnprocessableEntity, resp)
		return
	}
	c.JSON(http.StatusOK, output)
}
