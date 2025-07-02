package main

import (
	"github.com/andwct/ml-feature-store-api/routers"
	"github.com/andwct/ml-feature-store-api/storage"
)

func main() {
	store := storage.NewInMemoryStore()
	r := routers.SetupRouter(store)

	// 只掛載靜態檔案 (注意路徑與前端結構)
	r.Static("/static", "./frontend-ui/static")
	r.Static("/assets", "./frontend-ui/assets")

	// 這裡不要再註冊 r.GET("/") 了，避免重複路由錯誤

	r.Run(":8080")
}
