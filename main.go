package main

import (
	"GOSQL/conn"
	"GOSQL/handlers"
	"fmt"
	"net/http"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/go-chi/chi"
)

// func ConnectionToDB() (*sql.DB, error) {
// 	connStr := "server=localhost;user id=sa;password=golangdb123456.;database=testGoSQLdb"
// 	return sql.Open("sqlserver", connStr)
// }

// //TO CREATE:
// func createUserTable(db *sql.DB) {
// 	create := fmt.Sprintf("create table %s (id int NOT NULL identity(1,1) primary key ,name varchar(50) not null)", models.UserTabel())
// 	_, e := db.Query(create)
// 	if e != nil {
// 		fmt.Println(e)
// 	}
// 	defer db.Close()
// }

// //TO UPDATE TABLE:
// func updateUserTable(db *sql.DB) {
// 	alter := fmt.Sprintf("alter table %s\nalter column id int ", models.UserTabel())
// 	db.Query(alter)
// 	defer db.Close()
// }

// var dbInstance *sql.DB

func main() {
	fmt.Println("HELLO")
	//connStr := "server=localhost;user id=sa;password=golangdb123456.;database=testGoSQLdb"
	// db, err := conn.ConnectionToDB() //ConnectionToDB()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer db.Close()
	conn.Connection()
	// createUserTable(db)
	// updateUserTable(db)

	router := chi.NewRouter()
	userHandler := handlers.NewUserHandler()
	router.Route("/users", userHandler.Handle)

	http.ListenAndServe(":8000", router)
}
