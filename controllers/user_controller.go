package controllers

import (
	"fmt"
	"net/http"
	m "uts21/models"
	// "github.com/gorilla/mux"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)


func Login(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM users WHERE userPassword=?"

	errParseForm := r.ParseForm()
	if errParseForm != nil {
		sendErrorResponse(w, 500, "Failed to parse form")
		return
	}

	password := r.Form.Get("password")
	if password == "" {
		sendErrorResponse(w, 400, "Bad Request! You must incldue password!")
		return
	}

	_, err := db.Query(query, password)
	if err != nil {
		sendErrorResponse(w, 500, "Internal Server Error! Login Fail")
		return
	} 

	sendSuccessResponse(w, "Login Successful!")
}


func UpdateUser(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	queryGetUser := "SELECT * FROM users WHERE userId = ?"

	errParseForm := r.ParseForm()
	if errParseForm != nil {
		sendErrorResponse(w, 500, "Failed to parse form")
		return
	}

	id := r.Form.Get("id")

	result, errGetUser := db.Query(queryGetUser, id)
	if errGetUser != nil {
		sendErrorResponse(w, 500, "Internal Server Error! Database query failed!")
		return
	}

	var user m.User
	var users []m.User
	var userType int
	if result.Next() {
		errScan := result.Scan(&user.UserId, &user.UserName, &user.UserEmail, &user.UserPassword, &user.UserCountry, &userType)

		if errScan != nil {
			fmt.Println(errScan)
			sendErrorResponse(w, 500, "Internal Server Error! Fail to Scan!")
			return
		}
	}else{
		sendErrorResponse(w, 404, "account not found!")
		return
	}

	username := r.Form.Get("username")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	country := r.Form.Get("country")

	if username != "" {
		user.UserName = username
	}

	if email != "" {
		user.UserEmail = email
	}

	if password != "" {
		user.UserPassword = password
	}

	if country != "" {
		user.UserCountry = country
	}

	users = append(users, user)

	queryUpdate := "UPDATE users SET userName = ?, userEmail = ?, userPassword = ?, userCountry = ?, userType = ? WHERE userId = ?"

	resultUpdate, errUpdate := db.Exec(queryUpdate, user.UserName, user.UserEmail, user.UserPassword, user.UserCountry, userType, id)

	if errUpdate != nil {
		fmt.Println(errUpdate)
		sendErrorResponse(w, 500, "Database Query Fail")
		return
	}

	rowsAffected,_:= resultUpdate.RowsAffected()
	if rowsAffected == 0 {
		sendErrorResponse(w, 404, "Data not found!")
		return
	}

	sendUserResponse(w, "Update Successful", users)
}