package handlers

import (
	"GOSQL/conn"
	"GOSQL/handlers/param"
	"GOSQL/models"
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
	router.Route("/{id}", func(router chi.Router) {
		router.Get("/", h.getUserByID)
	})
	router.Get("/", h.getUser)
	router.Post("/", h.createUser)
	router.Put("/", h.updateUserTable)
	router.Delete("/", h.deleteUser)
}

//create : response e prblm ase!struct e convert korte partesi na sql.result re
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
	d, e := conn.Insert(user)
	if e != nil {
		fmt.Println("hi")
		json.NewEncoder(w).Encode(e)
		return
	}

	json.NewEncoder(w).Encode(d)
}

func (h *UserHandler) getUser(w http.ResponseWriter, r *http.Request) {

	// users, err := conn.GetAll()
	users, err := conn.GetALL()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(&users)
	json.NewEncoder(w).Encode(users)

}

func (h *UserHandler) getUserByID(w http.ResponseWriter, r *http.Request) {
	id := param.UInt(r, "id")
	// models.User
	user, err := conn.GetByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("content-type", "application/json")
		w.Write([]byte(`{"message" : "user not found"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) updateUserTable(w http.ResponseWriter, r *http.Request) {
	col := models.AddColumn{}
	err := json.NewDecoder(r.Body).Decode(&col)
	if err != nil {
		w.WriteHeader(400)
		w.Header().Add("content-type", "application/json")
		w.Write([]byte(`{"message" : "bad request error"}`))
		return
	}
	fmt.Println(col)

	conn.UpdateUserTable()
	fmt.Fprintf(w, "column added")
}

func (h *UserHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete user")
}
