package service_test

import (
	"log/slog"
	"os"
	"testing"

	"github.com/Shyyw1e/effective-mobile-test-task/internal/model"
	"github.com/Shyyw1e/effective-mobile-test-task/internal/repository"
	"github.com/Shyyw1e/effective-mobile-test-task/internal/service"
	"github.com/Shyyw1e/effective-mobile-test-task/pkg/logger"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type FakeClient struct{}

func (f *FakeClient) GetAge(name string) (int, error) {
	return 30, nil
}

func (f *FakeClient) GetGender(name string) (string, error) {
	return "male", nil
}

func (f *FakeClient) GetNationalities(name string) ([]string, error) {
	return []string{"RU"}, nil
}

func TestMain(m *testing.M) {
	logger.Log = logger.New(slog.LevelDebug) 
	os.Exit(m.Run())
}


func TestEnrichAndSave(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	err = db.AutoMigrate(&model.Person{})
	assert.NoError(t, err)

	repo := repository.NewPersonRepository(db)
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	svc := service.NewEnrichService(repo, &FakeClient{}, logger)

	p, err := svc.EnrichAndSave("Ivan", "Petrov", nil)
	assert.NoError(t, err)
	

	assert.Equal(t, "Ivan", p.Name)
	assert.Equal(t, "Petrov", p.Surname)
	assert.Nil(t, (*string)(nil), p.Patronymic)
	assert.Equal(t, 30, p.Age)
	assert.Equal(t, "male", p.Gender)
	assert.ElementsMatch(t, []string{"RU"}, p.Nationalities)

	var found model.Person
	err = db.First(&found, "name = ?", "Ivan").Error
	assert.NoError(t, err)
	assert.Equal(t, p.ID, found.ID)
}