package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{		
		v1.GET("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "GO primary routes",
			})
		})
		v1.POST("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "GO primary routes",
			})
		})
		v1.DELETE("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "GO primary routes",
			})
		})
		v1.PUT("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "GO primary routes",
			})
		})
		v1.GET("/openings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "All",
			})
		})
	}
}
