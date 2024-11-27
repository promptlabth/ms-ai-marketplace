package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/gorilla/mux"
	"github.com/promptlabth/ms-ai-marketplace/app/user/service"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{userService: userService}
}

func (h userHandler) NewUser(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("content-type") != "application/json" {
		http.Error(w, "Content-Type not application/json", http.StatusUnsupportedMediaType)
		return
	}

	var request service.NewUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}

	response, err := h.userService.NewUser(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h userHandler) GetUser(c *gin.Context) {
    firebaseID := c.Param("firebase_id") // Gin's way of accessing route parameters
    if firebaseID == "" {
        log.Println("firebaseID cannot be empty")
        c.JSON(http.StatusBadRequest, gin.H{"error": "firebaseID cannot be empty"})
        return
    }

    log.Printf("Fetching user with firebaseID: %s", firebaseID)
    response, err := h.userService.GetUser(firebaseID)
    if err != nil {
        log.Println(err) // Log the error
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, response)
}

