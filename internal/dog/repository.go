package dog

import "context"

type Repository interface {
	Colours(ctx context.Context) ([]string, error)
	Breeds(ctx context.Context) ([]string, error)
	All(ctx context.Context, colour, breed, orderBy string, limit int) ([]Dog, error)
}
