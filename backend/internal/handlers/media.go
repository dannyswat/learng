package handlers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/learng/backend/internal/utils"
)

type MediaHandler struct {
	uploadDir    string
	maxImageSize int64
	maxAudioSize int64
}

func NewMediaHandler(uploadDir string) *MediaHandler {
	return &MediaHandler{
		uploadDir:    uploadDir,
		maxImageSize: 5 * 1024 * 1024, // 5MB
		maxAudioSize: 2 * 1024 * 1024, // 2MB
	}
}

// UploadImage handles image file uploads
func (h *MediaHandler) UploadImage(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("No file uploaded"))
	}

	// Validate file size
	if file.Size > h.maxImageSize {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(
			fmt.Sprintf("File too large (max %dMB)", h.maxImageSize/(1024*1024)),
		))
	}

	// Validate file extension
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !isValidImageExtension(ext) {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(
			"Invalid file extension. Supported: .jpg, .jpeg, .png, .webp",
		))
	}

	// Validate MIME type
	mimeType := file.Header.Get("Content-Type")
	if !isValidImageType(mimeType) {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(
			"Invalid file type. Supported formats: JPEG, PNG, WebP",
		))
	}

	// Generate unique filename
	filename := uuid.New().String() + ext
	imageDir := filepath.Join(h.uploadDir, "images")

	// Ensure directory exists
	if err := os.MkdirAll(imageDir, 0755); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create upload directory"))
	}

	savePath := filepath.Join(imageDir, filename)

	// Save file
	if err := saveUploadedFile(file, savePath); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to save file"))
	}

	// Return response
	url := "/uploads/images/" + filename
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"url":      url,
		"filename": filename,
		"size":     file.Size,
		"mimeType": mimeType,
	})
}

// UploadAudio handles audio file uploads
func (h *MediaHandler) UploadAudio(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse("No file uploaded"))
	}

	// Validate file size
	if file.Size > h.maxAudioSize {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(
			fmt.Sprintf("File too large (max %dMB)", h.maxAudioSize/(1024*1024)),
		))
	}

	// Validate file extension (primary validation)
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !isValidAudioExtension(ext) {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(
			"Invalid file extension. Supported: .mp3, .wav, .webm",
		))
	}

	// Validate MIME type (allow application/octet-stream if extension is valid)
	mimeType := file.Header.Get("Content-Type")
	if !isValidAudioType(mimeType) {
		return c.JSON(http.StatusBadRequest, utils.ErrorResponse(
			"Invalid file type. Supported formats: MP3, WAV, WebM",
		))
	}

	// Generate unique filename
	filename := uuid.New().String() + ext
	audioDir := filepath.Join(h.uploadDir, "audio")

	// Ensure directory exists
	if err := os.MkdirAll(audioDir, 0755); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to create upload directory"))
	}

	savePath := filepath.Join(audioDir, filename)

	// Save file
	if err := saveUploadedFile(file, savePath); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.ErrorResponse("Failed to save file"))
	}

	// Return response
	url := "/uploads/audio/" + filename
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"url":      url,
		"filename": filename,
		"size":     file.Size,
		"mimeType": mimeType,
	})
}

// Helper functions

func isValidImageType(mimeType string) bool {
	validTypes := []string{
		"image/jpeg",
		"image/png",
		"image/webp",
		"application/octet-stream",
	}
	for _, t := range validTypes {
		if t == mimeType {
			return true
		}
	}
	return false
}

func isValidImageExtension(ext string) bool {
	validExts := []string{".jpg", ".jpeg", ".png", ".webp"}
	for _, e := range validExts {
		if e == ext {
			return true
		}
	}
	return false
}

func isValidAudioType(mimeType string) bool {
	validTypes := []string{
		"audio/mpeg",
		"audio/mp3",
		"audio/wav",
		"audio/wave",
		"audio/x-wav",
		"audio/webm",
		"application/octet-stream",
	}
	for _, t := range validTypes {
		if t == mimeType {
			return true
		}
	}
	return false
}

func isValidAudioExtension(ext string) bool {
	validExts := []string{".mp3", ".wav", ".webm"}
	for _, e := range validExts {
		if e == ext {
			return true
		}
	}
	return false
}

func saveUploadedFile(file *multipart.FileHeader, savePath string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
