package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/sky0621/go-transaction/domain"
	"github.com/sky0621/go-transaction/usecase"
	"net/http"
)

func ListTest(u usecase.Test) func(c *gin.Context) {
	return func(c *gin.Context) {
		results, err := u.ListTest(c.Request.Context())
		if err != nil {
			// MEMO: 適当
			panic(err)
		}
		// MEMO: 本来はパースする。
		c.JSON(http.StatusOK, results)
	}
}

func SaveTest(u usecase.Test) func(c *gin.Context) {
	return func(c *gin.Context) {
		// MEMO: 本来はリクエストボディの中身をパースする。
		r := &domain.TestModel{
			Score: 92, Result: "PASS",
		}

		// MEMO: context.Context に transaction 開始済みの *gorm.DB が隠されている。
		if err := u.SaveTest(c.Request.Context(), r); err != nil {
			// MEMO: 適当
			panic(err)
		}

		// MEMO: 適当
		c.JSON(http.StatusOK, gin.H{"result": "ok"})
	}
}
