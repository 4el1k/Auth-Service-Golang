package repo

import (
	"Avito-Test-Backend-Golang/internal/model"
	"context"
	"database/sql"
	"log"
)

func NewBannerRepo(log *log.Logger, db *sql.DB) *BannerRepo {
	return &BannerRepo{log, db}
}

type BannerRepo struct {
	log *log.Logger
	db  *sql.DB
}

func (r *BannerRepo) GetContentByTagIdAndFeatureId(ctx context.Context,
	req model.BannerGetRequest) (model.Content, error) {
	return (nil), nil
}

func (r *BannerRepo) GetPageByTagIdAndFeatureId(ctx context.Context,
	req model.BannerFilterGetRequest) ([]model.BannerFilterResponse, error) {
	return nil, nil
}

func (r *BannerRepo) Create(ctx context.Context, req model.BannerCreateOrUpdateRequest) (int, error) {
	return 0, nil
}

func (r *BannerRepo) Update(ctx context.Context, req model.BannerCreateOrUpdateRequest) error {
	return nil
}

func (r *BannerRepo) Delete(ctx context.Context, id int) error {
	return nil
}
