
# mysql

```go
/* *****************************************************************************
// Setup preferences
// ****************************************************************************/
// SET NAMES utf8 COLLATE 'utf8_unicode_ci';
// SET CHARACTER SET utf8;

/* *****************************************************************************
// Remove database (if it already exists)
// ****************************************************************************/
DROP DATABASE IF EXISTS gopherfacedb;

/* *****************************************************************************
// Create new database
// ****************************************************************************/
CREATE DATABASE gopherfacedb ;
USE gopherfacedb;

/* *****************************************************************************
// Create the table(s)
// ****************************************************************************/
CREATE TABLE user (
    id TINYINT(1) UNSIGNED NOT NULL AUTO_INCREMENT,
	username VARCHAR(18) NOT NULL,
	uuid VARCHAR(64) NOT NULL,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    password_hash CHAR(64) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_ts TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_ts TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	UNIQUE (username),
    PRIMARY KEY (id)
);

```

`定义模型`

```go
package models

import (
	"time"

	"github.com/EngineerKamesh/gofullstack/volume2/section5/gopherfacedb/common/utility"
)

type User struct {
	UUID              string `json:"uuid" bson:"uuid"`
	Username          string `json:"username" bson:"username"`
	FirstName         string `json:"firstName" bson:"firstName"`
	LastName          string `json:"lastName" bson:"lastName"`
	Email             string `json:"email" bson:"email"`
	PasswordHash      string `json:"passwordHash" bson:"passwordHash"`
	TimestampCreated  int64  `json:"timestampCreated" bson:"timestampCreated"`
	TimestampModified int64  `json:"timestampModified" bson:"timestampModified"`
}

func NewUser(username string, firstName string, lastName string, email string, password string) *User {

	passwordHash := utility.SHA256OfString(password)
	now := time.Now()
	unixTimestamp := now.Unix()
	u := User{UUID: utility.GenerateUUID(), Username: username, FirstName: firstName, LastName: lastName, Email: email, PasswordHash: passwordHash, TimestampCreated: unixTimestamp}
	return &u
}

```

`连库`

```go
package datastore

import (
	"errors"

	"github.com/EngineerKamesh/gofullstack/volume2/section5/gopherfacedb/models"
)

type Datastore interface {
	CreateUser(user *models.User) error
	GetUser(username string) (*models.User, error)
	Close()
}

const (
	MYSQL = iota
	MONGODB
	REDIS
)

func NewDatastore(datastoreType int, dbConnectionString string) (Datastore, error) {

	switch datastoreType {
	case MYSQL:
		return NewMySQLDatastore(dbConnectionString)
	case MONGODB:
		return NewMongoDBDatastore(dbConnectionString)
	case REDIS:
		return NewRedisDatastore(dbConnectionString)
	default:
		return nil, errors.New("The datastore you specified does not exist!")
	}

}

```


```go
package datastore

import (
	"database/sql"
	"log"

	"github.com/EngineerKamesh/gofullstack/volume2/section5/gopherfacedb/models"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDatastore struct {
	*sql.DB
}

func NewMySQLDatastore(dataSourceName string) (*MySQLDatastore, error) {

	connection, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &MySQLDatastore{
		DB: connection}, nil
}

func (m *MySQLDatastore) CreateUser(user *models.User) error {

	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO user(uuid, username, first_name, last_name, email, password_hash) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.UUID, user.Username, user.FirstName, user.LastName, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (m *MySQLDatastore) GetUser(username string) (*models.User, error) {

	stmt, err := m.Prepare("SELECT uuid, username, first_name, last_name, email, password_hash, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM user WHERE username = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(username)
	u := models.User{}
	err = row.Scan(&u.UUID, &u.Username, &u.FirstName, &u.LastName, &u.Email, &u.PasswordHash, &u.TimestampCreated, &u.TimestampModified)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return &u, err
}

func (m *MySQLDatastore) Close() {
	m.Close()
}

```

```go
初始化数据库
	db, err := datastore.NewDatastore(datastore.MYSQL, "gopherface:gopherface@/gopherfacedb")

	if err != nil {
		log.Print(err)
	}

	defer db.Close()

	env := common.Env{DB: db}

		r.Handle("/signup", handlers.SignUpHandler(&env)).Methods("GET", "POST")

```

```go
// 接收env 到请求里
func SignUpHandler(e *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		s := SignUpForm{}
		s.FieldNames = []string{"username", "firstName", "lastName", "email"}
		s.Fields = make(map[string]string)
		s.Errors = make(map[string]string)

		switch r.Method {

		case "GET":
			DisplaySignUpForm(w, r, &s)
		case "POST":
			ValidateSignUpForm(w, r, &s, e) // e
		default:
			DisplaySignUpForm(w, r, &s)
		}

	})
}
```

`使用`

```go
// 函数签名 接收mysql连接信息
func ValidateSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm, e *common.Env) {

// ...
	if len(s.Errors) > 0 {
    		DisplaySignUpForm(w, r, s)
    	} else {
    		ProcessSignUpForm(w, r, s, e) // e
    	}


```

```go
// ProcessSignUpForm
func ProcessSignUpForm(w http.ResponseWriter, r *http.Request, s *SignUpForm, e *common.Env) {

	u := models.NewUser(r.FormValue("username"), r.FormValue("firstName"), r.FormValue("lastName"), r.FormValue("email"), r.FormValue("password"))
	//fmt.Println("user: ", u)
	err := e.DB.CreateUser(u)

	if err != nil {
		log.Print(err)
	}

	user, err := e.DB.GetUser("EngineerKamesh") // 查用户
	if err != nil {
		log.Print(err)
	} else {
		fmt.Printf("Fetch User Result: %+v\n", user)
	}

	// Display form confirmation message
	DisplayConfirmation(w, r, s)

}

```