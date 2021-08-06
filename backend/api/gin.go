package api

import (
	"github.com/techotron/online-quote-book/backend/controllers"
	"github.com/techotron/online-quote-book/backend/log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginlogrus "github.com/toorop/gin-logrus"
)

// Setup API using gin with routes and cors settings
func Setup() *gin.Engine {
	router := gin.New()
	router.Use(ginlogrus.Logger(log.Logger), gin.Recovery())
	addAllowAllCors(router)
	addRoutes(router)
	return router
}

// addAllowAllCors sets cors config
func addAllowAllCors(g *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	g.Use(cors.New(corsConfig))
}

// addRoutes using existing controllers
func addRoutes(g *gin.Engine) {
	g.GET("/info", controllers.GetInfo)
	
	g.GET("/quotes/:quoteBook", controllers.GetQuotes)

	g.GET("/quotebooks", controllers.GetAllQuoteBooks)
	g.POST("/quotebook/:quoteBook", controllers.AddQuoteBook)
}
