package repo

import (
	"Avito-Test-Backend-Golang/internal/model"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgtype/pgxtype"
	"github.com/jackc/pgx/v4"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"time"
)

type AuthRepository struct {
	db              pgxtype.Querier
	log             *logrus.Logger
	refreshTokenTtl time.Duration
}

func NewAuthRepository(db pgxtype.Querier, log *logrus.Logger, ttlMilisec time.Duration) *AuthRepository {
	return &AuthRepository{
		db:              db,
		log:             log,
		refreshTokenTtl: ttlMilisec,
	}
}

const (
	createToken     = "INSERT INTO refresh_token (user_id, token_id, exp_date) VALUES ($1, $2, $3)"
	deleteToken     = "DELETE FROM refresh_token WHERE user_id = $1 AND token = $2"
	findUserByLogin = "SELECT id, password FROM account WHERE login = $1"
	findUserById    = "SELECT login, tag_id, is_admin FROM account WHERE id = $1"
	checkToken      = "SELECT 1 FROM refresh_token WHERE user_id = $1 AND token_id = $2 AND exp_date<NOW()"
	createUser      = "INSERT INTO account (login, password, is_admin) VALUES ($1, $2, $3)"
)

func (r *AuthRepository) Create(ctx context.Context, userId uuid.UUID) (*uuid.UUID, error) {
	r.log.Printf("start create token")
	refreshToken := uuid.NewV4()
	expDate := time.Now().Add(r.refreshTokenTtl).UTC()
	_, err := r.db.Exec(ctx, createToken, userId, refreshToken, expDate)
	if err != nil {
		r.log.Printf("error creating token: %v", err)
		return nil, err
	}
	return &refreshToken, nil
}

func (r *AuthRepository) Delete(ctx context.Context, userId uuid.UUID, refreshToken uuid.UUID) error {
	r.log.Printf("start delete refreshToken")
	_, err := r.db.Exec(ctx, deleteToken, userId, refreshToken)
	if err != nil {
		r.log.Printf("error deleting refreshToken: %v", err)
		return err
	}
	return nil
}

func (r *AuthRepository) CreateUser(ctx context.Context, user *model.UserSignUpRequest) error {
	r.log.Println("start CreateUser auth repo")
	_, err := r.db.Exec(ctx, createUser, user.Login, []byte(user.Password), user.IsAdmin)
	if err != nil {
		r.log.Printf("error creating user: %v", err)
		return err
	}
	return nil
}

func (r *AuthRepository) IsRottenToken(ctx context.Context, userId uuid.UUID, refreshToken uuid.UUID) (bool, error) {
	r.log.Println("start checkToken")
	var count int64
	row := r.db.QueryRow(ctx, checkToken, userId, refreshToken)
	err := row.Scan(&count)
	if err != nil {
		r.log.Printf("error checking token: %v", err)
		return false, nil
	}
	r.log.Println("end checkToken")
	return count > 0, nil
}

func (r *AuthRepository) FindUserByLogin(ctx context.Context, login string) (*model.User, error) {
	r.log.Println("start findByLogin auth repo")
	row := r.db.QueryRow(ctx, findUserByLogin, login)
	result := &model.User{Login: login}
	var pass []byte
	err := row.Scan(&result.Id, &pass)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = fmt.Errorf("error happened in db.QueryRow, func findByLogin: %w", err)
			r.log.Errorf("error in auth repo FindUserByLogin", err.Error())
			return &model.User{}, err
		}
	}
	result.Password = string(pass)
	r.log.Printf("end findByLogin auth repo, result: {id: %s} \n", result.Id)
	return result, nil
}

func (r *AuthRepository) FindUserById(ctx context.Context, id uuid.UUID) (*model.User, error) {
	r.log.Println("start findByLogin refresh token repo")
	row := r.db.QueryRow(ctx, findUserById, id)
	result := &model.User{Id: id}
	err := row.Scan(&result.Login, &result.IsAdmin, result.TagId)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			err = fmt.Errorf("error happened in db.QueryRow, func findById: %w", err)
			r.log.Errorf("error in auth repo, err.Error())")
			return &model.User{}, err
		}
	}
	r.log.Println("end findById auth repo, result: {id: %s}", result.Id)
	return result, nil
}
