package conn

import (
	"GOSQL/models"
	"context"
	"database/sql"

	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type Conn struct{}

func ConnectionToDB() (*sql.DB, error) {
	connStr := "server=localhost;user id=sa;password=golangdb123456.;database=testGoSQLdb"
	return sql.Open("sqlserver", connStr)
}

func Insert(user models.User) /*string*/ (sql.Result, error) {
	sqlStatement := `
	INSERT INTO users (name)
	VALUES ( $1)`
	d, err := dbInstance.Exec(sqlStatement, user.NAME)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("user created")
	}
	return d, err
	//return fmt.Sprintf("insert into %s (name)\nvalues(%s);", models.UserTabel(), user.NAME)
}

//TO CREATE:
func createUserTable(db *sql.DB) {
	create := fmt.Sprintf("if not exists(select[name] from sys.tables where[name]='%s') create table %s (id int NOT NULL identity(1,1) primary key ,name varchar(50) not null)", models.UserTabel(), models.UserTabel())
	//fmt.Println(create)
	_, e := db.Query(create)
	if e != nil {
		fmt.Println(e)
		fmt.Println("herer")
	}
	defer db.Close()
}

//TO UPDATE TABLE:
func UpdateUserTable(db *sql.DB) {
	// alter := fmt.Sprintf("alter table %s\nalter column id int ", models.UserTabel())
	// db.Query(alter)

	sqlStatement := `
IF COL_LENGTH('users', 'hoby') IS NULL
BEGIN
    ALTER TABLE users
    ADD [hoby] varchar(100) 
END`
	ctx := context.TODO()
	_, err := db.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("column addition failed")
	}
	defer db.Close()
}

var dbInstance *sql.DB

func Connection() {
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
	defer dbInstance.Close()

	createUserTable(db)
	// updateUserTable(db)
}
