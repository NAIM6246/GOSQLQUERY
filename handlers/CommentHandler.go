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

type CommnetHandler struct{}

func NewCommnetHandler() *CommnetHandler {
	return &CommnetHandler{}
}

func (h *CommnetHandler) Handle(router chi.Router) {
	router.Get("/", h.getCommnet)
	router.Post("/", h.createCommnet)
	router.Route("/{id}", func(router chi.Router) {
		router.Get("/", h.getCommnetByID)
		router.Delete("/", h.deleteCommnet)
	})
}

func (h *CommnetHandler) createCommnet(w http.ResponseWriter, r *http.Request) {
	input := models.Comments{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(400)
		w.Header().Add("content-type", "application/json")
		w.Write([]byte(`{"message" : "bad request error"}`))
		return
	}
	fmt.Println(input)
	d, e := conn.InsertComment(input)
	if e != nil {
		fmt.Println("hi")
		json.NewEncoder(w).Encode(e)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(d)
}

func (h *CommnetHandler) getCommnet(w http.ResponseWriter, r *http.Request) {
	articles, err := conn.GetAllComment()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(articles)
}

func (h *CommnetHandler) getCommnetByID(w http.ResponseWriter, r *http.Request) {
	id := param.UInt(r, "id")
	// models.User
	comment, err := conn.GetCommentByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("content-type", "application/json")
		w.Write([]byte(`{"message" : "comment not found"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(comment)
}

func (h *CommnetHandler) deleteCommnet(w http.ResponseWriter, r *http.Request) {
	id := param.UInt(r, "id")
	conn.Delete(id, models.CommentTable())
	w.WriteHeader(http.StatusNoContent)
	w.Header().Add("content-type", "application/json")
}
