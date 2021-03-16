// migrate create -ext sql -dir db/migrations -seq create_items_table
/*
testing:
	- docker-compose up --build
    - example of export
        - POSTGRESQL_URL="
            postgres://
            ${POSTGRES_USER}:${POSTGRES_PASSWORD}
            @localhost:5432/
            ${POSTGRES_DB}?sslmode=disable"
    - export POSTGRESQL_URL="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable"─╯
    - migrate -database ${POSTGRESQL_URL} -path db/migrations up
	- curl -X POST \
	http://localhost:8080/items -H "Content-type: application/json" \
	-d '{ "id": "065d8403-8a8f-484d-b602-9138ff7dedcf", "name": "Wadson marcia", "username": "wadson.marcia"}'
*/
// migrate -database ${POSTGRESQL_URL} -path db/migrations up

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PicPay/software-engineer-challenge/db"
	"github.com/PicPay/software-engineer-challenge/handler"
)

func main() {
	addr := ":8080"
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")
	database, err := db.Initialize(dbUser, dbPassword, dbName)

	if err != nil {
		log.Fatalf("Could not set up database: %v", err)
	}
	defer database.Conn.Close()

	httpHandler := handler.NewHandler(database)
	server := &http.Server{
		Handler: httpHandler,
	}

	go func() {
		resp := handler.Pfile()
		handler.Parsing(resp)	
		server.Serve(listener)
	}()
	defer Stop(server)

	log.Printf("Started server on %s", addr)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")
}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
