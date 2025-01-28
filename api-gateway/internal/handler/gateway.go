package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/morshedulmunna/api-gateway/internal/config"
	"github.com/morshedulmunna/api-gateway/internal/proxy"
)

type Gateway struct {
	config *config.Config
}

func NewGateway(config *config.Config) *Gateway {
	return &Gateway{
		config: config,
	}
}

func (h *Gateway) ProxyAuthService(c *gin.Context) {
	targetURL := h.config.Services["auth"].BaseURL + c.Request.URL.Path
	proxy.ReverseProxy(c, targetURL)
}

func (h *Gateway) ProxyUserService(c *gin.Context) {
	targetURL := h.config.Services["users"].BaseURL + c.Request.URL.Path
	proxy.ReverseProxy(c, targetURL)
}
