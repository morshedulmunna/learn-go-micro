package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/morshedulmunna/api-gateway/internal/config"
	"github.com/morshedulmunna/api-gateway/internal/proxy"
)

type GatewayHandler struct {
	config *config.Config
}

func NewGatewayHandler(config *config.Config) *GatewayHandler {
	return &GatewayHandler{
		config: config,
	}
}

func (h *GatewayHandler) ProxyAuthService(c *gin.Context) {
	targetURL := h.config.Services["auth"].BaseURL + c.Request.URL.Path
	proxy.ReverseProxy(c, targetURL)
}

func (h *GatewayHandler) ProxyUserService(c *gin.Context) {
	targetURL := h.config.Services["users"].BaseURL + c.Request.URL.Path
	proxy.ReverseProxy(c, targetURL)
}
