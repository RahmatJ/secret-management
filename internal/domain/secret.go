package domain

import (
	"secret-management/internal/dto"
	"secret-management/internal/entities"
)

type SecretUsecase interface {
	GetSecretByUserId(userId string) (*dto.GetSecretByUserIdResponse, error)
	DailySecretCheck() error
}

type SecretRepository interface {
	CreateSecret(secret *entities.SecretManagement) error
	GetSecret(userId string, currentTime string) (*entities.SecretManagement, error)
	GetExpiringSecret(endTime string) ([]entities.SecretManagement, error)
}

type SecretHelpers interface {
	GenerateSecret(length int) string
}
