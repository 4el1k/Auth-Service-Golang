package app

import (
	middleware "Avito-Test-Backend-Golang/internal/middlware"
	ah "Avito-Test-Backend-Golang/internal/pkg/auth/http"
	ar "Avito-Test-Backend-Golang/internal/pkg/auth/repo"
	as "Avito-Test-Backend-Golang/internal/pkg/auth/service"
	"Avito-Test-Backend-Golang/internal/pkg/config"
	"Avito-Test-Backend-Golang/internal/util/jwter"
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	logger *logrus.Logger
}

func NewApp(logger *logrus.Logger) *App {
	return &App{logger: logger}
}

func (a *App) Run() error {
	// ============================Config============================ //
	println("run start")
	wd, _ := os.Getwd()
	println(wd)
	cfg := config.MustLoad("config/config.yaml")
	// ============================Log============================ //
	loggger := a.logger
	// ============================DataBase============================ //
	loggger.Println("init database")
	//format := "postgres://%v:%v@%v:%v/%v?sslmode=disable"
	db, err := pgxpool.Connect(context.Background(), fmt.Sprintf(
		"user=%v password=%v host=%v port=%v dbname=%v",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName))
	if err != nil {
		err = fmt.Errorf("init database error %w", err)
		loggger.Error(err)
		return err
	}
	defer db.Close()
	if err = db.Ping(context.Background()); err != nil {
		err = fmt.Errorf("ping database error %w", err)
		loggger.Error(err)
		return err
	}
	// ============================Service================================== //
	loggger.Println("init service")
	jwtManager := jwter.NewJwtManager(cfg.AuthJwt.AccessExpirationTime, cfg.AuthJwt.Secret, cfg.Issuer)
	authRepo := ar.NewAuthRepository(db, loggger, cfg.AuthJwt.RefreshTokenTimeToLimit)
	authService := as.NewAuthService(authRepo, loggger, jwtManager)
	authHandler := ah.NewAuthHandler(loggger, authService)
	authMw := middleware.New(loggger, jwtManager)
	r := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.Handle("/logout", authMw(http.HandlerFunc(authHandler.Logout))).Methods(http.MethodPost)
		auth.Handle("/refresh", authMw(http.HandlerFunc(authHandler.RefreshToken))).Methods(http.MethodPost)
		auth.Handle("/sign-up", http.HandlerFunc(authHandler.SignUp)).Methods(http.MethodPost)
		auth.Handle("/sign-in", http.HandlerFunc(authHandler.SignIn)).Methods(http.MethodPost)
	}

	// ============================Server============================ //
	loggger.Println("init server")
	loggger.Println("addr: %s", cfg.Address)
	log1 := log.Default()
	srv := http.Server{
		Handler:           r,
		Addr:              cfg.Address,
		ReadTimeout:       cfg.Timeout,
		WriteTimeout:      cfg.Timeout,
		IdleTimeout:       cfg.IdleTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
		ErrorLog:          log1,
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("http server listen err: %v\n", err)
		}
	}()
	sig := <-quit
	a.logger.Debug("handle quit chanel: ", sig.String())
	a.logger.Info("server stopping...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		err = fmt.Errorf("error on server shutdown: %w", err)
		return err
	}
	return nil
}
