package http

import (
	"Avito-Test-Backend-Golang/internal/model"
	"Avito-Test-Backend-Golang/internal/pkg/auth/service"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func NewAuthHandler(log *logrus.Logger, authService *service.AuthService) *AuthHandler {
	return &AuthHandler{log, authService}
}

type AuthHandler struct {
	log     *logrus.Logger
	service *service.AuthService
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	h.log.Printf("start signin")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Printf("read body err: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	in := &model.UserSignInRequest{}
	err = in.UnmarshalJSON(body) // ToDo: easy JSON!!
	if err != nil {
		h.log.Printf("unmarshal body err: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	coupleTokens, err := h.service.GetCoupleToken(r.Context(), in)
	if err != nil {
		h.log.Printf("auth err: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	resp, err := coupleTokens.MarshalJSON()
	if err != nil {
		h.log.Printf("marshal body err: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		h.log.Printf("write err: %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	h.log.Printf("start signUp")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Printf("read body err: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	payload := &model.UserSignUpRequest{}
	err = payload.UnmarshalJSON(body) // ToDo: easy JSON!!
	if err != nil {
		h.log.Printf("unmarshal body err: %s", err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = h.service.CreateUser(r.Context(), payload)
	if err != nil {
		h.log.Printf("sign up err: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	h.log.Printf("start refresh token")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.log.Printf("read body err: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	token := &model.UserRefreshTokenRequest{}
	err = token.UnmarshalJSON(body)
	if err != nil {
		h.log.Printf("unmarshal body err: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	refreshToken, err := uuid.FromString(token.RefreshToken)
	if err != nil {
		h.log.Printf("refresh token err: %s", err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	jwtToken, err := h.service.RefreshJwtToken(r.Context(), refreshToken)
	if err != nil {
		h.log.Printf("auth err: %s", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(jwtToken)
	if err != nil {
		return
	}
	_, err = w.Write(resp)
	if err != nil {
		h.log.Printf("write err: %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	h.log.Printf("start logout")
	// ToDo: поручаю это задание стажеру :))
}
