package omise

import (
    // "os"
    "testing"

    "github.com/joho/godotenv"
    "github.com/stretchr/testify/assert"
)

func TestCreateCharge(t *testing.T) {
    // Load environment variables from .env file
    err := godotenv.Load("../../../cmd/.env")
    if err != nil {
        t.Fatalf("Error loading .env file: %v", err)
    }

    client, err := NewClient()
    if err != nil {
        t.Fatalf("Failed to create Omise client: %v", err)
    }

    amount := int64(100000) // Amount in smallest currency unit (e.g., cents)
    currency := "THB"
    card := "tokn_test_5f2bq7i4h8x9m5l0k3a" // Replace with a valid test card token

    charge, err := client.CreateCharge(amount, currency, card)
    if err != nil {
        t.Fatalf("Failed to create charge: %v", err)
    }

    assert.NotNil(t, charge)
    assert.Equal(t, amount, charge.Amount)
    assert.Equal(t, currency, charge.Currency)
    assert.Equal(t, "successful", charge.Status)
}