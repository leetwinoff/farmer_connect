package repo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"farmers_connect/internal/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
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
		logger.Error().Err(err).Msg("failed to ping db")
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}
	logger.Trace().Msg("ping db success")

	return &dbWrapper{DB: db}, nil
}

type dbWrapper struct {
	DB *sqlx.DB
}

func (db *dbWrapper) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return db.DB.ExecContext(ctx, query, args...)
}

func (db *dbWrapper) GetContext(ctx context.Context, dest interface{}, query string, args ...any) error {
	return db.DB.GetContext(ctx, dest, query, args...)
}

func (db *dbWrapper) SelectContext(ctx context.Context, dest interface{}, query string, args ...any) error {
	return db.DB.SelectContext(ctx, dest, query, args...)
}

func (db *dbWrapper) BeginTx(ctx context.Context) (Tx, error) {
	tx, err := db.DB.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (db *dbWrapper) Close(ctx context.Context) error {
	if err := db.DB.Close(); err != nil {
		return fmt.Errorf("failed to close db: %w", err)
	}
	return nil
}
