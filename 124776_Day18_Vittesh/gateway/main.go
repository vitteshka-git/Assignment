package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Task 11: Design service discovery (Function to dynamically resolve service URLs)
func discoverService(envKey string, fallback string) string {
	if val := os.Getenv(envKey); val != "" {
		return val
	}
	return fallback
}

// Task 10: Design API Gateway (HTTP Proxying and Header Forwarding)
func forwardRequest(targetBaseURL string, c *gin.Context, targetPath string) {
	client := &http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequestWithContext(c.Request.Context(), c.Request.Method, targetBaseURL+targetPath, c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to construct proxy request"})
		return
	}

	for key, values := range c.Request.Header {
		for _, val := range values {
			req.Header.Add(key, val)
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Target microservice unavailable"})
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	for key, values := range resp.Header {
		for _, val := range values {
			c.Header(key, val)
		}
	}
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func main() {
	_ = godotenv.Load()
	port := os.Getenv("GATEWAY_PORT")
	if port == "" {
		port = "8080"
	}

	// Task 1: Analyze the Employee Monolith
	// Task 2: Identify microservice boundaries
	employeeServiceURL := discoverService("EMPLOYEE_SERVICE_URL", "http://localhost:8081")
	departmentServiceURL := discoverService("DEPARTMENT_SERVICE_URL", "http://localhost:8082")
	authServiceURL := discoverService("AUTH_SERVICE_URL", "http://localhost:8083")

	router := gin.New()

	// Task 14: Create an observability architecture (Central request logging)
	router.Use(gin.Recovery(), func(c *gin.Context) {
		start := time.Now()
		c.Next()
		log.Printf("[Gateway] %s %s Status:%d Duration:%s", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), time.Since(start))
	})

	// Task 14: Create an observability architecture (Health & Readiness probes)
	router.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"service": "gateway", "status": "UP"}) })
	router.GET("/ready", func(c *gin.Context) { c.JSON(200, gin.H{"ready": true}) })

	// Task 4: Define Employee Service Routes via API Gateway
	router.GET("/api/employees", func(c *gin.Context) { forwardRequest(employeeServiceURL, c, "/employees") })
	router.POST("/api/employees", func(c *gin.Context) { forwardRequest(employeeServiceURL, c, "/employees") })
	router.GET("/api/employees/:id", func(c *gin.Context) { forwardRequest(employeeServiceURL, c, "/employees/"+c.Param("id")) })

	// Task 5: Define Department Service Routes via API Gateway
	router.GET("/api/departments", func(c *gin.Context) { forwardRequest(departmentServiceURL, c, "/departments") })
	router.POST("/api/departments", func(c *gin.Context) { forwardRequest(departmentServiceURL, c, "/departments") })

	// Task 6: Define Authentication Service Routes via API Gateway
	router.POST("/api/auth/login", func(c *gin.Context) { forwardRequest(authServiceURL, c, "/auth/login") })

	log.Printf("API Gateway running on port %s...", port)
	router.Run(":" + port)
}
