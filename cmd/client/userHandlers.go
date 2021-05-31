package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"userService/cmd/client/mappers"
	"userService/pkg/models/requestModels"
)

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
		app.serverError(w, err, http.StatusBadRequest)
		return
	}

	dbUser := mappers.DbUserMap(userRequest, uuid.New())

	userId, err := app.Db.Insert(dbUser)
	w.Write([]byte("User with id = " + userId.String() + " was created"))
}

func (app *Application) get(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.notAllowed(w)
		return
	}

	requestId, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil {
		app.serverError(w, err, http.StatusBadRequest)
		return
	}

	user, err := app.Db.Get(requestId)
	if err != nil {
		app.serverError(w, err, http.StatusBadRequest)
		return
	}

	jsonUser, err := json.Marshal(user)

	w.Write(jsonUser)
}

func (app *Application) delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		app.notAllowed(w)
		return
	}

	requestId, err := uuid.Parse(r.URL.Query().Get("id"))
	if err != nil {
		app.serverError(w, err, http.StatusBadRequest)
		return
	}

	result, err := app.Db.Delete(requestId)
	if err != nil || result == false {
		app.serverError(w, err, http.StatusBadRequest) // TODO add string with description
		return
	}

	w.Write([]byte("User with userId = " + requestId.String() + " was deleted"))
}
