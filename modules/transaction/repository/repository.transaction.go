package repository

import (
	"test-kreditplus/constant"
	consumenModel "test-kreditplus/modules/consumen/model"
	transaction "test-kreditplus/modules/transaction"
	"test-kreditplus/modules/transaction/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TransactionRepository struct {
	DBPostgres *gorm.DB
	//DBMongoDB
	//DBMinio, etc
}

func InitTransactionRepository(db *gorm.DB) transaction.Repository {
	return &TransactionRepository{
		DBPostgres: db,
	}
}

func (u *TransactionRepository) CreateTransactionRepository(transaction *model.Transaction) (error, int64) {
	db := u.DBPostgres

	transaction.InitAudit(constant.OPERATION_SQL_INSERT)

	//! Handling concurrent
	err := db.Transaction(func(tx *gorm.DB) error {
		var consumer consumenModel.Consumen
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&consumer, transaction.ConsumerID).Error; err != nil {
			return err
		}

		// Validate and update consumer limit here

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err, 0
	}

	return nil, transaction.ID

}
