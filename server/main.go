package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"entdemo/config"
	"entdemo/ent"
	resolvers "entdemo/resolvers"

	_ "entdemo/ent/runtime"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
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
	router := chi.NewRouter()
	config.InitConfig()
	dbUrl := config.GetDbUrl()

	// Create ent.Client and run the schema migration.
	client, err := Open(dbUrl)
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	defer client.Close()

	// Configure the server and start listening on :8080.
	srv := handler.NewDefaultServer(resolvers.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})

	createdProduct, err := client.Product.Create().
		SetName("test").SetDescription("test").SetPrice(100.0).Save(context.Background())
	if err != nil {
		panic(err)
	}
	log.Println("created product id: ", createdProduct.ID)
	product, err := client.Product.Query().All(context.Background())
	log.Println("err: ", err)
	if err != nil {
		log.Println("err: ", err)
	}
	for i := range product {
		log.Println("product: ", product[i])
	}

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Handle("/", playground.Handler("Graphql", "/graphql"))
	router.Handle("/graphql", srv)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
