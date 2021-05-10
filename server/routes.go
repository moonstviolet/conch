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
	//route(engine, http.MethodGet, "/example", handlers.ShowExample)
	//files := http.FileServer(http.Dir(config.Static))

	engine.Handle(http.MethodGet, "/", handlers.Index)
	engine.Handle(http.MethodGet, "/login", handlers.Login)
	engine.Handle(http.MethodPost, "/login", handlers.Login)
	engine.Handle(http.MethodGet, "/signup", handlers.Signup)
	engine.Handle(http.MethodPost, "/signup", handlers.Signup)

	route(engine, http.MethodGet, "/user/find", handlers.FindUser)

	logined := engine.Group("")
	logined.Use(middleware.Session())
	{
		logined.Handle(http.MethodGet, "/logout", handlers.Logout)
		logined.Handle(http.MethodGet, "/user/profile", handlers.Profile)
	}

	// route(engine, http.MethodPost, "/question/new", handlers.NewQuestion)
	// route(engine, http.MethodGet, "/question/read", handlers.ReadQuestion)

	// route(engine, http.MethodPost, "/answer/new", handlers.NewAnswer)
}
