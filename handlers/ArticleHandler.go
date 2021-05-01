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

type ArticleHandler struct{}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{}
}

func (h *ArticleHandler) Handle(router chi.Router) {
	router.Get("/", h.getArticle)
	router.Post("/", h.createArticle)
	router.Route("/{id}", func(router chi.Router) {
		router.Get("/", h.getArticleByID)
		router.Delete("/", h.deleteArticle)
		router.Get("/comments", h.getCommentsOfArticle)
		router.Route("/{user}", func(router chi.Router) {
			router.Get("/", h.getCommentsOfUserInArticle)
		})
	})
}

func (h *ArticleHandler) createArticle(w http.ResponseWriter, r *http.Request) {
	input := models.Article{}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("content-type", "application/json")
		w.Write([]byte(`{"message" : "bad request error"}`))
		return
	}
	fmt.Println(input)
	d, e := conn.InsertArticle(input)
	if e != nil {
		fmt.Println("hi")
		json.NewEncoder(w).Encode(e)
		return
	}

	json.NewEncoder(w).Encode(d)
}

func (h *ArticleHandler) getArticle(w http.ResponseWriter, r *http.Request) {
	articles, err := conn.GetAllArticle()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(articles)
}

func (h *ArticleHandler) getArticleByID(w http.ResponseWriter, r *http.Request) {
	id := param.UInt(r, "id")
	// models.User
	article, err := conn.GetArticleByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("content-type", "application/json")
		w.Write([]byte(`{"message" : "article not found"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func (h *ArticleHandler) deleteArticle(w http.ResponseWriter, r *http.Request) {
	id := param.UInt(r, "id")
	conn.Delete(id, models.ArticleTable())
	w.WriteHeader(http.StatusNoContent)
	w.Header().Add("content-type", "application/json")
}

func (h *ArticleHandler) getCommentsOfArticle(w http.ResponseWriter, r *http.Request) {
	id := param.UInt(r, "id")
	comments, err := conn.GetAllCommentFromArticle(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(comments)
}

func (h *ArticleHandler) getCommentsOfUserInArticle(w http.ResponseWriter, r *http.Request) {
	id := param.UInt(r, "id")
	uid := param.UInt(r, "user")
	comments, err := conn.GetAllCommentFromArticleOfUser(id, uid)
	if err != nil {
		w.WriteHeader(400)
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode(comments)
}
