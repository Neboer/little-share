package lib

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// FileInfo holds information about a stored file
type FileInfo struct {
	FileName         string `json:"name"`
	FileSizeBytes    int64  `json:"size"`
	FileCreationTime int64  `json:"creation_time"`
}

// StoreToLocal handles file uploads and stores them on disk
func StoreToLocal(c *gin.Context, maxStorageSize int64, fileStorePath string) error {
	defer func() {
		if r := recover(); r != nil {
			slog.Error("Panic during file upload", "error", r)
			if r == io.ErrUnexpectedEOF {
				c.AbortWithStatusJSON(http.StatusInternalServerError,
					gin.H{"error": "Upload interrupted by client"})
			}
		}
	}()

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": "No 'file' field in request"})
		slog.Error("Error reading form file", "error", err)
		return err
	}
	defer file.Close()

	filename := header.Filename
	currentSize := GetCurrentTotalFileSize(fileStorePath)
	availableSpace := maxStorageSize - currentSize

	// Check if file is too large
	if header.Size >= availableSpace {
		c.AbortWithStatusJSON(http.StatusRequestEntityTooLarge,
			gin.H{"error": fmt.Sprintf("File too large! Available space: %d bytes", availableSpace)})
		return fmt.Errorf("file too large: %d bytes, available: %d bytes", header.Size, availableSpace)
	}

	// Ensure storage path exists
	if err := os.MkdirAll(fileStorePath, 0755); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": "Failed to create storage directory"})
		slog.Error("Failed to create file store directory", "path", fileStorePath, "error", err)
		return err
	}

	// Save the file
	savePath := filepath.Join(fileStorePath, filename)
	dst, err := os.Create(savePath)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": "Failed to save file"})
		slog.Error("Failed to create output file", "path", savePath, "error", err)
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": "Failed to write file data"})
		slog.Error("Failed to write file", "path", savePath, "error", err)
		return err
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("File %s uploaded successfully", filename)})
	slog.Info("File uploaded successfully", "filename", filename, "size", header.Size)
	return nil
}

// DeleteOneFile deletes a file from storage
func DeleteOneFile(fileName string, fileStorePath string) error {
	filePath := filepath.Join(fileStorePath, fileName)

	// Check if file exists
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file not found: %s", fileName)
		}
		return fmt.Errorf("failed to stat file: %w", err)
	}

	// Delete the file
	if err := os.Remove(filePath); err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	slog.Info("File deleted successfully", "filename", fileName)
	return nil
}

// GetFileList returns a list of all files in the storage directory
func GetFileList(fileStorePath string) ([]FileInfo, error) {
	entries, err := os.ReadDir(fileStorePath)
	if err != nil {
		if os.IsNotExist(err) {
			return []FileInfo{}, nil // Return empty list if directory doesn't exist
		}
		return nil, fmt.Errorf("failed to read file directory: %w", err)
	}

	var files []FileInfo
	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				slog.Warn("Failed to get file info", "name", entry.Name(), "error", err)
				continue
			}

			files = append(files, FileInfo{
				FileName:         entry.Name(),
				FileCreationTime: info.ModTime().Unix(),
				FileSizeBytes:    info.Size(),
			})
		}
	}

	return files, nil
}

// GetCurrentTotalFileSize returns the total size of all files in storage
func GetCurrentTotalFileSize(fileStorePath string) int64 {
	entries, err := os.ReadDir(fileStorePath)
	if err != nil {
		if os.IsNotExist(err) {
			return 0
		}
		slog.Warn("Failed to read file directory", "path", fileStorePath, "error", err)
		return 0
	}

	var totalSize int64
	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				slog.Warn("Failed to get file info", "name", entry.Name(), "error", err)
				continue
			}
			totalSize += info.Size()
		}
	}

	return totalSize
}

// GetAvailableSpace returns the available storage space in bytes
func GetAvailableSpace(maxStorageSize int64, fileStorePath string) int64 {
	currentSize := GetCurrentTotalFileSize(fileStorePath)
	available := maxStorageSize - currentSize
	if available < 0 {
		return 0
	}
	return available
}
