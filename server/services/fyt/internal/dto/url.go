package dto

type CreateUrlRequest struct {
	Url            string `json:"url" validate:"required,url"`
	ExpirationDate string `json:"expiration_date" validate:"omitempty"`
	CustomAlias    string `json:"custom_alias" validate:"required,ascii"`
}
