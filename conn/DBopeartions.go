package conn

import (
	"GOSQL/models"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
)

// const db = dbInstance

//TO CREATE:
func createUserTable() {
	// var db = dbInstance

	// create := fmt.Sprintf("if not exists(select[name] from sys.tables where[name]='%s') create table %s (id int NOT NULL identity(1,1) primary key ,name varchar(50) not null, address varchar(50))", models.UserTabel(), models.UserTabel())
	//fmt.Println(create)
	// _, e := db.Query(create)
	// if e != nil {
	// 	fmt.Println(e)
	// 	fmt.Println("herer")
	// }
	// fmt.Println("hi")
	// defer db.Close()
	sqlStastement := `
if not exists(select[name] from sys.tables where[name]= 'users' ) 
create table "users" (id int NOT NULL identity(1,1) primary key ,name varchar(50) not null)
	`
	ctx := context.TODO()
	_, e := dbInstance.ExecContext(ctx, sqlStastement, sql.Named("tablename", models.UserTabel()), sql.Named("tab", models.UserTabel()))
	if e != nil {
		fmt.Println(e)
		fmt.Println("herer")
	}
	// defer db.Close()
}

//TO CREATE article:
func createArticleTable( /*db *sql.DB*/ ) {
	// create := fmt.Sprintf("if not exists(select[name] from sys.tables where[name]='%s') create table %s (id int NOT NULL identity(1,1) primary key ,name varchar(50) not null)", models.UserTabel(), models.UserTabel())
	//fmt.Println(create)
	sqlStastement := `
if not exists(select[name] from sys.tables where[name]='@tablename') create table @tab (id int NOT NULL identity(1,1) primary key ,name varchar(50) not null)
	`

	// var db = dbInstance
	ctx := context.TODO()
	_, e := dbInstance.ExecContext(ctx, sqlStastement, sql.Named("tablename", "articles"), sql.Named("tab", "articles"))
	if e != nil {
		fmt.Println(e)
		fmt.Println("herer")
	}
	// defer db.Close()
}

//TO UPDATE TABLE:
func UpdateUserTable() {
	// var db = dbInstance
	sqlStatement := `
IF COL_LENGTH('users', 'hoby') IS NULL
BEGIN
    ALTER TABLE users
    ADD [hoby] varchar(100) 
END`
	ctx := context.TODO()
	_, err := dbInstance.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("column addition failed")
	}
	// defer db.Close()
}

//create data in table
func Insert(user models.User) (sql.Result, error) {
	// var db = dbInstance

	//return sql result bt cant convert to struct!!
	sqlStatement := `
	INSERT INTO users 
	VALUES (@nam)`
	ctx := context.TODO()
	//sql.Naned working for nam bt not working for table name!!why?
	d, err := dbInstance.ExecContext(ctx, sqlStatement /*sql.Named("table", "users"),*/, sql.Named("nam", user.NAME))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("user created")
	}
	//err2 := jsonify.Jsonify(d)
	return d, err
}

func GetByID(id uint) (*models.User, error) {
	fmt.Println(id)
	sqlStatement := fmt.Sprintf("SELECT  id as id, name as \"name\" from %s where id=%d for json path;", models.UserTabel(), id)
	ctx := context.TODO()
	fmt.Println(sqlStatement)
	// d, err := dbInstance.QueryContext(ctx, sqlStatement, sql.Named("id", id))
	d := dbInstance.QueryRowContext(ctx, sqlStatement, sql.Named("ids", id))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }

	var u []*models.User
	var s string
	er := d.Scan(&s)
	if er != nil {
		fmt.Println(er)
		return nil, er
	}
	// for d.Next() {
	// 	er := d.Scan(&s)
	// 	fmt.Println(er)
	// }
	b := []byte(s)

	err2 := json.Unmarshal(b, &u)
	if err2 != nil {
		fmt.Println(err2)
		return nil, err2
	}
	fmt.Println(u)

	return u[0], nil
}

//Get all user using normal sql query
func GetAll() ([]*models.User, error) {
	// var db = dbInstance
	sqlStatement := `select * from users`
	ctx := context.TODO()

	d, err := dbInstance.QueryContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer d.Close()
	var users []*models.User
	var user models.User
	for d.Next() {
		err2 := d.Scan(&user.ID, &user.NAME)
		if err2 != nil {
			fmt.Println(err2)
			fmt.Println("here")
			return nil, err
		}
		// fmt.Println(user)
		users = append(users, &user)
	}
	// fmt.Println(&users)
	return users, nil
	// return d, nil
}

//get all using json
func GetALL() ([]*models.User, error) {
	// var db = dbInstance
	sqlStatement := `
	SELECT  id , name as "name"
	from users
	for json path;
	`
	ctx := context.TODO()

	d, err := dbInstance.QueryContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var u []*models.User
	var s string
	for d.Next() {
		er := d.Scan(&s)
		if er != nil {
			return nil, er
		}
	}
	b := []byte(s)
	//
	err2 := json.Unmarshal(b, &u)
	if err2 != nil {
		fmt.Println(err2)
		return nil, err2
	}
	fmt.Println(&u)

	return u, nil
	// return d, nil
}
