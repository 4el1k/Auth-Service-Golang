package model

import "time"

//go:generate easyjson -all /home/bebra/repositories-ITIS/Avito-Test-Backend-Golang/internal/model/banner.go

//easyjson:json
type Banner struct {
	TagIds     []int     `json:"tag_ids"`
	Feature_id []int     `json:"feature_ids"`
	Content    Content   `json:"content"`
	Is_active  bool      `json:"is_active"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Banner_id  int       `json:"banner_id"`
	Article    int       `json:"article"`
}

//easyjson:json
type Content struct {
	Title string `json:"title"`
	Text  string `json:"text"`
	Url   string `json:"url"`
}

//easyjson:json
type BannerCreateOrUpdateRequest struct {
	TagIds    []int   `json:"tag_ids"`
	FeatureId int     `json:"feature_id"`
	Content   Content `json:"content"`
	IsActive  bool    `json:"is_active"`
}

//easyjson:json
type BannerFilterResponse struct {
	BannerId  int       `json:"banner_id"`
	TagIds    []int     `json:"tag_ids"`
	FeatureId int       `json:"feature_id"`
	Content   Content   `json:"content"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//easyjson:json
type BannerFilterGetRequest struct {
	TagIds    int `json:"tag_ids"`
	FeatureId int `json:"feature_id"`
	Limit     int `json:"limit"`
	Offset    int `json:"offset"`
}

//easyjson:json
type BannerGetRequest struct {
	TagIds    int `json:"tag_ids"`
	FeatureId int `json:"feature_id"`
}
