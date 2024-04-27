package dto

import (
	"secret-management/internal/entities"
	"time"
)

type (
	GetSecretByUserIdResponse struct {
		UserId      string    `json:"user_id"`
		ApiKey      string    `json:"api_key"`
		ExpiredDate time.Time `json:"expired_date"`
	}
)

func (g *GetSecretByUserIdResponse) FromSecretEntities(entities *entities.SecretManagement) {
	g.UserId = entities.UserId
	g.ApiKey = entities.ApiKey
	g.ExpiredDate = entities.ExpiredDate
}

func (g *GetSecretByUserIdResponse) ToSecretEntities(customId string, expiredDate time.Time) entities.SecretManagement {
	return entities.SecretManagement{
		Id:          customId,
		UserId:      g.UserId,
		ApiKey:      g.ApiKey,
		ExpiredDate: expiredDate,
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}
