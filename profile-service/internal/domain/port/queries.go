package port

import (
	"context"

	"github.com/google/uuid"
	"github.com/protomem/socnet/profile-service/internal/domain/model"
)

type (
	QueryFunc[Params any] func(ctx context.Context, params Params) error

	QueryWithResultFunc[Param, Result any] func(ctx context.Context, param Param) (Result, error)
)

type (
	GetUserQuery QueryWithResultFunc[uuid.UUID, model.User]
)

type (
	GetUserByNicknameQuery QueryWithResultFunc[string, model.User]
)

type (
	GetUserByEmailQuery QueryWithResultFunc[string, model.User]
)

type (
	FindUsersQuery QueryWithResultFunc[[]uuid.UUID, []model.User]
)

type (
	CreateUserParams struct {
		Nickname string
		Email    string
		Password string
	}

	CreateUserQuery QueryWithResultFunc[CreateUserParams, model.User]
)

type (
	UpdateUserByNicknameParams struct {
		Nickname string
		Email    string
		Password string
	}

	UpdateUserByNicknameQuery QueryWithResultFunc[Pair[string, UpdateUserByNicknameParams], model.User]
)

type (
	DeleteUserByNicknameQuery QueryFunc[string]
)
