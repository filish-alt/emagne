package initator

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/filagot/emagne/internal/config"
	"github.com/filagot/emagne/internal/database/persistancedb"
	"github.com/filagot/emagne/internal/handler/rest"
	"github.com/filagot/emagne/internal/middleware"
	"github.com/filagot/emagne/internal/module/auth"
	"github.com/filagot/emagne/internal/storage/auth"
	"github.com/gin-gonic/gin"
)

// App represents the application dependencies
type App struct {
	Config      *config.Config
	DB          *persistancedb.PersistenceDB
	AuthStorage storage.AuthStorage
	AuthModule  module.AuthModule
	Handler     *rest.Handler
	Router      *gin.Engine
	Server      *http.Server
}

// NewApp creates and initializes the application
func NewApp() (*App, error) {
	// Load configuration
	cfg := config.Load()

	// Initialize database
	db, err := persistancedb.New(cfg.DBSource)
	if err != nil {
		return nil, err
	}

	// Test database connection
	if err := db.Ping(context.Background()); err != nil {
		return nil, err
	}

	// Initialize storage layer
	authStorage := auth.NewAuthStorage(db.Queries)

	// Initialize module layer
	authModule := auth.NewAuthModule(authStorage, cfg)

	// Initialize handler layer
	handler := rest.NewHandler(authModule)

	// Initialize Gin router
	router := gin.Default()

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// Setup routes
	setupRoutes(router, handler, cfg)

	// Create HTTP server
	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: router,
	}

	return &App{
		Config:      cfg,
		DB:          db,
		AuthStorage: authStorage,
		AuthModule:  authModule,
		Handler:     handler,
		Router:      router,
		Server:      server,
	}, nil
}

// setupRoutes configures all the application routes
func setupRoutes(router *gin.Engine, handler *rest.Handler, cfg *config.Config) {
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Auth routes
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", handler.Register)
		authGroup.POST("/login", handler.Login)
	}

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware(cfg))
	{
		protected.GET("/profile", handler.GetProfile)
	}
}

// Start starts the application server
func (app *App) Start() {
	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on %s", app.Config.ServerAddress)
		if err := app.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Give outstanding requests a deadline for completion
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Attempt graceful shutdown
	if err := app.Server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	// Close database connection
	if err := app.DB.Close(); err != nil {
		log.Printf("Error closing database: %v", err)
	}

	log.Println("Server exited")
}

// Close gracefully closes the application
func (app *App) Close() error {
	return app.DB.Close()
}
