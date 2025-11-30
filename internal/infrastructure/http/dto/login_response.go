package dto

type LoginResponse struct {
	Token *string `json:"token,omitempty"`
	Error *string `json:"error,omitempty"`
}

func NewRegisterResponseError(error string) *LoginResponse {
	return &LoginResponse{Token: nil, Error: &error}
}

func NewRegisterResponse(token string) *LoginResponse {
	return &LoginResponse{Token: &token, Error: nil}
}
