package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// MEMO: トランザクションをhandlerレベルで扱うことの検証に関する部分以外は適当に実装。
func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(transactionMiddleware())

	r.Run()
}

func transactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
	}
}
