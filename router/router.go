package router

import (
	"github.com/EmeraldLS/quote-generator/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(port string) error {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/api/quotes/populate", controller.PopulateDBFromFrontnend)
	r.POST("/api/quotes/insert", controller.CreateQuoteFromFrontend)
	r.GET("api/quotes", controller.GetAllQuotesFromFrontend)
	r.GET("/api/quotes/related", controller.GetRelatedQuoteFromFrontend)
	r.DELETE("/api/quotes/one", controller.DeleteOneFromFrontend)
	r.DELETE("/api/quotes", controller.DeleteAllFromFrontend)
	r.DELETE("/api/quotes/related", controller.DeleteSimilarFromFrontend)
	return r.Run(port)
}
