package usecases

import "context"

type Synchronizer interface {
	Synchronize(ctx context.Context, groups []int32, userId uint) error
}
