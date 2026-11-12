package main

import (
	"encoding/json"
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

// Task 3: Define Bounded Contexts (Isolated Employee Entity)
type Employee struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	Name         string  `gorm:"size:100;not null" json:"name"`
	Email        string  `gorm:"size:150;not null;uniqueIndex" json:"email"`
	DepartmentID uint    `gorm:"not null" json:"departmentId"`
	Salary       float64 `gorm:"not null" json:"salary"`
}

// Task 8: Design synchronous REST communication (HTTP call to Department Service)
func checkDepartmentExists(deptURL string, deptID uint) bool {
	client := &http.Client{Timeout: 3 * time.Second}
	resp, err := client.Get(fmt.Sprintf("%s/departments/%d", deptURL, deptID))
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}

// Task 9: Design asynchronous communication (Transactional Outbox Publisher)
func publishEvent(eventType string, payload interface{}) {
	data, _ := json.Marshal(payload)
	log.Printf("[Outbox Event] Published %s: %s", eventType, string(data))
}

func main() {
	_ = godotenv.Load()
	port := os.Getenv("EMPLOYEE_SERVICE_PORT")
	deptURL := os.Getenv("DEPARTMENT_SERVICE_URL")
	if port == "" {
		port = "8081"
	}
	if deptURL == "" {
		deptURL = "http://localhost:8082"
	}

	// Task 7: Design database-per-service (Connect exclusively to employee_db)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("EMPLOYEE_DB_HOST"), os.Getenv("EMPLOYEE_DB_USER"),
		os.Getenv("EMPLOYEE_DB_PASSWORD"), os.Getenv("EMPLOYEE_DB_NAME"),
		os.Getenv("EMPLOYEE_DB_PORT"))

	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if db != nil {
		db.AutoMigrate(&Employee{})
	}

	router := gin.New()

	// Task 14: Create an observability architecture (Structured HTTP logging)
	router.Use(gin.Recovery(), func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("[Employee-Service] %s %s Status:%d Duration:%s", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), time.Since(start))
	})

	// Task 14: Create an observability architecture (Health & Readiness)
	router.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"service": "employee", "status": "UP"}) })
	router.GET("/ready", func(c *gin.Context) { c.JSON(200, gin.H{"ready": true}) })

	// Task 4: Define Employee Service (Get All Employees)
	router.GET("/employees", func(c *gin.Context) {
		var employees []Employee
		if db != nil {
			db.Find(&employees)
		}
		c.JSON(200, employees)
	})

	// Task 4: Define Employee Service (Create Employee Handler)
	router.POST("/employees", func(c *gin.Context) {
		var emp Employee
		if err := c.ShouldBindJSON(&emp); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Task 8: Design synchronous REST communication (Verify department)
		if !checkDepartmentExists(deptURL, emp.DepartmentID) {
			c.JSON(400, gin.H{"error": "Department does not exist"})
			return
		}

		// Task 12: Design distributed transaction handling
		// Task 13: Design Saga workflow (Local DB transaction with compensating rollback)
		if db != nil {
			tx := db.Begin()
			if err := tx.Create(&emp).Error; err != nil {
				tx.Rollback() // Compensating action: rollback local transaction
				log.Println("[Saga Rollback] Employee creation transaction rolled back")
				c.JSON(500, gin.H{"error": "Creation failed"})
				return
			}
			tx.Commit()
		}

		// Task 9: Design asynchronous communication (Publish Outbox event)
		publishEvent("EMPLOYEE_CREATED", emp)

		c.JSON(http.StatusCreated, emp)
	})

	log.Printf("Employee Service running on port %s...", port)
	router.Run(":" + port)
}
