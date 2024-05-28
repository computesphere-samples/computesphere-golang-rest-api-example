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
