package repository

import (
	"context"
	"github.com/sky0621/go-transaction/domain/model"
)

type Test interface {
	ListTest(ctx context.Context) ([]*model.TestModel, error)
	SaveTestScore(ctx context.Context, score int) error
	SaveTestResult(ctx context.Context, result string) error
}
