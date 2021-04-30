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

	sqlStastement := fmt.Sprintf("if not exists(select[name] from sys.tables where[name]='%s') create table %s (id int NOT NULL identity(1,1) primary key ,name varchar(50) not null)", models.UserTabel(), models.UserTabel())
	fmt.Println(sqlStastement)
	// 	sqlStastement := `
	// if not exists(select[name] from sys.tables where[name]= 'users' )
	// create table "users" (id int NOT NULL identity(1,1) primary key ,name varchar(50) not null)
	// 	`
	ctx := context.TODO()
	_, e := dbInstance.ExecContext(ctx, sqlStastement)
	if e != nil {
		fmt.Println(e)
		fmt.Println("herer")
	}
}

//TO UPDATE TABLE:
func UpdateUserTable() {
	// var db = dbInstance
	sqlStatement := fmt.Sprintf("IF COL_LENGTH('%s', 'hoby') IS NULL BEGIN ALTER TABLE %s ADD [hoby] varchar(100) END", models.UserTabel(), models.UserTabel())
	// 	sqlStatement := `
	// IF COL_LENGTH('users', 'hoby') IS NULL
	// BEGIN
	//     ALTER TABLE users
	//     ADD [hoby] varchar(100)
	// END`
	ctx := context.TODO()
	_, err := dbInstance.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println("column addition failed")
	}
}

//delete table data by id
func Delete(id uint, tableName string) {
	sqlStatement := fmt.Sprintf("DELETE FROM %s WHERE ID=%d", tableName, id)
	_, err := dbInstance.Exec(sqlStatement)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(d)
}

//create data in table
func Insert(user models.User) (sql.Result, error) {
	//return sql result bt cant convert to struct!!
	sqlStatement := fmt.Sprintf("INSERT INTO %s VALUES ('%s')", models.UserTabel(), user.NAME)
	ctx := context.TODO()
	d, err := dbInstance.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("user created")
	}
	return d, err
}

func GetByID(id uint) (*models.User, error) {

	sqlStatement := fmt.Sprintf("SELECT  id as id, name as \"name\" from %s where id=%d for json path;", models.UserTabel(), id)
	ctx := context.TODO()
	var s string
	err := dbInstance.QueryRowContext(ctx, sqlStatement).Scan(&s)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	/*
		var s string
		er := d.Scan(&s)
		if er != nil {
			fmt.Println(er)
			return nil, er
		}
	*/

	var u []*models.User
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
	sqlStatement := fmt.Sprintf("SELECT * FROM %s", models.UserTabel())
	ctx := context.TODO()

	d, err := dbInstance.QueryContext(ctx, sqlStatement)
	if err != nil {
		return nil, err
	}
	defer d.Close()
	var users []*models.User
	var user models.User
	for d.Next() {
		err2 := d.Scan(&user.ID, &user.NAME)
		if err2 != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

//get all using json
func GetALL() ([]*models.User, error) {
	// var db = dbInstance
	sqlStatement := fmt.Sprintf("SELECT  id , name as \"name\" from %s for json path;", models.UserTabel())
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
}

//TO CREATE article:
func createArticleTable() {
	sqlStastement := fmt.Sprintf("IF NOT EXISTS(SELECT[name] FROM sys.tables WHERE[name]='%s') CREATE TABLE %s (id int NOT NULL identity(1,1) primary key, body text not null, userID int not null)", models.ArticleTable(), models.ArticleTable())
	fmt.Println(sqlStastement)
	// 	sqlStastement := `
	// if not exists(select[name] from sys.tables where[name]='@tablename') create table @tab (id int NOT NULL identity(1,1) primary key ,name varchar(50) not null)
	// 	`

	// var db = dbInstance
	ctx := context.TODO()
	_, e := dbInstance.ExecContext(ctx, sqlStastement)
	if e != nil {
		fmt.Println(e)
		fmt.Println("herer")
	}
}

func InsertArticle(article models.Article) (sql.Result, error) {
	sqlStatement := fmt.Sprintf("INSERT INTO %s (body,userID) VALUES ('%s',%d)", models.ArticleTable(), article.BODY, article.USERID)
	ctx := context.TODO()
	d, err := dbInstance.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("article created")
	}
	return d, err
}

func GetArticleByID(id uint) (*models.Article, error) {

	sqlStatement := fmt.Sprintf("SELECT  id , body as body, userID as userID  from %s where id=%d for json path;", models.ArticleTable(), id)
	ctx := context.TODO()
	var s string
	err := dbInstance.QueryRowContext(ctx, sqlStatement).Scan(&s)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var u []*models.Article
	b := []byte(s)
	err2 := json.Unmarshal(b, &u)
	if err2 != nil {
		fmt.Println(err2)
		return nil, err2
	}
	fmt.Println(u)

	return u[0], nil
}

//get all article
func GetAllArticle() ([]*models.Article, error) {
	// var db = dbInstance
	sqlStatement := fmt.Sprintf("SELECT  id , body as body, userID as userID from %s for json path;", models.ArticleTable())
	ctx := context.TODO()

	d, err := dbInstance.QueryContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var article []*models.Article
	var s string
	for d.Next() {
		er := d.Scan(&s)
		if er != nil {
			return nil, er
		}
	}
	b := []byte(s)
	//
	err2 := json.Unmarshal(b, &article)
	if err2 != nil {
		fmt.Println(err2)
		return nil, err2
	}
	fmt.Println(&article)

	return article, nil
}

//To create comment table
func createCommentTable() {
	sqlStastement := fmt.Sprintf("IF NOT EXISTS(SELECT[name] FROM sys.tables WHERE[name]='%s') CREATE TABLE %s (id int NOT NULL identity(1,1) primary key, body text not null, userID int not null, articleID int not null)", models.CommentTable(), models.CommentTable())
	ctx := context.TODO()
	_, e := dbInstance.ExecContext(ctx, sqlStastement)
	if e != nil {
		fmt.Println(e)
		fmt.Println("herer")
	}
}

//create comment
func InsertComment(comment models.Comments) (sql.Result, error) {
	sqlStatement := fmt.Sprintf("INSERT INTO %s (body,userID,articleID) VALUES ('%s',%d,%d)", models.CommentTable(), comment.BODY, comment.USERID, comment.ARTICLEID)
	ctx := context.TODO()
	d, err := dbInstance.ExecContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		fmt.Println("comment created")
	}
	return d, err
}

