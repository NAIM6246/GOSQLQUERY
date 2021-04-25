package models

type User struct {
	ID   int    ` json:"id"`
	NAME string `varchar(50);json:"name"`
}

func UserTabel() string {
	return "users"
}
