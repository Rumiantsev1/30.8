package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"30.8/pkg/storage"

	"github.com/jackc/pgx/v4"
)

func main() {
	dsn := "postgresql://postgres:postgres@localhost:5434/postgres"
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	db, err := storage.New(dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database NEWWW: %v\n", err)
		os.Exit(1)
	}
	log.Println(db.Tasks(1, 1))
}
