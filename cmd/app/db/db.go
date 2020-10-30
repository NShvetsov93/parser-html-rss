package db

import (
	"context"
	"log"
	"os"

	// "github.com/jackc/pgx/v4"

	"github.com/jackc/pgx/v4/pgxpool"
)

func Connect() *pgxpool.Pool {
	login, exists := os.LookupEnv("DB_USER")
	if !exists {
		log.Fatalf("No found DB_USER")
	}
	pass, exists := os.LookupEnv("DB_PASS")
	if !exists {
		log.Fatalf("No found DB_PASS")
	}
	url, exists := os.LookupEnv("DB_URL")
	if !exists {
		log.Fatalf("No found DB_URL")
	}
	dbName, exists := os.LookupEnv("DB_NAME")
	if !exists {
		log.Fatalf("No found DB_NAME")
	}
	// conn, error := pgx.Connect(context.Background(), "postgres://"+login+":"+pass+"@"+url+":5432/"+dbName)
	conn, error := pgxpool.Connect(context.Background(), "postgres://"+login+":"+pass+"@"+url+":5432/"+dbName)
	if error != nil {
		log.Fatalf("Can't connect to db : %v", error)
	}
	// error = conn.Ping(context.Background())
	// if error != nil {
	// 	log.Fatalf("can't ping db: %v", error)
	// }
	return conn
}

func Close(conn *pgxpool.Pool) {
	conn.Close()
	// if error := conn.Close(context.Background()); error != nil {
	// 	log.Fatalf("Can't disconnect to db")
	// }
}