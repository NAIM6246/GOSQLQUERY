package models

type User struct {
	ID   int    ` json:"id"`
	NAME string `varchar(50);json:"name"`
}

type AddColumn struct {
	ColName string `json:"colname"`
}

func UserTabel() string {
	return "users"
}
