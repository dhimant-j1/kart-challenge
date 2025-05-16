package main

import (
	_ "backend-challenge/docs" // Import docs for swagger
	"backend-challenge/internal/api/handlers"
	"backend-challenge/internal/middleware"
	"backend-challenge/internal/service"
	"log"
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Order Food Online API
// @version         1.0
// @description     This is a food ordering API based on the OpenAPI 3.1 specification.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name api_key
// @description API key authentication for order operations

func main() {
	// Initialize services
	productService := service.NewProductService()
	orderService := service.NewOrderService(productService)

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productService)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Initialize router
	router := gin.Default()

	// Apply global middleware
	router.Use(middleware.CORS())
	router.Use(middleware.APIKeyAuth())

	// API routes
	apiGroup := router.Group("/api")
	{
		// Product endpoints
		apiGroup.GET("/product", productHandler.ListProducts)
		apiGroup.GET("/product/:productId", productHandler.GetProduct)

		// Order endpoints
		apiGroup.POST("/order", orderHandler.PlaceOrder)
	}

	// Swagger documentation endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Serve OpenAPI spec for Scalar
	router.GET("/openapi.json", func(c *gin.Context) {
		c.File("./docs/swagger.json")
	})

	// Scalar API Reference endpoint
	router.GET("/reference", func(c *gin.Context) {
		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "http://localhost:8080/openapi.json", // Use the URL that serves your OpenAPI spec
			CustomOptions: scalar.CustomOptions{
				PageTitle: "Food Ordering API Documentation",
			},
			DarkMode: true,
		})

		if err != nil {
			c.String(http.StatusInternalServerError, "Error generating API reference: %v", err)
			return
		}

		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, htmlContent)
	})

	// Start server
	log.Println("Starting server on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
