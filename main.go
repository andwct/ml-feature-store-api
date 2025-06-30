package main

import (
	"ml-feature-store-api/routers"
	"ml-feature-store-api/storage"
)

func main() {
	store := storage.NewInMemoryStore()
	r := routers.SetupRouter(store)
	r.Run(":8080")
}
