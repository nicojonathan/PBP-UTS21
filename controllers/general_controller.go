package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func getUserType(w http.ResponseWriter, db *sql.DB, password string) int {
	queryUser := "SELECT userType FROM users WHERE userPassword = ?"

	if password == "" {
		sendErrorResponse(w, 400, "Bad Request! Password required!")
		return 0
	}

	result, errQueryUser := db.Query(queryUser, password)
	if errQueryUser != nil {
		fmt.Println(errQueryUser)
		sendErrorResponse(w, 500, "Internal Server Error!")

		return 0
	}

	var userType int
	if result.Next() {
		result.Scan(&userType)
		return userType
	} else {
		sendErrorResponse(w, 404, "account not found!")
		return 0
	}

	//fmt.Println("type: ", userType)
}