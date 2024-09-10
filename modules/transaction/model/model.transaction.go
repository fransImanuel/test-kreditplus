package model

import (
	"test-kreditplus/constant"
	"test-kreditplus/modules/consumen/model"
	"test-kreditplus/schemas"
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	schemas.FullAudit

	ConsumerID  *uint          `json:"consumer_id"  gorm:"column:consumer_id"`
	Consumen    model.Consumen `json:"consumer" gorm:"foreignKey:consumer_id"`
	ContractNo  *string        `json:"contract_no"  gorm:"column:contract_no"`
	OTR         *float64       `json:"otr"  gorm:"column:otr"`
	AdminFee    *float64       `json:"admin_fee"  gorm:"column:admin_fee"`
	Installment *float64       `json:"installment"  gorm:"column:installment"`
	Interest    *float64       `json:"interest"  gorm:"column:interest"`
	AssetName   *string        `json:"asset_name"  gorm:"column:asset_name"`
}

// ? this is just gorm way of custom table name
func (t *Transaction) TableName() string {
	return constant.TABLE_TRANSACTION_NAME
}

func (Transaction) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&Transaction{})
	if err != nil {
		return err
	}

	return nil
}

func (t *Transaction) InitAudit(operation string /*, user string, user_id int64*/) {
	timeNow := time.Now()
	switch operation {
	case constant.OPERATION_SQL_INSERT:
		// t.CreatedByUserName = user
		t.CreatedTime = timeNow
		// t.ModifiedByUserName = user
		// t.ModifiedTime = timeNow
	case constant.OPERATION_SQL_UPDATE:
		// t.ModifiedByUserName = user
		t.ModifiedTime = timeNow
	case constant.OPERATION_SQL_DELETE:
		// t.DeletedByUserId = &user_id
		t.DeletedTime = gorm.DeletedAt{Time: timeNow, Valid: true}
	}
}
