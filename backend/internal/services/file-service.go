package services

import (
	"backend/config"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"time"
)

// UploadFileToGCS uploads a file to GCS and returns the public URL
func UploadFileToGCS(file multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	ctx := context.Background()

	// Prepare file for upload
	fileName := fmt.Sprintf("%d-%s", time.Now().Unix(), fileHeader.Filename)

	buffer := bytes.NewBuffer(nil)
	_, err := io.Copy(buffer, file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	// Upload to GCS
	wc := config.GCSClient.Bucket(config.BucketName).Object(fileName).NewWriter(ctx)
	wc.ContentType = fileHeader.Header.Get("Content-Type")

	if _, err := io.Copy(wc, buffer); err != nil {
		return "", fmt.Errorf("error writing to GCS: %v", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("error closing GCS writer: %v", err)
	}

	fileURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s", config.BucketName, fileName)
	log.Println("✅ File uploaded to GCS:", fileURL)

	return fileURL, nil
}
