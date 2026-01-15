export interface CreateShortenURLRequest {
  url: string;
  custom_alias: string;
  expiration_date?: string;
}

export interface CreateShortenURLResponse {
  message?: string;
  error?: string;
}

export interface GetShortenURLRequest {
  custom_alias: string;
}

export interface GetShortenURLResponse {
  url?: string;
  error?: string;
}
