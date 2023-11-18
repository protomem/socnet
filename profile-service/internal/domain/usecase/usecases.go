package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/protomem/socnet/profile-service/internal/domain/model"
	"github.com/protomem/socnet/profile-service/internal/domain/port"
	"golang.org/x/crypto/bcrypt"
)

func GetUser(
	getUserQuery port.GetUserQuery,
) port.GetUserUseCase {
	return func(ctx context.Context, id uuid.UUID) (model.User, error) {
		const op = "usecase.GetUser"

		user, err := getUserQuery(ctx, id)
		if err != nil {
			return model.User{}, fmt.Errorf("%s: %w", op, err)
		}

		return user, nil
	}
}

func GetUserByNickname(
	getUserByNicknameQuery port.GetUserByNicknameQuery,
) port.GetUserByNicknameUseCase {
	return func(ctx context.Context, nickname string) (model.User, error) {
		const op = "usecase.GetUserByNickname"

		user, err := getUserByNicknameQuery(ctx, nickname)
		if err != nil {
			return model.User{}, fmt.Errorf("%s: %w", op, err)
		}

		return user, nil
	}
}

func GetUserByEmailAndPassword(
	getUserByEmailQuery port.GetUserByEmailQuery,
) port.GetUserByEmailAndPasswordUseCase {
	return func(ctx context.Context, dto port.GetUserByEmailAndPasswordDTO) (model.User, error) {
		const op = "usecase.GetUserByEmailAndPassword"
		var err error

		user, err := getUserByEmailQuery(ctx, dto.Email)
		if err != nil {
			return model.User{}, fmt.Errorf("%s: %w", op, err)
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(dto.Password))
		if err != nil {
			if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				return model.User{}, fmt.Errorf("%s: %w", op, model.ErrUserNotFound)
			}

			return model.User{}, fmt.Errorf("%s: %w", op, err)
		}

		return user, nil
	}
}

func FindUsers(
	findUsersQuery port.FindUsersQuery,
) port.FindUsersUseCase {
	return func(ctx context.Context, ids []uuid.UUID) ([]model.User, error) {
		const op = "usecase.FindUsers"

		users, err := findUsersQuery(ctx, ids)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		return users, nil
	}
}

func CreateUser(
	createUserQuery port.CreateUserQuery,
) port.CreateUserUseCase {
	return func(ctx context.Context, dto port.CreateUserDTO) (model.User, error) {
		const op = "usecase.CreateUser"

		// TODO: Validate ...

		hashPass, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			return model.User{}, fmt.Errorf("%s: %w", op, err)
		}

		user, err := createUserQuery(ctx, port.CreateUserParams{
			Nickname: dto.Nickname,
			Password: string(hashPass),
			Email:    dto.Email,
		})
		if err != nil {
			return model.User{}, fmt.Errorf("%s: %w", op, err)
		}

		return user, nil
	}
}

func UpdateUserByNickname(
	getUserByNicknameQuery port.GetUserByNicknameQuery,
	updateUserByNicknameQuery port.UpdateUserByNicknameQuery,
) port.UpdateUserByNicknameUseCase {
	return func(ctx context.Context, dto port.Pair[string, port.UpdateUserByNicknameDTO]) (model.User, error) {
		const op = "usecase.UpdateUserByNickname"
		var err error

		// TODO: Validate ...

		oldUser, err := getUserByNicknameQuery(ctx, dto.Left())
		if err != nil {
			return model.User{}, fmt.Errorf("%s: %w", op, err)
		}

		updateData := port.UpdateUserByNicknameParams{}

		if dto.Right().Nickname != nil {
			updateData.Nickname = dto.Right().Nickname
		}

		if dto.Right().Email != nil {
			updateData.Email = dto.Right().Email
			*updateData.Verified = false
		}

		if dto.Right().OldPassword != nil && dto.Right().NewPassword != nil {
			err = bcrypt.CompareHashAndPassword([]byte(oldUser.Password), []byte(*dto.Right().OldPassword))
			if err != nil && !errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				return model.User{}, fmt.Errorf("%s: %w", op, err)
			}

			updateData.Password = dto.Right().NewPassword
		}

		newUser, err := updateUserByNicknameQuery(ctx, port.NewPair(dto.Left(), updateData))
		if err != nil {
			return model.User{}, fmt.Errorf("%s: %w", op, err)
		}

		return newUser, nil
	}
}

func DeleteUserByNickname(
	deleteUserByNicknameQuery port.DeleteUserByNicknameQuery,
) port.DeleteUserByNicknameQuery {
	return func(ctx context.Context, nickname string) error {
		const op = "usecase.DeleteUserByNickname"

		err := deleteUserByNicknameQuery(ctx, nickname)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		return nil
	}
}
