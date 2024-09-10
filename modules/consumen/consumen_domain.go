package users

import (
	"test-kreditplus/modules/consumen/model"
	"test-kreditplus/schemas"
)

type Repository interface {
	CreateConsumenRepository(consumen *model.Consumen) (error, int64)
	GetConsumenRepository() (*[]model.Consumen, error)
}

type Service interface {
	CreateConsumenService(consumen schemas.CreateConsumenRequest) (error, int64)
	GetConsumenService() (*[]model.Consumen, error)
}
