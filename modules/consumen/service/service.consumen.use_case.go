package service

import (
	consumen "test-kreditplus/modules/consumen"
	"test-kreditplus/modules/consumen/model"
	"test-kreditplus/schemas"
)

type ConsumenService struct {
	ConsumenRepository consumen.Repository
}

func InitConsumenService(ConsumenRepository consumen.Repository) consumen.Service {
	return &ConsumenService{
		ConsumenRepository: ConsumenRepository,
	}
}

func (i *ConsumenService) CreateConsumenService(consumen schemas.CreateConsumenRequest) (error, int64) {
	consumenModel := &model.Consumen{
		Name:        &consumen.Name,
		NIK:         &consumen.NIK,
		FullName:    &consumen.FullName,
		LegalName:   &consumen.LegalName,
		BirthPlace:  &consumen.BirthPlace,
		BirthDate:   &consumen.BirthDate,
		Salary:      &consumen.Salary,
		KTPPhoto:    &consumen.KTPPhoto,
		SelfiePhoto: &consumen.SelfiePhoto,
	}
	err, id := i.ConsumenRepository.CreateConsumenRepository(consumenModel)
	if err != nil {
		return err, 0
	}

	return nil, id
}

func (i *ConsumenService) GetConsumenService() (*[]model.Consumen, error) {
	consumens, err := i.ConsumenRepository.GetConsumenRepository()
	if err != nil {
		return nil, err
	}

	return consumens, nil
}
