package transport

import (
	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
	"github.com/protomem/socnet/user-service/internal/domain/model"
	"github.com/protomem/socnet/user-service/internal/domain/port"
)

var ErrInvalidRequest = errors.New("invalid request")

type Endpoints struct {
	GetUser                   endpoint.Endpoint
	GetUserByNickname         endpoint.Endpoint
	GetUserByEmailAndPassword endpoint.Endpoint
	FindUsers                 endpoint.Endpoint
	CreateUser                endpoint.Endpoint
	UpdateUserByNickname      endpoint.Endpoint
	DeleteUserByNickname      endpoint.Endpoint
}

func MakeEndpoints(ucs port.UseCases) Endpoints {
	return Endpoints{
		GetUser:                   MakeGetUserEndpoint(ucs.GetUserUseCase),
		GetUserByNickname:         MakeGetUserByNicknameEndpoint(ucs.GetUserByNicknameUseCase),
		GetUserByEmailAndPassword: MakeGetUserByEmailAndPasswordEndpoint(ucs.GetUserByEmailAndPasswordUseCase),
		FindUsers:                 MakeFindUsersEndpoint(ucs.FindUsersUseCase),
		CreateUser:                MakeCreateUserEndpoint(ucs.CreateUserUseCase),
		UpdateUserByNickname:      MakeUpdateUserByNicknameEndpoint(ucs.UpdateUserByNicknameUseCase),
		DeleteUserByNickname:      MakeDeleteUserByNicknameEndpoint(ucs.DeleteUserByNicknameUseCase),
	}
}

type (
	GetUserRequest struct {
		ID uuid.UUID
	}

	GetUserResponse struct {
		User model.User `json:"user"`

		Err error `json:"error,omitempty"`
	}
)

func MakeGetUserEndpoint(uc port.GetUserUseCase) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req, ok := request.(GetUserRequest)
		if !ok {
			return nil, ErrInvalidRequest
		}

		user, err := uc(ctx, req.ID)
		if err != nil {
			return GetUserResponse{Err: err}, nil
		}

		return GetUserResponse{User: user}, nil
	}
}

type (
	GetUserByNicknameRequest struct {
		Nickname string
	}

	GetUserByNicknameResponse struct {
		User model.User `json:"user"`

		Err error `json:"error,omitempty"`
	}
)

func MakeGetUserByNicknameEndpoint(uc port.GetUserByNicknameUseCase) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req, ok := request.(GetUserByNicknameRequest)
		if !ok {
			return nil, ErrInvalidRequest
		}

		user, err := uc(ctx, req.Nickname)
		if err != nil {
			return GetUserByNicknameResponse{Err: err}, nil
		}

		return GetUserByNicknameResponse{User: user}, nil
	}
}

type (
	GetUserByEmailAndPasswordRequest struct {
		Email    string
		Password string
	}

	GetUserByEmailAndPasswordResponse struct {
		User model.User `json:"user"`

		Err error `json:"error,omitempty"`
	}
)

func MakeGetUserByEmailAndPasswordEndpoint(uc port.GetUserByEmailAndPasswordUseCase) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req, ok := request.(GetUserByEmailAndPasswordRequest)
		if !ok {
			return nil, ErrInvalidRequest
		}

		user, err := uc(ctx, port.GetUserByEmailAndPasswordDTO{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			return GetUserByEmailAndPasswordResponse{Err: err}, nil
		}

		return GetUserByEmailAndPasswordResponse{User: user}, nil
	}
}

type (
	FindUsersRequest struct {
		IDs []uuid.UUID
	}

	FindUsersResponse struct {
		Users []model.User `json:"users"`

		Err error `json:"error,omitempty"`
	}
)

func MakeFindUsersEndpoint(uc port.FindUsersUseCase) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req, ok := request.(FindUsersRequest)
		if !ok {
			return nil, ErrInvalidRequest
		}

		users, err := uc(ctx, req.IDs)
		if err != nil {
			return FindUsersResponse{Err: err}, nil
		}

		return FindUsersResponse{Users: users}, nil
	}
}

type (
	CreateUserRequest struct {
		Nickname string
		Email    string
		Password string
	}

	CreateUserResponse struct {
		User model.User `json:"user"`

		Err error `json:"error,omitempty"`
	}
)

func MakeCreateUserEndpoint(uc port.CreateUserUseCase) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req, ok := request.(CreateUserRequest)
		if !ok {
			return nil, ErrInvalidRequest
		}

		user, err := uc(ctx, port.CreateUserDTO(req))
		if err != nil {
			return CreateUserResponse{Err: err}, nil
		}

		return CreateUserResponse{User: user}, nil
	}
}

type (
	UpdateUserByNicknameRequest struct {
		Nickname string

		Data struct {
			Nickname    *string
			Email       *string
			OldPassword *string
			NewPassword *string
		}
	}

	UpdateUserByNicknameResponse struct {
		User model.User `json:"user"`

		Err error `json:"error,omitempty"`
	}
)

func MakeUpdateUserByNicknameEndpoint(uc port.UpdateUserByNicknameUseCase) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req, ok := request.(UpdateUserByNicknameRequest)
		if !ok {
			return nil, ErrInvalidRequest
		}

		user, err := uc(ctx, port.NewPair(req.Nickname, port.UpdateUserByNicknameDTO(req.Data)))
		if err != nil {
			return UpdateUserByNicknameResponse{Err: err}, nil
		}

		return UpdateUserByNicknameResponse{User: user}, nil
	}
}

type (
	DeleteUserByNicknameRequest struct {
		Nickname string
	}

	DeleteUserByNicknameResponse struct {
		Err error `json:"error,omitempty"`
	}
)

func MakeDeleteUserByNicknameEndpoint(uc port.DeleteUserByNicknameUseCase) endpoint.Endpoint {
	return func(ctx context.Context, request any) (any, error) {
		req, ok := request.(DeleteUserByNicknameRequest)
		if !ok {
			return nil, ErrInvalidRequest
		}

		err := uc(ctx, req.Nickname)
		if err != nil {
			return DeleteUserByNicknameResponse{Err: err}, nil
		}

		return DeleteUserByNicknameResponse{}, nil
	}
}
