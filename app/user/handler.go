package user

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userUsecaser interface {
	NewUser(c context.Context, user User) error
	GetUser(c context.Context, firebase_id string) (*UserEntity, error)
}

type Handler struct {
	userUsecase userUsecaser
}

func NewHandler(u userUsecaser) *Handler {
	return &Handler{userUsecase: u}
}

func (h *Handler) NewUser(c *gin.Context) {
	var req NewUserRequest

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
		return
	}

	user := User{
		FirebaseID: req.FirebaseID,
		Name:       req.Name,
		Email:      req.Email,
		Platform:   req.Platform,
		StripeID:   req.StripeID,
		PlanID:     req.PlanID,
		Password:   req.Password,
	}

	if err := h.userUsecase.NewUser(context.Background(), user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (h *Handler) GetUser(c *gin.Context) {
	firebaseID := c.Param("id")
	userByID, err := h.userUsecase.GetUser(c.Request.Context(), firebaseID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userByID})
}
