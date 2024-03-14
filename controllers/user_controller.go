package controllers

import (
	"encoding/json"
	"fmt"
	model "latihanbackend/models"
	"log"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT * FROM users"

	//Read from Header
	// name := r.Header.Get("name")
	// if name != "" {
	// 	query += " WHERE name ='" + name + "'"
	// }

	// Read from Query Param
	name := r.URL.Query()["name"]
	age := r.URL.Query()["age"]
	if name != nil {
		fmt.Println(name[0])
		//fmt.Println(name[1])
		query += " WHERE name ='" + name[0] + "'"
	}

	if age != nil {
		if name[0] != "" {
			query += " AND"
		} else {
			query += " WHERE"
		}
		query += " age ='" + age[0] + "'"
	}

	rows, err := db.Query(query)

	if err != nil {
		log.Println(err)
		//send error response
		return
	}

	var user model.User
	var users []model.User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.UserType); err != nil {
			log.Print(err)
			//send error response
			return
		} else {
			users = append(users, user)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	//var response UsersResponse
	//if len(users) < 5 {
	var response model.UsersResponse
	response.Status = "200"
	response.Message = "Success"
	response.Data = users
	json.NewEncoder(w).Encode(response)
	// } else {
	// var response ErrorResponse
	// 	response.Status = 400
	// response.Message = "Error Array Size not correct"
}
