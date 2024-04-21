package dto

type (
	GetSecretByUserIdResponse struct {
		UserId string `json:"user_id"`
		ApiKey string `json:"api_key"`
	}
)
