package domain

import "secret-management/internal/dto"

type SecretUsecase interface {
	GetSecretByUserId(userId string) (dto.GetSecretByUserIdResponse, error)
}

type SecretRepository interface {
}
