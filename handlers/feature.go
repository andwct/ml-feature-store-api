package handlers

import (
	"log"
	"net/http"

	"github.com/andwct/ml-feature-store-api/models"
	"github.com/andwct/ml-feature-store-api/storage"
	"github.com/gin-gonic/gin"
)

type FeatureHandler struct {
	Store *storage.InMemoryStore
}

func NewFeatureHandler(store *storage.InMemoryStore) *FeatureHandler {
	return &FeatureHandler{Store: store}
}

func (h *FeatureHandler) SetFeature(c *gin.Context) {
	var feature models.Feature
	if err := c.ShouldBindJSON(&feature); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	key := feature.HistDate + ":" + feature.Name
	h.Store.Set(key, feature)
	c.JSON(http.StatusOK, gin.H{"message": "feature saved"})
}

func (h *FeatureHandler) GetFeature(c *gin.Context) {
	entityID := c.Query("entity_id")
	name := c.Query("name")
	log.Printf("GetFeature called with entity_id=%s, name=%s\n", entityID, name)

	if entityID == "" || name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "entity_id and name query params are required"})
		return
	}
	key := entityID + ":" + name
	if val, ok := h.Store.Get(key); ok {
		c.JSON(http.StatusOK, val)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "feature not found"})
	}
}