//get comment by id
func GetCommentByID(id uint) (*models.Comments, error) {

	sqlStatement := fmt.Sprintf("SELECT  id , body as body, userID as userID, articleID as articleID from %s where id=%d for json path;", models.CommentTable(), id)
	ctx := context.TODO()
	var s string
	err := dbInstance.QueryRowContext(ctx, sqlStatement).Scan(&s)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var comment []*models.Comments
	b := []byte(s)
	err2 := json.Unmarshal(b, &comment)
	if err2 != nil {
		fmt.Println(err2)
		return nil, err2
	}

	return comment[0], nil
}

//get all comment
func GetAllComment() ([]*models.Comments, error) {
	// var db = dbInstance
	sqlStatement := fmt.Sprintf("SELECT  id , body as body, userID as userID, articleID as articleID from %s for json path;", models.CommentTable())
	ctx := context.TODO()

	d, err := dbInstance.QueryContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var comments []*models.Comments
	var s string
	for d.Next() {
		er := d.Scan(&s)
		if er != nil {
			return nil, er
		}
	}
	b := []byte(s)
	//
	err2 := json.Unmarshal(b, &comments)
	if err2 != nil {
		fmt.Println(err2)
		return nil, err2
	}
	fmt.Println(&comments)

	return comments, nil
}

//get all comments of a article
func GetAllCommentFromArticle(id uint) ([]*models.Comments, error) {
	// var db = dbInstance
	sqlStatement := fmt.Sprintf("SELECT  id , body as body, userID as userID,articleID as articleID from %s where articleID=%d for json path;", models.CommentTable(), id)
	ctx := context.TODO()

	d, err := dbInstance.QueryContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var comments []*models.Comments
	var s string
	for d.Next() {
		er := d.Scan(&s)
		if er != nil {
			return nil, er
		}
	}
	b := []byte(s)
	//
	err2 := json.Unmarshal(b, &comments)
	if err2 != nil {
		fmt.Println(err2)
		return nil, err2
	}

	return comments, nil
}

func GetAllCommentFromArticleOfUser(articleID uint, userID uint) ([]*models.Comments, error) {
	// var db = dbInstance
	sqlStatement := fmt.Sprintf("SELECT  id , body as body, userID as userID,articleID as articleID from %s where articleID=%d and userID=%d for json path;", models.CommentTable(), articleID, userID)
	ctx := context.TODO()

	d, err := dbInstance.QueryContext(ctx, sqlStatement)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var comments []*models.Comments
	var s string
	for d.Next() {
		er := d.Scan(&s)
		if er != nil {
			return nil, er
		}
	}
	b := []byte(s)
	//
	err2 := json.Unmarshal(b, &comments)
	if err2 != nil {
		fmt.Println(err2)
		return nil, err2
	}

	return comments, nil
}
