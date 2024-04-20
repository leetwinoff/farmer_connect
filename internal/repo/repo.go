package repo

import (
	"context"
	"database/sql"
	"farmers_connect/internal/config"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"time"
)

type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	GetContext(context.Context, interface{}, string, ...interface{}) error
	SelectContext(context.Context, interface{}, string, ...interface{}) error

	BeginTx(ctx context.Context) (Tx, error)
	Close(ctx context.Context) error
}

type Tx interface {
	Commit() error
	Rollback() error

	SelectContext(context.Context, interface{}, string, ...any) error
	ExecContext(context.Context, string, ...any) (sql.Result, error)
	GetContext(context.Context, interface{}, string, ...any) error
}

type TxStarter interface {
	BeginTx(ctx context.Context) (Tx, error)
}

func NewDB(ctx context.Context, cfg config.DB) (*dbWrapper, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SSLMode)
	pgxCfg, err := pgx.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}
	pgxCfg.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	rawDB := stdlib.OpenDB(*pgxCfg)
	db := sqlx.NewDb(rawDB, "pgx")
	db = db.Unsafe()
	db.SetMaxOpenConns(cfg.PoolSize)
	db.SetMaxIdleConns(cfg.PoolSize)

	logger := zerolog.Ctx(ctx)
	logger.Trace().
		Int("pool_size", cfg.PoolSize).
		Msg("start ping db")
	pingCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := db.PingContext(pingCtx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	logger.Trace().Msg("ping db success")

	return &dbWrapper{DB: db}, nil
}

type dbWrapper struct {
	DB *sqlx.DB
}

func (db *dbWrapper) Close(ctx context.Context) {
	if err := db.DB.Close(); err != nil {
		log.Error().Err(err).Msg("failed to close database")
	}
}
