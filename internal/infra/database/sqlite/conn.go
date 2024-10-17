package sqlite

import (
	"database/sql"
	"errors"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tclemos/go-expert-desafio-client-server-api/config"
)

const (
	driverName = "sqlite3"
)

func initDB(cfg config.SQLiteConfig) {
	dbFilePath, err := filepath.Abs(cfg.Path)
	if err != nil {
		panic(err)
	}
	if _, err := os.Stat(dbFilePath); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(dbFilePath)
		if err != nil {
			panic(err)
		}
		defer f.Close()
	}
}

func initTables(db *sql.DB) error {
	createCotacoesTableSQL := `CREATE TABLE IF NOT EXISTS cotacoes ("date" TIMESTAMP, "bid" TEXT);`

	statement, err := db.Prepare(createCotacoesTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = statement.Exec()
	return err
}

func MustOpenConn(cfg config.SQLiteConfig) *sql.DB {
	initDB(cfg)

	dbFilePath, err := filepath.Abs(cfg.Path)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(driverName, dbFilePath)
	if err != nil {
		panic(err)
	}

	err = initTables(db)
	if err != nil {
		panic(err)
	}

	return db
}
