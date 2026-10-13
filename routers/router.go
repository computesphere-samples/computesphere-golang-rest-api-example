package routers

import (
	"github.com/gin-gonic/gin"

	v1 "computesphere.com/computesphere-golang-rest-api-example/routers/api/v1"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// The root answers too, so opening the app's URL shows it is up and
	// lists what it serves instead of a 404.
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ComputeSphere Go REST API example",
			"endpoints": []string{
				"GET /ping",
				"GET /api/v1/articles",
				"GET /api/v1/articles/:id",
				"POST /api/v1/articles",
				"PUT /api/v1/articles/:id",
				"DELETE /api/v1/articles/:id",
			},
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1Group := r.Group("/api/v1")
	{
		v1Group.GET("/articles", v1.GetArticles)
		v1Group.GET("/articles/:id", v1.GetArticle)
		v1Group.POST("/articles", v1.AddArticle)
		v1Group.PUT("/articles/:id", v1.EditArticle)
		v1Group.DELETE("/articles/:id", v1.DeleteArticle)
	}
	return r
}
