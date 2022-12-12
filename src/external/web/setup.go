package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sky0621/go-transaction/adapter/rest"
	"github.com/sky0621/go-transaction/usecase"
	"gorm.io/gorm"
	"net/http"
)

func Setup(db *gorm.DB, uTest usecase.Test) {
	r := gin.Default()

	// POSTだけに適用する方法って、あるんだっけ？
	r.Use(transactionMiddleware(db))

	r.POST("/test", rest.SaveTest(uTest))
	r.GET("/test", rest.ListTest(uTest))

	r.Run()
}

func transactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method

		if method == http.MethodPost || method == http.MethodPatch || method == http.MethodPut || method == http.MethodDelete {
			tx := db.Begin()
			defer func() {
				if r := recover(); r != nil {
					tx.Rollback()
				}
			}()

			// MEMO: gin.Context への set は結局 context.Context への再変換が必要になるので使わない。
			context.WithValue(context.Background(), "TX", tx)

			ctx.Next()

			tx.Commit()

			return
		}

		// MEMO: gin.Context への set は結局 context.Context への再変換が必要になるので使わない。
		context.WithValue(context.Background(), "DB", db)

		ctx.Next()
	}
}
