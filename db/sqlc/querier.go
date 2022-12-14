// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"context"
	"encoding/json"
)

type Querier interface {
	GetMostUsedRequest(ctx context.Context) (FizzbuzzStatistic, error)
	IncrementRequest(ctx context.Context, request json.RawMessage) (FizzbuzzStatistic, error)
}

var _ Querier = (*Queries)(nil)
