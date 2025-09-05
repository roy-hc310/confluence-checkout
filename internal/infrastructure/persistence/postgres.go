package persistence

import (
	"context"
	"fmt"
	"time"

	"confluence-checkout/internal/infrastructure/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type PostgresInfraHandler struct {
	DbRead  *pgxpool.Pool
	DbWrite *pgxpool.Pool
}

func NewPostgresInfraHandler() *PostgresInfraHandler {
	dsnRead := fmt.Sprintf("host=%s user=%s password=%s dbname=%s search_path=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.GlobalEnv.DBReadHost, config.GlobalEnv.DBReadUser, config.GlobalEnv.DBReadPass, config.GlobalEnv.DBReadName, config.GlobalEnv.DBReadSchema, config.GlobalEnv.DBReadPort)

	dbReadConfig, err := pgxpool.ParseConfig(dsnRead)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	dbReadConfig.MaxConns = 10
	dbReadConfig.MinConns = 0
	dbReadConfig.MaxConnLifetime = time.Hour
	dbReadConfig.MaxConnIdleTime = time.Minute * 30
	dbReadConfig.HealthCheckPeriod = time.Minute * 5

	dbRead, err := pgxpool.NewWithConfig(context.Background(), dbReadConfig)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := dbRead.Ping(ctx); err != nil {
		log.Fatal().Msg(err.Error())
		println("Can't Connect to Postgres (Read)")
	}

	// readQuery := sqlc.New(dbReadPool)

	dsnWrite := fmt.Sprintf("host=%s user=%s password=%s dbname=%s search_path=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.GlobalEnv.DBWriteHost, config.GlobalEnv.DBWriteUser, config.GlobalEnv.DBWritePass, config.GlobalEnv.DBWriteName, config.GlobalEnv.DBWriteSchema, config.GlobalEnv.DBWritePort)

	dbWriteConfig, err := pgxpool.ParseConfig(dsnWrite)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	dbWriteConfig.MaxConns = 10
	dbWriteConfig.MinConns = 0
	dbWriteConfig.MaxConnLifetime = time.Hour
	dbWriteConfig.MaxConnIdleTime = time.Minute * 30
	dbWriteConfig.HealthCheckPeriod = time.Minute * 5

	dbWrite, err := pgxpool.NewWithConfig(context.Background(), dbWriteConfig)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	if err := dbWrite.Ping(ctx); err != nil {
		log.Fatal().Msg(err.Error())
		println("Can't Connect to Postgres (Read)")
	}

	// writeQuery := sqlc.New(dbWritePool)

	return &PostgresInfraHandler{
		DbRead:  dbRead,
		DbWrite: dbWrite,
	}
}

func (p *PostgresInfraHandler) BeginTx(ctx context.Context, options pgx.TxOptions) (pgx.Tx, error) {
	tx, err := p.DbWrite.BeginTx(ctx, options)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (p *PostgresInfraHandler) Query(ctx context.Context, queryString string, params []interface{}) (pgx.Rows, error) {
	rows, err := p.DbRead.Query(ctx, queryString, params...)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (p *PostgresInfraHandler) TotalItems(totalItems *string, ctx context.Context, queryString string, params []interface{}) error {
	err := p.DbRead.QueryRow(ctx, queryString, params...).Scan(&totalItems)
	if err != nil {
		return err
	}

	return nil
}
