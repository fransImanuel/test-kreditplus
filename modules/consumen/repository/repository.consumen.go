package repository

import (
	"test-kreditplus/constant"
	consumen "test-kreditplus/modules/consumen"
	"test-kreditplus/modules/consumen/model"

	"gorm.io/gorm"
)

type ConsumenRepository struct {
	DBPostgres *gorm.DB
	//DBMongoDB
	//DBMinio, etc
}

func InitConsumenRepository(db *gorm.DB) consumen.Repository {
	return &ConsumenRepository{
		DBPostgres: db,
	}
}

func (u *ConsumenRepository) CreateConsumenRepository(consumen *model.Consumen) (error, int64) {
	db := u.DBPostgres

	consumen.InitAudit(constant.OPERATION_SQL_INSERT)

	results := db.Create(&consumen)
	if results.Error != nil {
		return results.Error, 0
	}

	return nil, consumen.ID

}

func (u *ConsumenRepository) GetConsumenRepository() (*[]model.Consumen, error) {
	var items *[]model.Consumen
	db := u.DBPostgres

	// Get all records
	results := db.Find(&items)
	// SELECT * FROM ITEMS;
	if results.Error != nil {
		return nil, results.Error
	}

	return items, nil
}
