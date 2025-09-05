package persistence

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type DatabaseInfra interface {
	BeginTx(ctx context.Context, options pgx.TxOptions) (pgx.Tx, error)
	Query(ctx context.Context, queryString string, params []interface{}) (pgx.Rows, error)
	TotalItems(totalItems *string, ctx context.Context, queryString string, params []interface{}) error
}
