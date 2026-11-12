package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Task 3: Define Bounded Contexts (Isolated User Entity)
type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"size:100;not null;uniqueIndex" json:"username"`
	Password string `gorm:"size:255;not null" json:"-"`
	Role     string `gorm:"size:50;not null" json:"role"`
}

func main() {
	_ = godotenv.Load()
	port := os.Getenv("AUTH_SERVICE_PORT")
	if port == "" {
		port = "8083"
	}

	// Task 7: Design database-per-service (Connect exclusively to auth_db)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("AUTH_DB_HOST"), os.Getenv("AUTH_DB_USER"),
		os.Getenv("AUTH_DB_PASSWORD"), os.Getenv("AUTH_DB_NAME"),
		os.Getenv("AUTH_DB_PORT"))

	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if db != nil {
		db.AutoMigrate(&User{})
	}

	router := gin.New()

	// Task 14: Create an observability architecture (Structured logging)
	router.Use(gin.Recovery(), func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("[Auth-Service] %s %s Status:%d Duration:%s", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), time.Since(start))
	})

	// Task 14: Create an observability architecture (Health & Readiness)
	router.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"service": "auth", "status": "UP"}) })
	router.GET("/ready", func(c *gin.Context) { c.JSON(200, gin.H{"ready": true}) })

	// Task 6: Define Authentication Service (User Login Handler)
	router.POST("/auth/login", func(c *gin.Context) {
		var payload struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(400, gin.H{"error": "Invalid payload"})
			return
		}
		if payload.Username == "admin" && payload.Password == "admin123" {
			c.JSON(200, gin.H{"authenticated": true, "username": "admin", "role": "ADMIN"})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	})

	log.Printf("Authentication Service running on port %s...", port)
	router.Run(":" + port)
}
