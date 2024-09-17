package http

import (
	"Avito-Test-Backend-Golang/internal/pkg/banner/service"
	"github.com/sirupsen/logrus"
	"net/http"
)

type BannerHandler struct {
	log     *logrus.Logger
	service service.BannerService
}

func NewBannerHandler(service service.BannerService, log *logrus.Logger) *BannerHandler {
	return &BannerHandler{log, service}
}

func (h *BannerHandler) GetBannerContent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *BannerHandler) GetPageBanner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *BannerHandler) PostBanner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *BannerHandler) PatchBanner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func (h *BannerHandler) DeleteBanner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
