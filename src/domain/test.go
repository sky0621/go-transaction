package domain

import "context"

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
	//TODO implement me
	panic("implement me")
}

func (t test) SaveTestScore(ctx context.Context, score int) error {
	//TODO implement me
	panic("implement me")
}

func (t test) SaveTestResult(ctx context.Context, result string) error {
	//TODO implement me
	panic("implement me")
}
