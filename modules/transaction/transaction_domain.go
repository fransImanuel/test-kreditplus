package transaction

import (
	"test-kreditplus/modules/transaction/model"
	"test-kreditplus/schemas"
)

type Repository interface {
	CreateTransactionRepository(transaction *model.Transaction) (error, int64)
}

type Service interface {
	CreateTransactionService(transaction schemas.CreateTransactionRequest) (error, int64)
}
