package routes

import (
	"sd3971-go/internal/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers all application routes
func RegisterRoutes(router *gin.Engine, productHandler *handlers.ProductHandler) {
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Product routes
	products := router.Group("/products")
	{
		// GET /products - retrieve all products
		products.GET("", productHandler.GetProducts)

		// POST /products - create a new product
		products.POST("", productHandler.CreateProduct)

		// GET /products/:id - retrieve a single product
		products.GET("/:id", productHandler.GetProductByID)

		// PUT /products/:id - update a product
		products.PUT("/:id", productHandler.UpdateProduct)

		// DELETE /products/:id - delete a product
		products.DELETE("/:id", productHandler.DeleteProduct)
	}
}
