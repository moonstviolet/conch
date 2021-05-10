package server

import (
	"conch/middleware"
	"conch/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func route(routes gin.IRoutes, method string, path string, function interface{}) {
	routes.Handle(method, path, middleware.CreateHandlerFunc(function))
}

func routes(engine *gin.Engine) {
	//route(engine, http.MethodGet, "/example", handlers.ShowExample)
	//files := http.FileServer(http.Dir(config.Static))

	engine.Handle(http.MethodGet, "/", func(c *gin.Context) {
		qlist, _ := models.QuestionList()
		c.HTML(200, "index.html", gin.H{
			"QList": qlist,
		})
	})

	//route(engine, http.MethodGet, "/", handlers.Index)
	// route(engine, http.MethodGet, "/login", handlers.Login)
	// route(engine, http.MethodPost, "/logout", handlers.Logout)
	// route(engine, http.MethodPost, "/signup", handlers.Signup)
	// route(engine, http.MethodGet, "/user/find", handlers.FindUser)
	// route(engine, http.MethodGet, "/user/profile", handlers.Profile)

	// route(engine, http.MethodPost, "/question/new", handlers.NewQuestion)
	// route(engine, http.MethodGet, "/question/read", handlers.ReadQuestion)

	// route(engine, http.MethodPost, "/answer/new", handlers.NewAnswer)
}
