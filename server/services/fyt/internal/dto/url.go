package dto

type CreateUrlRequest struct {
	Url            string `json:"url" validate:"required,url"`
	ExpirationDate string `json:"expiration_date" validate:"omitempty,datetime=14-01-2026"`
	CustomAlias    string `json:"custom_alias" validate:"required,ascii"`
}
