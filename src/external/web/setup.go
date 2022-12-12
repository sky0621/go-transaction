package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sky0621/go-transaction/adapter/controller"
	"github.com/sky0621/go-transaction/usecase"
	"gorm.io/gorm"
	"net/http"
)

func Setup(db *gorm.DB, uTest usecase.Test) {
	r := gin.Default()

	// POSTだけに適用する方法って、あるんだっけ？
	r.Use(transactionMiddleware(db))

	r.POST("/test", controller.SaveTestFunc(uTest))
	r.GET("/test", controller.ListTestFunc(uTest))

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

					ctx.JSON(http.StatusInternalServerError, r)
				}
			}()

			ctx.Set("TX", tx)

			ctx.Next()

			tx.Commit()

			return
		}

		ctx.Set("DB", db)

		ctx.Next()
	}
}
