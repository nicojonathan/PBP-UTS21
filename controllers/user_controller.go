package controllers

import (
	"net/http"
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