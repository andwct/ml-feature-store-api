package routers

import (
	"ml-feature-store-api/handlers"
	"ml-feature-store-api/storage"

	"github.com/gin-gonic/gin"
)

func SetupRouter(store *storage.InMemoryStore) *gin.Engine {
	r := gin.Default()

	featureHandler := handlers.NewFeatureHandler(store)

	r.POST("/features", featureHandler.SetFeature)
	r.GET("/features/:entity_id/:name", featureHandler.GetFeature)

	return r
}
