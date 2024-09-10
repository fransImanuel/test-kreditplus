package model

import (
	"test-kreditplus/constant"
	"test-kreditplus/schemas"
	"time"

	"gorm.io/gorm"
)

type Consumen struct {
	schemas.FullAudit

	Name        *string  `json:"name"  gorm:"column:name"`
	NIK         *string  `json:"nik"  gorm:"column:nik"`
	FullName    *string  `json:"full_name"  gorm:"column:full_name"`
	LegalName   *string  `json:"legal_name"  gorm:"column:legal_name"`
	BirthPlace  *string  `json:"birth_place"  gorm:"column:birth_place"`
	BirthDate   *string  `json:"birth_date"  gorm:"column:birth_date"`
	Salary      *float64 `json:"salary"  gorm:"column:salary"`
	KTPPhoto    *string  `json:"ktp_photo"  gorm:"column:ktp_photo"`
	SelfiePhoto *string  `json:"selfie_photo"  gorm:"column:selfie_photo"`
	// Limits      *[]Limit `gorm:"foreignKey:ConsumerID" json:"limits"`
}

type Limit struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	ConsumerID uint     `json:"consumer_id"  gorm:"consumer_id"`
	Consumen   Consumen `json:"consumer" gorm:"foreignKey:consumer_id"`
	Tenor      int      `json:"tenor"  gorm:"tenor"`
	Limit      float64  `json:"limit"  gorm:"limit"`
}

// ? this is just gorm way of custom table name
func (t *Consumen) TableName() string {
	return constant.TABLE_CONSUMEN_NAME
}

func (Consumen) Migrate(tx *gorm.DB) error {
	err := tx.AutoMigrate(&Consumen{})
	if err != nil {
		return err
	}
	err = tx.AutoMigrate(&Limit{})
	if err != nil {
		return err
	}

	return nil
}

func (t *Consumen) InitAudit(operation string /*, user string, user_id int64*/) {
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
