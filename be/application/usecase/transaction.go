package usecase

import "context"

//go:generate mockgen -source=transaction.go -destination=../../mock/usecase/transaction.go
type Transaction interface {
	BeginTx(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
}
