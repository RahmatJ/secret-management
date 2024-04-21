package entities

import "time"

type SecretManagement struct {
	Id          string    `json:"id" gorm:"id"`
	UserId      string    `json:"user_id" gorm:"user_id"`
	ApiKey      string    `json:"api_key" gorm:"api_key"`
	ExpiredDate time.Time `json:"expired_date" gorm:"expired_date"`
	CreatedAt   time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
}
