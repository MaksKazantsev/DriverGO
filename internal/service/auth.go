package service

import (
	"context"
	"fmt"
	"github.com/MaksKazantsev/DriverGO/internal/errors"
	"github.com/MaksKazantsev/DriverGO/internal/repositories"
	"github.com/MaksKazantsev/DriverGO/internal/service/models"
	"github.com/MaksKazantsev/DriverGO/internal/utils"
	"github.com/google/uuid"
)

type Authorization interface {
	Register(ctx context.Context, data models.RegisterReq) (models.AuthResponse, error)
	Login(ctx context.Context, data models.LoginReq) (models.AuthResponse, error)
	Refresh(ctx context.Context, token string) (models.AuthResponse, error)
}

type Auth struct {
	repo repositories.Auth
}

func NewAuth(repo repositories.Auth) Authorization {
	return &Auth{repo: repo}
}

func (a *Auth) Register(ctx context.Context, data models.RegisterReq) (models.AuthResponse, error) {
	fmt.Println(data.Password)
	hash, err := utils.Hash(data.Password)
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("failed to hash password: %w", err)
	}
	data.Password = hash

	id := uuid.New().String()
	data.ID = id

	RToken, err := utils.NewToken(utils.REFRESH, utils.TokenData{ID: data.ID})
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}
	data.RToken = RToken

	AToken, err := utils.NewToken(utils.ACCESS, utils.TokenData{ID: data.ID, Role: "user"})
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "service layers successfully passed")
	res, err := a.repo.Register(ctx, UserReqToEntity(data))
	if err != nil {
		return models.AuthResponse{}, errors.ErrorRepoWrapper(err)
	}
	res.AccessToken = AToken

	return res, nil
}

func (a *Auth) Login(ctx context.Context, data models.LoginReq) (models.AuthResponse, error) {
	id, hash, err := a.repo.GetPasswordAndID(ctx, data.Email)
	if err != nil {
		return models.AuthResponse{}, errors.ErrorRepoWrapper(err)
	}

	if err = utils.CompareHash(hash, data.Password); err != nil {
		return models.AuthResponse{}, err
	}

	RToken, err := utils.NewToken(utils.REFRESH, utils.TokenData{ID: id})
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}
	data.RToken = RToken

	AToken, err := utils.NewToken(utils.ACCESS, utils.TokenData{ID: id, Role: "user"})
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "service layers successfully passed")

	res, err := a.repo.Login(ctx, data)
	if err != nil {
		return models.AuthResponse{}, errors.ErrorRepoWrapper(err)
	}

	res.AccessToken = AToken

	return res, nil
}

func (a *Auth) Refresh(ctx context.Context, token string) (models.AuthResponse, error) {
	cl, err := utils.ParseToken(token)
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("failed to parse token: %w", err)
	}

	id := cl["id"].(string)

	RToken, err := utils.NewToken(utils.REFRESH, utils.TokenData{ID: id})
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}

	AToken, err := utils.NewToken(utils.ACCESS, utils.TokenData{ID: id, Role: "user"})
	if err != nil {
		return models.AuthResponse{}, fmt.Errorf("failed to generate token: %w", err)
	}

	utils.ExtractLogger(ctx).Trace(utils.ExtractIdempotencyKey(ctx), "service layers successfully passed")

	res, err := a.repo.Refresh(ctx, id, RToken)
	if err != nil {
		return models.AuthResponse{}, errors.ErrorRepoWrapper(err)
	}

	res.AccessToken = AToken

	return res, nil
}
