package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"proxy-server/internal/domain"
)

type ProxyHandler struct {
	responseRepo domain.ResponseRepository
}

func NewProxyHandler(responseRepo domain.ResponseRepository) *ProxyHandler {
	return &ProxyHandler{responseRepo: responseRepo}
}

func (h *ProxyHandler) ProxyRequest(c *gin.Context) {
	var reqBody domain.ProxyRequest
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	ctx := c.Request.Context()

	// TODO: Implement proxying logic and saving response to repository

	c.JSON(http.StatusOK, domain.ProxyResponse{})
}

func (h *ProxyHandler) GetResponse(c *gin.Context) {
	id := c.Param("id")

	ctx := c.Request.Context()

	response, err := h.responseRepo.GetResponse(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Response not found"})
		return
	}

	c.JSON(http.StatusOK, response)
}
