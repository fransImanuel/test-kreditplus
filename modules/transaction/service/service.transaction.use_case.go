package service

import (
	"test-kreditplus/modules/transaction"
	"test-kreditplus/modules/transaction/model"
	"test-kreditplus/schemas"
)

type TransactionService struct {
	TransactionRepository transaction.Repository
}

func InitTransactionService(TransactionRepository transaction.Repository) transaction.Service {
	return &TransactionService{
		TransactionRepository: TransactionRepository,
	}
}

func (i *TransactionService) CreateTransactionService(transaction schemas.CreateTransactionRequest) (error, int64) {
	transactionModel := &model.Transaction{
		ConsumerID:  &transaction.ConsumerID,
		ContractNo:  &transaction.ContractNo,
		OTR:         &transaction.OTR,
		AdminFee:    &transaction.AdminFee,
		Installment: &transaction.Installment,
		Interest:    &transaction.Interest,
		AssetName:   &transaction.AssetName,
	}
	err, id := i.TransactionRepository.CreateTransactionRepository(transactionModel)
	if err != nil {
		return err, 0
	}

	return nil, id
}
