package service

import (

	"github.com/Shyyw1e/effective-mobile-test-task/internal/client"
	"github.com/Shyyw1e/effective-mobile-test-task/internal/model"
	"github.com/Shyyw1e/effective-mobile-test-task/internal/repository"
	"github.com/Shyyw1e/effective-mobile-test-task/pkg/logger"
)

type EnrichService struct {
	repo *repository.PersonRepository
}

func NewEnrichService(repo *repository.PersonRepository) *EnrichService {
	return &EnrichService{repo: repo}
}

func (s *EnrichService) EnrichAndSave(name, surname string, patronymic *string) (*model.Person, error) {
	age, err := client.GetAge(name)
	if err != nil {
		logger.Log.Error("Failed to get age", "err", err)
		return nil, err
	}
	gender, err := client.GetGender(name)
	if err != nil {
		logger.Log.Error("Failed to get gender", "err", err)
		return nil, err
	}
	nationalities, err := client.GetNationalities(name)
	if err != nil {
		logger.Log.Error("Failed to get nationalities", "err", err)
		return nil, err
	}

	person := &model.Person{
		Name: name,
		Surname: surname,
		Patronymic: patronymic,
		Age: age,
		Gender: gender,
		Nationalities: nationalities,
	}

	if err := s.repo.Save(person); err != nil {
		logger.Log.Error("Failed to save person", "err", err)
		return nil, err
	}

	logger.Log.Info("Person enriched and saved", "id", person.ID)
	return person, nil
}

