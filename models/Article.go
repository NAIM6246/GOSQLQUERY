package models

type Article struct {
	ID     uint   `json:"id"`
	BODY   string `json:"body"`
	USERID uint   `json:"userID"`
}

func ArticleTable() string {
	return "articles"
}
