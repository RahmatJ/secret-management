package usecase

import (
	"fmt"
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

func (uc *SecretUsecase) GetSecretByUserId(userId string) (*dto.GetSecretByUserIdResponse, error) {
	funcName := fmt.Sprintf("%s.GetSecretByUserId", uc.name)

	currentTime := time.Now()
	secret, err := uc.repository.GetSecret(userId, currentTime)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%s.Err", funcName))
	}

	if secret != nil {
		result := dto.GetSecretByUserIdResponse{}
		result.FromSecretEntities(secret)
		return &result, nil
	}

	apiKey := uc.helpers.GenerateSecret(20)
	expiredDate := time.Now().Add(30 * 24 * time.Hour)

	_ = entities.SecretManagement{
		// TODO(Rahmat): add function to generate UUID
		Id:          "",
		UserId:      userId,
		ApiKey:      apiKey,
		ExpiredDate: expiredDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result := dto.GetSecretByUserIdResponse{}

	return &result, nil
}
