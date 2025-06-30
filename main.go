package main

import (
	"github.com/andwct/ml-feature-store-api/routers"
	"github.com/andwct/ml-feature-store-api/storage"
)

func main() {
	store := storage.NewInMemoryStore()
	r := routers.SetupRouter(store)
	r.Run(":8080")
}
