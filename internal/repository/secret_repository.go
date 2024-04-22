package repository

import (
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"secret-management/internal/entities"
	"time"
)

type SecretRepository struct {
	name string
	db   *gorm.DB
}

func NewSecretRepository(db *gorm.DB) *SecretRepository {
	return &SecretRepository{
		name: "SecretRepository",
		db:   db,
	}
}

func (repo *SecretRepository) CreateSecret(secret *entities.SecretManagement) error {
	funcName := fmt.Sprintf("%s.CreateSecret", repo.name)

	err := repo.db.Create(&secret).Error
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("%s.Err", funcName))
	}

	return nil
}

func (repo *SecretRepository) GetSecret(userId string, currentTime time.Time) (*entities.SecretManagement, error) {
	funcName := fmt.Sprintf("%s.GetSecret", repo.name)

	var result entities.SecretManagement

	err := repo.db.Table(result.TableName()).
		Order("updated_at desc").
		Where("user_id = ? and expired_date < ?", userId, currentTime).
		First(&result).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("%s.Err", funcName))
	}

	return &result, nil
}
