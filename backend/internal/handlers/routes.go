package handlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/weaver/api/internal/middleware"
)

func SetupRoutes(router *gin.Engine, db *sql.DB) {
	// Apply middleware
	router.Use(middleware.CORSMiddleware())

	// Health check endpoint
	authHandler := NewAuthHandler(db)
	router.GET("/health", authHandler.HealthCheck)

	// Authentication routes
	authRoutes := router.Group("/api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.GET("/profile", authHandler.GetProfile)
		authRoutes.POST("/logout", authHandler.Logout)
	}

	// Supplier routes
	supplierHandler := NewSupplierHandler(db)
	supplierRoutes := router.Group("/api/suppliers")
	{
		supplierRoutes.GET("", supplierHandler.GetSuppliers)
		supplierRoutes.GET("/:id", supplierHandler.GetSupplier)
		supplierRoutes.POST("", supplierHandler.CreateSupplier)
		supplierRoutes.PUT("/:id", supplierHandler.UpdateSupplier)
		supplierRoutes.DELETE("/:id", supplierHandler.DeleteSupplier)
	}

	// Weaver routes
	weaverHandler := NewWeaverHandler(db)
	weaverRoutes := router.Group("/api/weavers")
	{
		weaverRoutes.GET("", weaverHandler.GetWeavers)
		weaverRoutes.POST("", weaverHandler.CreateWeaver)
	}

	// Buyer routes
	buyerHandler := NewBuyerHandler(db)
	buyerRoutes := router.Group("/api/buyers")
	{
		buyerRoutes.GET("", buyerHandler.GetBuyers)
		buyerRoutes.POST("", buyerHandler.CreateBuyer)
	}

	// Raw Silk Purchase routes
	rawSilkHandler := NewRawSilkHandler(db)
	rawSilkRoutes := router.Group("/api/raw-silk-purchases")
	{
		rawSilkRoutes.GET("", rawSilkHandler.GetRawSilkPurchases)
		rawSilkRoutes.POST("", rawSilkHandler.CreateRawSilkPurchase)
	}

	// Colouring routes
	colouringHandler := NewColouringHandler(db)
	colouringRoutes := router.Group("/api/colouring")
	{
		colouringRoutes.GET("", colouringHandler.GetColouringBatches)
		colouringRoutes.POST("", colouringHandler.CreateColouringBatch)
	}

	// Inventory routes
	inventoryHandler := NewInventoryHandler(db)
	inventoryRoutes := router.Group("/api/inventory")
	{
		inventoryRoutes.GET("/stock", inventoryHandler.GetInventoryStock)
		inventoryRoutes.GET("/movements", inventoryHandler.GetInventoryMovements)
	}

	// Dashboard routes (placeholder)
	router.GET("/api/dashboard/stats", func(c *gin.Context) {
		stats := map[string]interface{}{
			"total_suppliers":     5,
			"total_weavers":       5,
			"total_buyers":        5,
			"inventory_value":     0,
			"pending_purchases":   0,
			"processing_batches":  0,
		}
		c.JSON(200, stats)
	})
}
