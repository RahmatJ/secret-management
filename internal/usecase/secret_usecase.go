package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"log"
	"secret-management/config"
	"secret-management/internal/constants"
	"secret-management/internal/domain"
	"secret-management/internal/dto"
	"secret-management/internal/entities"
	"time"
)

type SecretUsecase struct {
	name       string
	repository domain.SecretRepository
	helpers    domain.SecretHelpers
	config     config.EnvConfig
}

func NewSecretUsecase(repository domain.SecretRepository, helpers domain.SecretHelpers, config config.EnvConfig) *SecretUsecase {
	return &SecretUsecase{
		name:       "SecretUsecase",
		repository: repository,
		helpers:    helpers,
		config:     config,
	}
}

func generateUUID() string {
	id := uuid.New()
	return id.String()
}

func (uc *SecretUsecase) generateSecretPayload(userId string) entities.SecretManagement {
	apiKey := uc.helpers.GenerateSecret(20)
	ttl := time.Duration(uc.config.ApiKeyTTL)
	expiredDate := time.Now().UTC().Add(ttl * time.Second)
	id := generateUUID()

	return entities.SecretManagement{
		Id:          id,
		UserId:      userId,
		ApiKey:      apiKey,
		ExpiredDate: expiredDate,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
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

	input := uc.generateSecretPayload(userId)

	err = uc.repository.CreateSecret(&input)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%s.Err", funcName))
	}

	result.FromSecretEntities(&input)

	return &result, nil
}

func getStartEndDay(daytime time.Time) time.Time {
	location, _ := time.LoadLocation(constants.Constant.Localization)
	y, m, d := daytime.In(location).Date()
	return time.Date(y, m, d, 23, 59, 59, 0, location)
}

func (uc *SecretUsecase) DailySecretCheck() error {
	funcName := fmt.Sprintf("%s.DailySecretCheck", uc.name)

	format := time.RFC3339Nano
	expiryDate := time.Now().AddDate(0, 0, 7)
	endOfDay := getStartEndDay(expiryDate)

	expiringSecret, err := uc.repository.GetExpiringSecret(endOfDay.Format(format))
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("%s.Err", funcName))
	}

	fmt.Printf("Expiring: \n %+v", expiringSecret)
	//loop expiring secret, generating new secret
	for _, secret := range expiringSecret {
		fmt.Printf("Secret: %+v", secret)
		input := uc.generateSecretPayload(secret.UserId)

		err = uc.repository.CreateSecret(&input)
		if err != nil {
			log.Printf("%s.Err: Failed to create secret. Error: %+v ", funcName, err)
		}
	}

	return nil
}
