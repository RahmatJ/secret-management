package usecase

import (
	"fmt"
	"secret-management/internal/dto"
)

type SecretUsecase struct {
}

func NewSecretUsecase() *SecretUsecase {
	return &SecretUsecase{}
}

func (uc *SecretUsecase) GetSecretByUserId(userId string) (dto.GetSecretByUserIdResponse, error) {
	result := dto.GetSecretByUserIdResponse{
		UserId: userId,
		ApiKey: fmt.Sprintf("%s-key", userId),
	}

	return result, nil
}
