package repository

import (
	"context"
	"github.com/sky0621/go-transaction/domain/model"
)

type Test interface {
	GetTest(ctx context.Context, id int) (*model.TestModel, error)
	SaveTestScore(ctx context.Context, score int) error
	SaveTestResult(ctx context.Context, result string) error
}
