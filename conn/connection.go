package conn

import (
	"database/sql"

	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func ConnectionToDB() (*sql.DB, error) {
	connStr := "server=localhost;user id=sa;password=golangdb123456.;database=testGoSQLdb"
	return sql.Open("sqlserver", connStr)
}

var dbInstance *sql.DB

func Connection() *sql.DB {
	db, err := ConnectionToDB()
	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Connection successfull")
	dbInstance = db
	// defer dbInstance.Close()

	createUserTable()
	createArticleTable()
	createCommentTable()

	// updateUserTable(db)
	return dbInstance
}
