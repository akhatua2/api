package controller

import (
	"encoding/json"
	"github.com/HackIllinois/api-commons/errors"
	"github.com/HackIllinois/api-user/models"
	"github.com/HackIllinois/api-user/service"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

func SetupController(route *mux.Route) {
	router := route.Subrouter()

	router.Handle("/{id}/", alice.New().ThenFunc(GetUserInfo)).Methods("GET")
	router.Handle("/", alice.New().ThenFunc(SetUserInfo)).Methods("POST")
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user_info, err := service.GetUserInfo(id)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(user_info)
}

func SetUserInfo(w http.ResponseWriter, r *http.Request) {
	var user_info models.UserInfo
	json.NewDecoder(r.Body).Decode(&user_info)

	if user_info.ID == "" {
		panic(errors.UnprocessableError("Must provide id parameter"))
	}

	err := service.SetUserInfo(user_info.ID, user_info)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	updated_info, err := service.GetUserInfo(user_info.ID)

	if err != nil {
		panic(errors.UnprocessableError(err.Error()))
	}

	json.NewEncoder(w).Encode(updated_info)
}
