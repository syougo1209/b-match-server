package usecase

import "context"

type Transaction interface {
	BeginTx(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
}
