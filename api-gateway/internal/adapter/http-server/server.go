package service

import (
	"github.com/gin-gonic/gin"
	"github.com/trippingcoin/MooDu/api-gateway/config"
)

type API struct {
	server *gin.Engine
	cfg    config.HTTPServer
	addr   string

	// clientHandler          *handler.Client
	// clientStatisticHandler *handler.ClientStatistic
}

// func New(cfg config.Server)
