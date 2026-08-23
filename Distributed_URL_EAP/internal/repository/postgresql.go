package repository

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(
		context.Background(),
		"postgres://myuser:distributedSQL@localhost:5433/mydatabase",
	)

	if err != nil {
		return nil, err
	}

	slog.Info("postgreSql connection success...")
	return conn, nil
}

func QueryData(conn *pgx.Conn) {
	rows, err := conn.Query(context.Background(), "SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("User ID: %d, Name: %s\n", id, name)
	}
}
