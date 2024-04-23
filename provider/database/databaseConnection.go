package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type SSLMode string

var (
	Todo *sqlx.DB
)

const (
	SSLModeEnable  SSLMode = "enable"
	SSLModeDisable SSLMode = "disable"
)

func DBConnection(host, port, databaseName, user, password string, sslMode SSLMode) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, databaseName, sslMode) // string formatter
	fmt.Println(connStr)
	var DB *sqlx.DB
	DB, err := sqlx.Open("postgres", connStr)

	if err != nil {
		return DB, err
	}

	err = DB.Ping()

	if err != nil {
		return DB, err
	}

	err = migrateUp(DB)
	if err != nil {
		return DB, err
	}

	return DB, nil
}

func migrateUp(db *sqlx.DB) error {
	driver, err := postgres.WithInstance(db.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://provider/database/migrations",
		"todo", driver)

	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

// this is transaction function
func Tx(fn func(tx *sqlx.Tx) error) error {
	tx, err := Todo.Beginx() // TODO is a global variable which has the type of sqlx.DB

	if err != nil {
		fmt.Print("getting error while hitting the transaction  function")
		return err
	}

	// write defere function for -> if any thing goes wrong it will help to rollback the database.
	// and defere function runs while all the surrounding function executes.
	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(); rollBackErr != nil {
				fmt.Print("rollback error in TX funcion", rollBackErr)
			}
			return

		}
		if commitErr := tx.Commit(); commitErr != nil {
			fmt.Print("commit error in TX function", commitErr)
		}
	}()

	err = fn(tx)

	return err
}
