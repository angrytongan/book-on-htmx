package search

import "context"

type Repository interface {
	Term(ctx context.Context, term string) ([]string, error)
}
