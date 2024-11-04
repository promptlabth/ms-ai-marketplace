package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/promptlabth/ms-ai-marketplace/app/user/service"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{userService: userService}
}

func (h userHandler) NewUser(w http.ResponseWriter, r *http.Request){
	if r.Header.Get("content-type") != "application/json" {
		log.Fatal("content-type not JSON")
		return
	}

	request := service.NewUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal("Error decode json")
		return
	}

	response , err := h.userService.NewUser(request)
		if err != nil {
			log.Fatal(err)
			return
	}
	

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type" , "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    firebaseID := vars["id"]

    response, err := h.userService.GetUser(firebaseID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("content-type", "application/json")
    json.NewEncoder(w).Encode(response)
}
