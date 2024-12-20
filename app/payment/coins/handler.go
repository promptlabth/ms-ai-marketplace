package coins

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type usecase interface {
	CreateCoins(ctx context.Context, coins CoinsEntity) (*uint, error)
	GetSumOfCoinsByFirebaseIDAndAgentID(ctx context.Context, firebaseID string, agentID int) (int, error)
	SetCoinsToZeroByFirebaseIDAndAgentID(ctx context.Context, firebaseID string, agentID int) error
}

type Handler struct {
	usecase usecase
}

func NewHandler(u usecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) CreateCoins(c *gin.Context) {
	var coinReq CoinReq
	if err := c.ShouldBindJSON(&coinReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	firebaseID, exists := c.Get("firebase_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Firebase ID not found in token"})
		return
	}

	coins := CoinsEntity{
		FirebaseID: firebaseID.(string),
		AgentID:    coinReq.AgentID,
		Coins:      0, // Set default coins to zero
	}

	id, err := h.usecase.CreateCoins(c.Request.Context(), coins)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) AddCoins(c *gin.Context) {
	var coins CoinsEntity
	if err := c.ShouldBindJSON(&coins); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	firebaseID, exists := c.Get("firebase_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Firebase ID not found in token"})
		return
	}

	coins.FirebaseID = firebaseID.(string)

	existingCoins, err := h.usecase.GetSumOfCoinsByFirebaseIDAndAgentID(c.Request.Context(), coins.FirebaseID, coins.AgentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	coins.Coins += existingCoins

	_, err = h.usecase.CreateCoins(c.Request.Context(), coins)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Coins added successfully"})
}

func (h *Handler) GetCoins(c *gin.Context) {
	firebaseID, exists := c.Get("firebase_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Firebase ID not found in token"})
		return
	}

	agentIDStr := c.Param("agentID")
	agentID, err := strconv.Atoi(agentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid agent ID"})
		return
	}

	totalCoins, err := h.usecase.GetSumOfCoinsByFirebaseIDAndAgentID(c.Request.Context(), firebaseID.(string), agentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"coins": totalCoins})
}

func (h *Handler) SetCoinsToZero(c *gin.Context) {
	firebaseID, exists := c.Get("firebase_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Firebase ID not found in token"})
		return
	}

	agentIDStr := c.Param("agentID")
	agentID, err := strconv.Atoi(agentIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid agent ID"})
		return
	}

	totalCoins, err := h.usecase.GetSumOfCoinsByFirebaseIDAndAgentID(c.Request.Context(), firebaseID.(string), agentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	fmt.Println("totalCoins", totalCoins)

	err = h.usecase.SetCoinsToZeroByFirebaseIDAndAgentID(c.Request.Context(), firebaseID.(string), agentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Coins set to zero", "totalCoins": totalCoins})
}
