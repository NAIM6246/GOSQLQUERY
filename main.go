package main

import (
	"GOSQL/conn"
	"GOSQL/handlers"
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/chi"
)

func main() {
	db := conn.Connection()

	router := chi.NewRouter()
	userHandler := handlers.NewUserHandler()
	router.Route("/users", userHandler.Handle)
	fmt.Println("serving on port :8080")
	http.ListenAndServe(":8000", router)

	defer db.Close()
}
