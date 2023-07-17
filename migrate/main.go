package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	"entdemo/ent"
	"entdemo/ent/migrate"

	config "entdemo/config"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func Open(databaseUrl string) (*ent.Client, error) {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv)), nil
}

func main() {
	//env
	config.InitConfig()
	dbUrl := config.GetDbUrl()

	// Create ent.Client and run the schema migration.
	client, err := Open(dbUrl)
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	defer client.Close()

	if err := client.Debug().Schema.Create(
		context.Background(),
		schema.WithAtlas(true),
		migrate.WithDropIndex(true),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

}
