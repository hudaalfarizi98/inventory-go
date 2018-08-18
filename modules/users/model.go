package users

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

type LoginRequest struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type LoginResponse struct {
	Result string `json:"res"`
}
type RegisterRequest struct {
	Username  string `form:"username"`
	Password  string `form:"password"`
	FirstName string `form:"firstname"`
	LastName  string `form:"lastname"`
}
type RegisterResponse struct {
	Result string `json:"res"`
}

func dbConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_db_inventory")

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func (lr *LoginRequest) ValidateFields() error {
	if lr.Username == "" || lr.Password == "" {
		return fmt.Errorf("Missing Username or Password")
	}

	return nil
}
func (rr *RegisterRequest) ValidateFields() error {
	if rr.Username == "" || rr.Password == "" {
		return fmt.Errorf("Missing Username or Password")
	}

	return nil
}

func (lr *LoginRequest) CheckLogin() (error, LoginResponse) {

	var password string
	rows, err := dbConnection().Query("select password from users where username=?", lr.Username)
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&password)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	var password_tes = bcrypt.CompareHashAndPassword([]byte(password), []byte(lr.Password))
	if password_tes != nil {
		fmt.Print(err.Error())
	}
	return nil, LoginResponse{}

}
func (rr *RegisterRequest) InsertRegister() (error, RegisterResponse) {
	password := rr.Password
	username := rr.Username
	firstname := rr.FirstName
	lastname := rr.LastName
	rows, err := dbConnection().Query("select username from users where username=?", username)
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&username)
		if err == nil {
			fmt.Print(err.Error())
		}
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	stmt, err := dbConnection().Prepare("INSERT users SET username=?, password=?, first_name=?, last_name=?")
	if err == nil {
		_, err := stmt.Exec(username, hashedPassword, firstname, lastname)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return nil, RegisterResponse{}
}
