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
