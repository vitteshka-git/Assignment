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

// Task 3: Define Bounded Contexts (Isolated Department Entity)
type Department struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:100;not null;uniqueIndex" json:"name"`
	Description string `gorm:"size:255" json:"description"`
}

func main() {
	_ = godotenv.Load()
	port := os.Getenv("DEPARTMENT_SERVICE_PORT")
	if port == "" {
		port = "8082"
	}

	// Task 7: Design database-per-service (Connect exclusively to department_db)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DEPARTMENT_DB_HOST"), os.Getenv("DEPARTMENT_DB_USER"),
		os.Getenv("DEPARTMENT_DB_PASSWORD"), os.Getenv("DEPARTMENT_DB_NAME"),
		os.Getenv("DEPARTMENT_DB_PORT"))

	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if db != nil {
		db.AutoMigrate(&Department{})
	}

	router := gin.New()

	// Task 14: Create an observability architecture (Structured logging)
	router.Use(gin.Recovery(), func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("[Department-Service] %s %s Status:%d Duration:%s", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), time.Since(start))
	})

	// Task 14: Create an observability architecture (Health & Readiness)
	router.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"service": "department", "status": "UP"}) })
	router.GET("/ready", func(c *gin.Context) { c.JSON(200, gin.H{"ready": true}) })

	// Task 5: Define Department Service (Get All Departments)
	router.GET("/departments", func(c *gin.Context) {
		var depts []Department
		if db != nil {
			db.Find(&depts)
		}
		c.JSON(200, depts)
	})

	// Task 5: Define Department Service (Get Department by ID)
	router.GET("/departments/:id", func(c *gin.Context) {
		var dept Department
		if db != nil {
			if err := db.First(&dept, c.Param("id")).Error; err != nil {
				c.JSON(404, gin.H{"error": "Department not found"})
				return
			}
		}
		c.JSON(200, dept)
	})

	// Task 5: Define Department Service (Create Department)
	router.POST("/departments", func(c *gin.Context) {
		var dept Department
		if err := c.ShouldBindJSON(&dept); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if db != nil {
			db.Create(&dept)
		}
		c.JSON(http.StatusCreated, dept)
	})

	log.Printf("Department Service running on port %s...", port)
	router.Run(":" + port)
}
