package models

type Comments struct {
	ID        uint   `json:"id"`
	BODY      string `json:"body"`
	USERID    uint   `json:"userID"`
	ARTICLEID uint   `json:"articleID"`
}

func CommentTable() string {
	return "comments"
}
