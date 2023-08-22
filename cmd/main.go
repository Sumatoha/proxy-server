package cmd

import (
	"awesomeProject1/internal/api"
	"github.com/gin-gonic/gin"
)

package main

import (
"context"
"github.com/gin-gonic/gin"
"proxy-server/internal/api"
"proxy-server/internal/domain"
"proxy-server/internal/infrastructure"
)

func main() {
	r := gin.Default()

	repo := infrastructure.NewInMemoryRepository()
	handler := api.NewProxyHandler(repo)

	r.POST("/proxy", handler.ProxyRequest)
	r.GET("/response/:id", handler.GetResponse)

	r.Run(":8080")
}