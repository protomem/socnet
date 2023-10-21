package port

import (
	"context"
)

type UseCaseFunc[DTO, Result any] func(ctx context.Context, dto DTO) (Result, error)
