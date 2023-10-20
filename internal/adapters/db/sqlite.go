package db

import (
	"database/sql"

	"github.com/sergiovirahonda/echo-api/internal/cfg"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func NewConnection(config cfg.Config) *bun.DB {
	conn, err := sql.Open("sqlite3", config.Database.File)
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(conn, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	return db
}

func NewTestConnection() *bun.DB {
	conn, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(conn, sqlitedialect.New())
	return db
}
