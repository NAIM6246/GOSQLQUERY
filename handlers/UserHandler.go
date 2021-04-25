package handlers

import (
	"GOSQL/conn"
	"GOSQL/models"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) Handle(router chi.Router) {
	router.Get("/", h.getUser)
	router.Post("/", h.createUser)
	router.Delete("/", h.deleteUser)
}

func insert(db *sql.DB, user models.User) (sql.Result, error) {
	sqlStatement := `
	INSERT INTO users (name)
	VALUES (@nam)`
	ctx := context.TODO()
	d, err := db.ExecContext(ctx, sqlStatement, sql.Named("nam", user.NAME))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("user created")
	}

	return d, err
}

//create
func (h *UserHandler) createUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(400)
		w.Header().Add("content-type", "application/json")
		w.Write([]byte(`{"message" : "bad request error"}`))
		return
	}
	fmt.Println(user)
	db, er := conn.ConnectionToDB()
	if er != nil {
		json.NewEncoder(w).Encode(er)
	}
	d, e := insert(db, user)
	if e != nil {
		fmt.Println("hi")
		json.NewEncoder(w).Encode(e)
	}

	json.NewEncoder(w).Encode(d)
}

func (h *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {

	db, er := conn.ConnectionToDB()
	if er != nil {
		json.NewEncoder(w).Encode(er)
	}
	conn.UpdateUserTable(db)
	fmt.Fprintf(w, "column added")
}

func (h *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete user")
}
