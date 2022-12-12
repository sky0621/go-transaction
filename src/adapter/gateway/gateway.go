package gateway

import (
	"context"
	"github.com/sky0621/go-transaction/domain/model"
	"github.com/sky0621/go-transaction/domain/repository"
	"gorm.io/gorm"
)

func NewTest() repository.Test {
	return &test{}
}

type test struct {
}

func (t test) GetTest(ctx context.Context, id int) (*model.TestModel, error) {
	db := ctx.Value("DB").(*gorm.DB)

	// MEMO: 超適当
	var resScore *TestScore
	db.Model(&TestScore{ID: id}).First(&resScore)
	var resResult *TestResult
	db.Model(&TestScore{ID: id}).First(&resResult)
	return &model.TestModel{
		Score:  resScore.Score,
		Result: resResult.Result,
	}, nil
}

func (t test) SaveTestScore(ctx context.Context, score int) error {
	tx := ctx.Value("TX").(*gorm.DB)

	res := tx.Create(&TestScore{
		ID:    1,
		Score: score,
	})
	return res.Error
}

func (t test) SaveTestResult(ctx context.Context, result string) error {
	tx := ctx.Value("TX").(*gorm.DB)
	res := tx.Create(&TestResult{
		ID:     1,
		Result: result,
	})
	return res.Error
}

type TestScore struct {
	ID    int
	Score int
}

type TestResult struct {
	ID     int
	Result string
}
