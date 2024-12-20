package omise

import (
    "os"

    "github.com/omise/omise-go"
    "github.com/omise/omise-go/operations"
)

type Client struct {
    client *omise.Client
}

func NewClient() (*Client, error) {
    publicKey := os.Getenv("OMISE_PUBLIC_TEST")
    secretKey := os.Getenv("OMISE_SECRET_TEST")

    client, err := omise.NewClient(publicKey, secretKey)
    if err != nil {
        return nil, err
    }

    return &Client{client: client}, nil
}

func (c *Client) CreateCharge(amount int64, currency string, card string) (*omise.Charge, error) {
    charge, create := &omise.Charge{}, &operations.CreateCharge{
        Amount:   amount,
        Currency: currency,
        Card:     card,
    }
    if err := c.client.Do(charge, create); err != nil {
        return nil, err
    }
    return charge, nil
}