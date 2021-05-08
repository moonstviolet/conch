package server

import (
	"conch/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer() *http.Server {
	if err := config.Load(); err != nil {
		log.Fatalf("load config with error: %v:", err)
		return nil
	}
	gConfig := config.GetConfig()
	gin.SetMode(gConfig.Server.RunMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	routes(router)
	server := &http.Server{
		Addr:           ":" + gConfig.Server.HttpPort,
		Handler:        router,
		ReadTimeout:    gConfig.Server.ReadTimeout,
		WriteTimeout:   gConfig.Server.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	return server
}