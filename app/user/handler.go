package user

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type usecase interface {
	NewUser(c context.Context, user User) error
	GetUser(c context.Context, firebase_id string) (*UserEntity, error) 
}

type Handler struct {
	usecase usecase
}

func NewHandler(u usecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) NewUser(c *gin.Context) {
	var req NewUserRequest

	if err :=c .Bind(&req); err != nil {
		c.JSON(404, map[string]string{
			"error": err.Error(),
		})
		return
	}

	user := User{
		FirebaseID:     req.FirebaseID,
		Name:           req.Name,
		Email:          req.Email,
		Platform:       req.Platform,
		StripeID:       req.StripeID,
		PlanID:         req.PlanID,
		Password:       req.Password,

	}

	if err := h.usecase.NewUser(context.Background(), user); err != nil {
		c.AbortWithStatus(500)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

func (h *Handler) GetUser(c *gin.Context) {
	firebaseID := c.Param("id")
	log.Printf("firebaseID : %+v\n", firebaseID)
    userByID, err := h.usecase.GetUser(c.Request.Context(), firebaseID)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to get user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": userByID})
}