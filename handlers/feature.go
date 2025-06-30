package handlers

import (
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
	entityID := c.Param("entity_id")
	name := c.Param("name")
	key := entityID + ":" + name
	if val, ok := h.Store.Get(key); ok {
		c.JSON(http.StatusOK, val)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "feature not found"})
	}
}
