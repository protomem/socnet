package port

import "context"

type QueryFunc[Param, Result any] func(ctx context.Context, param Param) (Result, error)
