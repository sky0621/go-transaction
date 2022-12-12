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

func (t test) ListTest(ctx context.Context) ([]*model.TestModel, error) {
	db := ctx.Value("DB").(*gorm.DB)
}

func (t test) SaveTestScore(ctx context.Context, score int) error {
	tx := ctx.Value("TX").(*gorm.DB)
}

func (t test) SaveTestResult(ctx context.Context, result string) error {
	tx := ctx.Value("TX").(*gorm.DB)
}
