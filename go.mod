// to install migration use this (download)
// go get -u github.com/golang-migrate/migrate/v4

// to create the migration you have to use
// migrate create -ext psql -dir migrations init
// and for to do migration up you have to use
// migrate -database "postgres://postgres:1234@localhost/todo?sslmode=disable" -path migrations up

module github.com/AmitPaliwal1810/go-lang-todo

go 1.22.1

require github.com/jmoiron/sqlx v1.3.5

require (
	github.com/go-chi/chi/v5 v5.0.12 // indirect
	github.com/golang-migrate/migrate/v4 v4.17.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/crypto v0.22.0 // indirect
)
