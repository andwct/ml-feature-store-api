package routers

import (
	"github.com/andwct/ml-feature-store-api/storage"

	"github.com/andwct/ml-feature-store-api/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(store *storage.InMemoryStore) *gin.Engine {
	r := gin.Default()

	featureHandler := handlers.NewFeatureHandler(store)

	r.POST("/features", featureHandler.SetFeature)
	r.GET("/features", featureHandler.GetFeature) // change from /features/:entity_id/:name
	r.GET("/", func(c *gin.Context) {
		c.File("./frontend-ui/index.html")
	})

	return r
}
