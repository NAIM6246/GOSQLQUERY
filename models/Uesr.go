package models

type User struct {
	ID   int    ` json:"id"`
	NAME string `json:"name"`
}

type AddColumn struct {
	ColName string `json:"colname"`
}

func UserTabel() string {
	return "users"
}
