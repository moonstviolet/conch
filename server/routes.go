package server

import (
	"conch/handlers"
	"conch/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func route(routes gin.IRoutes, method string, path string, function interface{}) {
	routes.Handle(method, path, middleware.CreateHandlerFunc(function))
}

func routes(engine *gin.Engine) {
	route(engine, http.MethodGet, "/example", handlers.ShowExample)
}
