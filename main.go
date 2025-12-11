package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Neboer/little-share/config"
	"github.com/Neboer/little-share/lib"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration from command line or defaults
	cfg := config.LoadWithFlag()

	slog.Info("Starting little-share server",
		"host", cfg.Host,
		"port", cfg.Port,
		"frontend_path", cfg.FrontendPath,
		"file_store_path", cfg.FileStorePath,
		"max_storage_size", cfg.MaxStorageSize)

	// Initialize Gin router
	r := gin.Default()

	// Serve frontend
	r.Use(static.Serve("/", static.LocalFile(cfg.FrontendPath, true)))

	// API: Get available storage space
	r.GET("/maxSpace", func(c *gin.Context) {
		available := lib.GetAvailableSpace(cfg.MaxStorageSize, cfg.FileStorePath)
		c.String(http.StatusOK, strconv.FormatInt(available, 10))
	})

	// API: List all stored files
	r.GET("/files", func(c *gin.Context) {
		fileList, err := lib.GetFileList(cfg.FileStorePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list files"})
			slog.Error("Failed to list files", "error", err)
			return
		}

		if len(fileList) == 0 {
			c.JSON(http.StatusOK, []lib.FileInfo{})
		} else {
			c.JSON(http.StatusOK, fileList)
		}
	})

	// API: Delete a specific file
	r.DELETE("/files/:file", func(c *gin.Context) {
		fileName := c.Param("file")
		err := lib.DeleteOneFile(fileName, cfg.FileStorePath)
		if err != nil {
			c.String(http.StatusNotFound, fmt.Sprintf("Error deleting file: %v", err))
			slog.Warn("Failed to delete file", "filename", fileName, "error", err)
		} else {
			c.String(http.StatusOK, fmt.Sprintf("File %s deleted successfully", fileName))
		}
	})

	// Serve file downloads
	r.Use(static.Serve("/files", static.LocalFile(cfg.FileStorePath, false)))

	// API: Upload file
	r.POST("/upload", func(c *gin.Context) {
		_ = lib.StoreToLocal(c, cfg.MaxStorageSize, cfg.FileStorePath)
	})

	// Start server
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	slog.Info("Server listening on", "address", addr)
	if err := r.Run(addr); err != nil {
		slog.Error("Failed to start server", "error", err)
	}
}
