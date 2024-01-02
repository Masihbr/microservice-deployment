package main

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"net/http"
)

var client *redis.Client

func init() {
	// Initialize Redis client
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", 
		Password: "",
		DB:       0,           
	})
}

func main() {
	r := gin.Default()

	r.POST("/api/", postNote)
	r.GET("/api/", getNote)
	r.PUT("/api/", putNote)
	r.DELETE("/api/", deleteNote)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

type PostRequest struct {
	Note string `json:"note"`
}

type GetRequest struct {
	ID   string `json:"id"`
	Hash string `json:"hash"`
}

type PutRequest struct {
	ID   string `json:"id"`
	Hash string `json:"hash"`
	Note string `json:"note"`
}

type DeleteRequest struct {
	ID   string `json:"id"`
	Hash string `json:"hash"`
}

type PostResponse struct {
	ID   string `json:"id"`
	Hash string `json:"hash"`
}

func postNote(c *gin.Context) {
	var req PostRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate UUID
	id := uuid.New().String()

	// Calculate hash
	hash := calculateHash(req.Note)

	// Save to Redis
	err := SetValueRedis(id, req.Note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save note"})
		return
	}

	c.JSON(http.StatusOK, PostResponse{ID: id, Hash: hash})
}

func getNote(c *gin.Context) {
	var req GetRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve from Redis
	data, err := GetValueRedis(req.ID)
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get note"})
		return
	}

	// Check if hash matches
	if calculateHash(data) != req.Hash {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash mismatch"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"note": data})
}

func putNote(c *gin.Context) {
	var req PutRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if ID exists
	exists, err := CheckKeyRedis(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check ID existence"})
		return
	}
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
		return
	}

	// Retrieve from Redis
	data, err := GetValueRedis(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get note"})
		return
	}

	// Check if hash matches
	if calculateHash(data) != req.Hash {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash mismatch"})
		return
	}

	// Update note in Redis
	err = SetValueRedis(req.ID, req.Note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note updated successfully", "hash": calculateHash(req.Note)})
}

func deleteNote(c *gin.Context) {
	var req DeleteRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve from Redis
	data, err := GetValueRedis(req.ID)
	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get note"})
		return
	}

	// Check if hash matches
	if calculateHash(data) != req.Hash {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Hash mismatch"})
		return
	}

	// Delete note from Redis
	err = DeleteKeyRedis(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted successfully"})
}

func calculateHash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
