package main

import (
	"github.com/sky0621/go-transaction/external/rdb"
	"github.com/sky0621/go-transaction/external/web"
)

// MEMO: トランザクションをhandlerレベルで扱うことの検証に関する部分以外は適当に実装。
func main() {
	db := rdb.Setup()
	web.Setup(db)
}

// Test
//   Score
//   Result
