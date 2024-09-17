package service

import (
	"Avito-Test-Backend-Golang/internal/model"
	"Avito-Test-Backend-Golang/internal/pkg/auth/repo"
	"Avito-Test-Backend-Golang/internal/util/hasher"
	"Avito-Test-Backend-Golang/internal/util/jwter"
	"context"
	"errors"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

type AuthService struct {
	authRepo   *repo.AuthRepository
	log        *logrus.Logger
	jwtManager *jwter.JwtManager
}

func NewAuthService(tokensRepo *repo.AuthRepository, log *logrus.Logger, jwtManager *jwter.JwtManager) *AuthService {
	return &AuthService{
		authRepo:   tokensRepo,
		jwtManager: jwtManager,
		log:        log,
	}
}

var (
	ErrPassMismatch = errors.New("password does not match")
)

func (s *AuthService) GetCoupleToken(ctx context.Context, in *model.UserSignInRequest) (*model.UserCoupleTokenResponse, error) {
	s.log.Println("start auth service")
	user, err := s.authRepo.FindUserByLogin(ctx, in.Login)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}
	println(user.Password)
	if !hasher.CheckPass([]byte(user.Password), in.Password) {
		s.log.Errorf(ErrPassMismatch.Error())
		return nil, ErrPassMismatch
	}
	return s.createCoupleToken(ctx, user)
}

func (s *AuthService) RefreshJwtToken(ctx context.Context, refreshToken uuid.UUID) (string, error) {
	s.log.Println("start refresh token service")
	userId := ctx.Value("userId").(uuid.UUID)
	isRottenToken, err := s.authRepo.IsRottenToken(ctx, userId, refreshToken)
	if err != nil {
		s.log.Println("error while checking if token is rotten")
		return "", err
	}
	if isRottenToken {
		s.log.Println("token is rotten")
		return "", err
	}
	user, err := s.authRepo.FindUserById(ctx, userId)
	newJwtToken, err := s.jwtManager.EncodeToken(user)
	if err != nil {
		s.log.Println("error while generating jwt token")
		return "", err
	}
	return newJwtToken, nil
}

func (s *AuthService) Logout(ctx context.Context, refreshToken uuid.UUID) error {
	s.log.Println("start logout service")
	userId := ctx.Value("userId").(uuid.UUID)
	err := s.authRepo.Delete(ctx, userId, refreshToken)
	if err != nil {
		s.log.Error(err)
		return err
	}
	return nil
}

func (s *AuthService) CreateUser(ctx context.Context, in *model.UserSignUpRequest) error {
	s.log.Println("start create user")
	_, err := s.authRepo.FindUserByLogin(ctx, in.Login)
	if err == nil {
		s.log.Println("user already exists")
		return errors.New("user already exists")
	}
	in.Password = string(hasher.HashPass(in.Password))
	err = s.authRepo.CreateUser(ctx, in)
	if err != nil {
		s.log.Println("error create user", err)
		return err
	}
	s.log.Println("end create user")
	return nil
}

func (s *AuthService) createCoupleToken(ctx context.Context, user *model.User) (*model.UserCoupleTokenResponse, error) {
	token, err := s.jwtManager.EncodeToken(user)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.authRepo.Create(ctx, user.Id)
	if err != nil {
		return nil, err
	}
	return &model.UserCoupleTokenResponse{AccessToken: token, RefreshToken: refreshToken.String()}, nil
}
