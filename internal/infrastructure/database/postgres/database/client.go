package database

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func New(ctx context.Context) (*PGDatabase, error) {
	dsn := generateDsn()
	log.Println()
	pool, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
	return NewDatabase(pool), nil
}

func generateDsn() string {
	connData := getConnectData()
	log.Println("::::::::::::::::::::::::::::::::::::::::::::::::::")
	log.Println("host=", connData.host)
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		connData.host, connData.port, connData.user, connData.password, connData.dbName)
}
