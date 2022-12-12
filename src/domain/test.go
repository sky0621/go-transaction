package domain

import (
	"context"
	"errors"
)

type TestModel struct {
	Score  int
	Result string
}

type Test interface {
	ListTest(ctx context.Context) ([]*TestModel, error)
	SaveTestScore(ctx context.Context, score int) error
	SaveTestResult(ctx context.Context, result string) error
}

func NewTest() Test {
	return &test{}
}

type test struct {
}

func (t test) ListTest(ctx context.Context) ([]*TestModel, error) {
	return nil, errors.New("implement me")
}

func (t test) SaveTestScore(ctx context.Context, score int) error {
	return errors.New("implement me")
}

func (t test) SaveTestResult(ctx context.Context, result string) error {
	return errors.New("implement me")
}
