package controllers

import (
	"encoding/json"
	"fmt"
	"gorilla-test/data"
	datatypes "gorilla-test/dataTypes"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) { // L is capital bcz I declared as global function
	var user datatypes.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error while decoidn json requred body", http.StatusBadRequest)
		return
	}

	_, exist := data.Users[user.Username]

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user not found"))
		return
	}

	if data.Users[user.Username] == user.Password {
		w.WriteHeader(http.StatusFound)
		w.Write([]byte("valid user"))
	} else {
		w.Write([]byte("Invalid password"))
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) { // S is capital bcz I declared as global function
	var newUser datatypes.UserCredentials
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "error while decoidn json requred body", http.StatusBadRequest)
		return
	}

	_, exist := data.Users[newUser.Username]

	if exist {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("user already exist"))
		return
	}

	data.Users[newUser.Username] = newUser.Password
	w.Write([]byte("new user has been created"))
}
