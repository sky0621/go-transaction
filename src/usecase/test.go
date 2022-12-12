package usecase

import (
	"context"
	"github.com/sky0621/go-transaction/domain"
)

type Test interface {
	ListTest(ctx context.Context) ([]*domain.TestModel, error)
	SaveTest(ctx context.Context, r *domain.TestModel) error
}

func NewTest(d domain.Test) Test {
	return &test{d: d}
}

type test struct {
	d domain.Test
}

func (t test) ListTest(ctx context.Context) ([]*domain.TestModel, error) {
	return t.d.ListTest(ctx)
}

func (t test) SaveTest(ctx context.Context, r *domain.TestModel) error {
	if err := t.d.SaveTestScore(ctx, r.Score); err != nil {
		return err
	}
	if err := t.d.SaveTestResult(ctx, r.Result); err != nil {
		return err
	}
	return nil
}
