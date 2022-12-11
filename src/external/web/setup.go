package web

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Setup(db *gorm.DB) {
	r := gin.Default()
	r.Use(transactionMiddleware(db))

	r.POST("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"": ""})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello!"})
	})

	r.Run()
}

func transactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tx := db.Begin()
		ctx.Set("TX", tx)
		ctx.Next()
	}
}
