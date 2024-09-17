package service

import (
	"Avito-Test-Backend-Golang/internal/model"
	"Avito-Test-Backend-Golang/internal/pkg/banner/repo"
	"context"
	"github.com/sirupsen/logrus"
)

func NewBannerService(log *logrus.Logger, repo repo.BannerRepo) *BannerService {
	return &BannerService{log, repo}
}

type BannerService struct {
	log  *logrus.Logger
	repo repo.BannerRepo
}

func (s *BannerService) GetBannerContent(ctx context.Context, req model.BannerGetRequest) (model.Content, error) {
	return (nil), nil
}

func (s *BannerService) GetBannerPage(ctx context.Context,
	req model.BannerFilterGetRequest) ([]model.BannerFilterResponse, error) {
	return nil, nil
}

func (s *BannerService) CreateBanner(ctx context.Context, req model.BannerCreateOrUpdateRequest) (int, error) {
	return 0, nil
}

func (s *BannerService) UpdateBanner(ctx context.Context, req model.BannerCreateOrUpdateRequest) error {
	return nil
}

func (s *BannerService) DeleteBanner(ctx context.Context, id int) error {
	return nil
}
