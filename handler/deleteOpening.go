package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCreateOpeningHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "GO primary routes",
	})
}
