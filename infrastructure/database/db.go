package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/syougo1209/b-match-server/config"
)

func NewDB(ctx context.Context, cfg *config.Config) (*sqlx.DB, func(), error) {
	mysqlDB, err := sql.Open("mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?parseTime=true&loc=Local",
			cfg.MySQLUser, cfg.MYSQLRootPassword,
			cfg.MYSQLAddr, cfg.MYSQLDbName,
		),
	)
	if err != nil {
		return nil, nil, err
	}

	if err := mysqlDB.PingContext(ctx); err != nil {
		return nil, nil, err
	}
	xdb := sqlx.NewDb(mysqlDB, "mysql")
	return xdb, func() { _ = mysqlDB.Close() }, nil
}

// sqlxの使用するメソッドを定義 ref: https://pkg.go.dev/github.com/jmoiron/sqlx#section-readme
type DbConnection interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	SelectContext(ctx context.Context, dest interface{}, query string, args ...any) error
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}
