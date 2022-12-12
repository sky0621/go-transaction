package main

import (
	"github.com/sky0621/go-transaction/domain"
	"github.com/sky0621/go-transaction/external/rdb"
	"github.com/sky0621/go-transaction/external/web"
	"github.com/sky0621/go-transaction/usecase"
)

// MEMO: トランザクションをhandlerレベルで扱うことの検証に関する部分以外は適当に実装。
func main() {
	db := rdb.Setup()

	// MEMO: 本来なら wire に任せる。
	dTestResult := domain.NewTest()
	uTestResult := usecase.NewTest(dTestResult)

	web.Setup(db, uTestResult)
}
