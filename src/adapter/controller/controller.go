package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/sky0621/go-transaction/domain/model"
	"github.com/sky0621/go-transaction/usecase"
	"net/http"
)

func ListTestFunc(u usecase.Test) func(c *gin.Context) {
	return func(c *gin.Context) {
		// MEMO: 本当は middleware で context.WithValue しときたい。。
		ctx := context.WithValue(c.Request.Context(), "DB", c.MustGet("DB"))

		// MEMO: 本来はGETパラメータからIDを取得して渡す。
		result, err := u.GetTest(ctx, 1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
		// MEMO: 本来はパースする。
		c.JSON(http.StatusOK, result)
	}
}

func SaveTestFunc(u usecase.Test) func(c *gin.Context) {
	return func(c *gin.Context) {
		// MEMO: 本当は middleware で context.WithValue しときたい。。
		ctx := context.WithValue(c.Request.Context(), "TX", c.MustGet("TX"))

		// MEMO: 本来はリクエストボディの中身をパースする。
		r := &model.TestModel{
			Score: 92, Result: "PASS",
		}

		// MEMO: context.Context に transaction 開始済みの *gorm.DB が隠されている。
		if err := u.SaveTest(ctx, r); err != nil {
			// MEMO: transactionMiddleware の recover に引っ掛けて rollback させる。。
			panic(err)
		}

		// MEMO: 適当
		c.JSON(http.StatusOK, gin.H{"result": "ok"})
	}
}
