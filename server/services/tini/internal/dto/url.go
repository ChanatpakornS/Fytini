package dto

type GetShortenUrlRequest struct {
	CustomAlias string `json:"custom_alias" validate:"required,ascii"`
}
