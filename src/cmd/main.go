package main

import (
	"github.com/sky0621/go-transaction/adapter/gateway"
	"github.com/sky0621/go-transaction/domain"
	"github.com/sky0621/go-transaction/external/rdb"
	"github.com/sky0621/go-transaction/external/web"
	"github.com/sky0621/go-transaction/usecase"
)

// MEMO: トランザクションをhandlerレベルで扱うことの検証に関する部分以外は適当に実装。
func main() {
	db := rdb.Setup()

	// MEMO: 本来なら wire に任せる。
	gTest := gateway.NewTest()
	dTestResult := domain.NewTest(gTest)
	uTestResult := usecase.NewTest(dTestResult)

	web.Setup(db, uTestResult)
}
