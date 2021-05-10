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
	engine.Handle(http.MethodGet, "/", handlers.Index)
	engine.Handle(http.MethodGet, "/login", handlers.Login)
	engine.Handle(http.MethodPost, "/login", handlers.Login)
	engine.Handle(http.MethodGet, "/signup", handlers.Signup)
	engine.Handle(http.MethodPost, "/signup", handlers.Signup)
	engine.Handle(http.MethodGet, "/question/read", handlers.ReadQuestion)

	route(engine, http.MethodGet, "/user/find", handlers.FindUser)

	logined := engine.Group("")
	logined.Use(middleware.Session())
	{
		logined.Handle(http.MethodGet, "/logout", handlers.Logout)
		logined.Handle(http.MethodGet, "/user/profile", handlers.Profile)
		logined.Handle(http.MethodGet, "/question/new", handlers.NewQuestion)
		logined.Handle(http.MethodPost, "/question/new", handlers.NewQuestion)
		//logined.Handle(http.MethodPost, "/answer/new", handlers.NewAnswer)
	}
}
