package upload

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"

	"cloud.google.com/go/storage"
)

type Core struct {
	storeage *storage.Client
}

func NewCore(
	storeage *storage.Client,

) *Core {
	return &Core{
		storeage: storeage,
	}
}

func (c *Core) Upload(ctx context.Context, file multipart.File, filename string) (string, error) {
	bucketName := "promptlab-image-storage"
	uploadPath := "marketplace/"

	// Upload an object with storage.Writer.
	wc := c.storeage.Bucket(bucketName).Object(uploadPath + filename).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return "e", fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %v", err)
	}

	// Construct the URL of the uploaded file
	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s%s", bucketName, uploadPath, filename)

	return url, nil
}
