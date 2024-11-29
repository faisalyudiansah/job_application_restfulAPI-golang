package services

import (
	"context"

	"job-application/apperrors"
	"job-application/dtos"
	"job-application/helpers"
	"job-application/models"
	"job-application/repositories"
)

type UserService interface {
	PostRegisterUserService(context.Context, dtos.RequestRegisterUser) (*dtos.ResponseRegisterUserAndUserProfile, error)
	PostLoginUserService(context.Context, dtos.RequestLoginUser) (*dtos.ResponseAccessToken, error)
}

type UserServiceImplementation struct {
	UserRepository         repositories.UserRepository
	TransactionsRepository repositories.TransactionRepository
	Bcrypt                 helpers.Bcrypt
	Jwt                    helpers.JWTProvider
}

func NewUserServiceImplementation(us repositories.UserRepository, tx repositories.TransactionRepository, bcr *helpers.BcryptStruct, jwt *helpers.JwtProviderHS256) *UserServiceImplementation {
	return &UserServiceImplementation{
		UserRepository:         us,
		TransactionsRepository: tx,
		Bcrypt:                 bcr,
		Jwt:                    jwt,
	}
}

func (us *UserServiceImplementation) PostRegisterUserService(ctx context.Context, reqBody dtos.RequestRegisterUser) (*dtos.ResponseRegisterUserAndUserProfile, error) {
	result, err := us.TransactionsRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		if IsEmailAlreadyRegistered := us.UserRepository.IsEmailAlreadyRegistered(cForTx, reqBody.Email); IsEmailAlreadyRegistered {
			return nil, apperrors.ErrUserEmailAlreadyExists
		}
		hashPassword, err := us.Bcrypt.HashPassword(reqBody.Password, 10)
		if err != nil {
			return nil, apperrors.ErrUserFailedRegister
		}
		user, err := us.UserRepository.PostUser(cForTx, reqBody, string(hashPassword))
		if err != nil {
			return nil, apperrors.ErrUserFailedRegister
		}
		userProfile, err := us.UserRepository.PostUserProfile(cForTx, reqBody, user.ID)
		if err != nil {
			return nil, apperrors.ErrUserFailedRegister
		}
		dataUser := models.UserAndProfile{
			User:        *user,
			UserProfile: *userProfile,
		}
		return dataUser, nil
	})
	if err != nil {
		return nil, err
	}
	return helpers.ModelToDtoResponseRegister(result.(models.UserAndProfile)), nil
}

func (us *UserServiceImplementation) PostLoginUserService(ctx context.Context, reqBody dtos.RequestLoginUser) (*dtos.ResponseAccessToken, error) {
	result, err := us.TransactionsRepository.Atomic(ctx, func(cForTx context.Context) (any, error) {
		user, err := us.UserRepository.GetUserByEmail(cForTx, reqBody.Email)
		if err != nil {
			return nil, apperrors.ErrUserInvalidEmailPassword
		}
		isValid, err := us.Bcrypt.CheckPassword(reqBody.Password, []byte(user.Password))
		if err != nil || !isValid {
			return nil, apperrors.ErrUserInvalidEmailPassword
		}
		accessToken, err := us.Jwt.CreateToken(int64(user.ID), user.Role)
		if err != nil {
			return nil, apperrors.ErrISE
		}
		return accessToken, nil
	})
	if err != nil {
		return nil, err
	}
	return helpers.AccessTokenToDtoResponseUserAccessToken(result.(string)), nil
}
