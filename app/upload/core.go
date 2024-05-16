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

func (c *Core) Upload(ctx context.Context, file multipart.File, filename string) error {

	bucketName := "promtlab-image-storage"
	uploadPath := "marketplace/"

	// os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "prompt-lab-383408-512938be4baf.json")
	// cl, err := storage.NewClient(ctx)
	// if err != nil {
	// 	log.Fatalf("Failed to create client: %v", err)
	// }

	// Upload an object with storage.Writer.
	wc := c.storeage.Bucket(bucketName).Object(uploadPath + filename).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}
