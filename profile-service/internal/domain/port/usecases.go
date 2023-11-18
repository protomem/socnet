package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/protomem/socnet/profile-service/internal/domain/model"
)

type (
	UseCaseFunc[DTO any] func(ctx context.Context, dto DTO) error

	UseCaseWithResultFunc[DTO, Result any] func(ctx context.Context, dto DTO) (Result, error)
)

type (
	GetUserUseCase UseCaseWithResultFunc[uuid.UUID, model.User]
)

type (
	GetUserByNicknameUseCase UseCaseWithResultFunc[string, model.User]
)

type (
	GetUserByEmailAndPasswordDTO struct {
		Email    string
		Password string
	}

	GetUserByEmailAndPasswordUseCase UseCaseWithResultFunc[GetUserByEmailAndPasswordDTO, model.User]
)

type (
	FindUsersUseCase UseCaseWithResultFunc[[]uuid.UUID, []model.User]
)

type (
	CreateUserDTO struct {
		Nickname string
		Email    string
		Password string
	}

	CreateUserUseCase UseCaseWithResultFunc[CreateUserDTO, model.User]
)

type (
	UpdateUserByNicknameDTO struct {
		Nickname *string
		Email    *string

		OldPassword *string
		NewPassword *string
	}

	UpdateUserByNicknameUseCase UseCaseWithResultFunc[Pair[string, UpdateUserByNicknameDTO], model.User]
)

type (
	DeleteUserByNicknameUseCase UseCaseFunc[string]
)

type UseCases struct {
	GetUserUseCase
	GetUserByNicknameUseCase
	GetUserByEmailAndPasswordUseCase

	FindUsersUseCase

	CreateUserUseCase

	UpdateUserByNicknameUseCase

	DeleteUserByNicknameUseCase
}
