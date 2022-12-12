package domain

import (
	"context"
	"github.com/sky0621/go-transaction/domain/model"
	"github.com/sky0621/go-transaction/domain/repository"
)

type Test interface {
	GetTest(ctx context.Context, id int) (*model.TestModel, error)
	SaveTestScore(ctx context.Context, score int) error
	SaveTestResult(ctx context.Context, result string) error
}

func NewTest(repo repository.Test) Test {
	return &test{repo: repo}
}

type test struct {
	repo repository.Test
}

func (t test) GetTest(ctx context.Context, id int) (*model.TestModel, error) {
	return t.repo.GetTest(ctx, id)
}

func (t test) SaveTestScore(ctx context.Context, score int) error {
	return t.repo.SaveTestScore(ctx, score)
}

func (t test) SaveTestResult(ctx context.Context, result string) error {
	return t.repo.SaveTestResult(ctx, result)
}
