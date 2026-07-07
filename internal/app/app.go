package app

import (
	"database/sql"
	"fmt"
	"log"

	"sd3971-go/config"
	"sd3971-go/internal/handlers"
	"sd3971-go/internal/infrastructure"
	"sd3971-go/internal/repositories"
	"sd3971-go/internal/routes"
	"sd3971-go/internal/services"

	"github.com/gin-gonic/gin"
)

// App represents the application instance
type App struct {
	Router *gin.Engine
	DB     *sql.DB
	Config *config.Config
}

// New creates and initializes a new application instance
func New(cfg *config.Config) (*App, error) {
	// Initialize database
	db, err := infrastructure.InitDB(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	// Run migrations
	if err := infrastructure.RunMigrations(db); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Initialize layers
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	// Register routes
	routes.RegisterRoutes(router, productHandler)

	app := &App{
		Router: router,
		DB:     db,
		Config: cfg,
	}

	log.Println("Application initialized successfully")
	return app, nil
}

// Start starts the application server
func (a *App) Start() error {
	addr := fmt.Sprintf(":%d", a.Config.AppPort)
	log.Printf("Starting server on %s\n", addr)
	return a.Router.Run(addr)
}

// Close closes database connection and other resources
func (a *App) Close() error {
	if a.DB != nil {
		return a.DB.Close()
	}
	return nil
}
