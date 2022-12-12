package usecase

import (
	"context"
	"github.com/sky0621/go-transaction/domain"
	"github.com/sky0621/go-transaction/domain/model"
)

type Test interface {
	GetTest(ctx context.Context, id int) (*model.TestModel, error)
	SaveTest(ctx context.Context, r *model.TestModel) error
}

func NewTest(d domain.Test) Test {
	return &test{d: d}
}

type test struct {
	d domain.Test
}

func (t test) GetTest(ctx context.Context, id int) (*model.TestModel, error) {
	return t.d.GetTest(ctx, id)
}

func (t test) SaveTest(ctx context.Context, r *model.TestModel) error {
	if err := t.d.SaveTestScore(ctx, r.Score); err != nil {
		return err
	}
	if err := t.d.SaveTestResult(ctx, r.Result); err != nil {
		return err
	}
	return nil
}
