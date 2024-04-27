package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"secret-management/internal/domain"
	"secret-management/internal/dto"
	"secret-management/internal/entities"
	"time"
)

type SecretUsecase struct {
	name       string
	repository domain.SecretRepository
	helpers    domain.SecretHelpers
}

func NewSecretUsecase(repository domain.SecretRepository, helpers domain.SecretHelpers) *SecretUsecase {
	return &SecretUsecase{
		name:       "SecretUsecase",
		repository: repository,
		helpers:    helpers,
	}
}

func generateUUID() string {
	id := uuid.New()
	return id.String()
}

func (uc *SecretUsecase) GetSecretByUserId(userId string) (*dto.GetSecretByUserIdResponse, error) {
	funcName := fmt.Sprintf("%s.GetSecretByUserId", uc.name)

	format := time.RFC3339Nano
	currentTime := time.Now().UTC()
	secret, err := uc.repository.GetSecret(userId, currentTime.Format(format))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%s.Err", funcName))
	}
	result := dto.GetSecretByUserIdResponse{}

	if secret != nil {
		result.FromSecretEntities(secret)
		return &result, nil
	}

	apiKey := uc.helpers.GenerateSecret(20)
	expiredDate := time.Now().UTC().Add(30 * 24 * time.Hour)
	id := generateUUID()
	input := entities.SecretManagement{
		Id:          id,
		UserId:      userId,
		ApiKey:      apiKey,
		ExpiredDate: expiredDate,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	err = uc.repository.CreateSecret(&input)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%s.Err", funcName))
	}

	result.FromSecretEntities(&input)

	return &result, nil
}
