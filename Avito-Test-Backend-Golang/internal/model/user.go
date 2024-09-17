package model

import uuid "github.com/satori/go.uuid"

//go:generate easyjson -all /home/bebra/repositories-ITIS/Avito-Test-Backend-Golang/internal/model/user.go

//easyjson:json
type User struct {
	Id       uuid.UUID `json:"id"`
	Login    string    `json:"login"`
	Password string    `json:"password"`
	IsAdmin  bool      `json:"is_admin"`
	TagId    int       `json:"tag_id"`
}

//easyjson:json
type UserSignInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

//easyjson:json
type UserRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

//easyjson:json
type UserCoupleTokenResponse struct {
	AccessToken  string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

//easyjson:json
type UserSignUpRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
	TagId    int    `json:"tag_id"`
}
