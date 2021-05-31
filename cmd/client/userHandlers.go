package main

import (
	"encoding/json"
	"net/http"
	"userService/cmd/client/mappers"
	"userService/pkg/models/requestModels"
)

var id = 1 // TODO use UUID

func (app *Application) create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		app.notAllowed(w)
		return
	}

	if r.Body == nil {
		http.Error(w, "Please send a request body", http.StatusBadRequest)
		return
	}

	var userRequest requestModels.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest) // TODO add validation

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dbUser := mappers.DbUserMap(userRequest, id)
	id++

	print(dbUser) // TODO add to db
}

func (app *Application) get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.notAllowed(w)
		return
	}

	//requestId := r.URL.Query().Get("id")

	w.Write([]byte("Hello get user")) // TODO get user
}

func (app *Application) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		app.notAllowed(w)
		return
	}

	//requestId := r.URL.Query().Get("id")

	w.Write([]byte("Hello delete user")) // TODO delete user
}
