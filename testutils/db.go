package testutils

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewDBForTest(t *testing.T) *sqlx.DB {
	t.Helper()

	mysqlDB, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/root?parseTime=true")

	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(
		func() { _ = mysqlDB.Close() },
	)

	return sqlx.NewDb(mysqlDB, "mysql")
}
