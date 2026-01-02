package initator

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/filagot/emagne/initator/domain"
	"github.com/filagot/emagne/initator/foundation"
	"github.com/filagot/emagne/internal/config"
	"github.com/filagot/emagne/internal/database/persistancedb"
	"github.com/gin-gonic/gin"
)

// App represents the application dependencies
type App struct {
	Config      *config.Config
	AuthStorage domain.PersistanceLayer
	AuthModule  *domain.Module
	Handler     *domain.Handler
	Router      *gin.Engine
	Server      *http.Server
}

// NewApp creates and initializes the application
func NewApp() (*App, error) {
	// Load configuration
	cfg := config.Load()

	// Run database migrations first
	InitiateMigration(cfg)

	// Initialize database
	log.Println("Initializing database...")
	pgxConn := foundation.InitDB(cfg.DBSource)
	log.Println("Database initialized")

	// Initialize persistence layer

	// Initialize storage layer
	persistanceLayer := domain.InitPersistance(persistancedb.New(pgxConn))

	// Initialize module layer
	moduleLayer := domain.InitModule(persistanceLayer, cfg)

	// Initialize handler layer
	handler := domain.InitHandler(moduleLayer)

	// Initialize Gin router
	router := gin.New()
	//server := gin.New()
	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		origin := cfg.FrontendOrigin
		if origin == "" {
			origin = "*"
		}
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if origin != "*" {
			c.Header("Vary", "Origin")
		}
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	mainGroup := router.Group("/api")
	// Setup routes using domain routing
	domain.InitiateRouting(mainGroup, router, handler, cfg)

	// Create HTTP server
	srv := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: router,
	}

	return &App{
		Config:      cfg,
		AuthStorage: *persistanceLayer,
		AuthModule:  moduleLayer,
		Handler:     handler,
		Router:      router,
		Server:      srv,
	}, nil
}

// InitiateMigration runs database migrations
func InitiateMigration(cfg *config.Config) {
	log.Println("Starting database migration...")

	// Set migration path
	migrationPath := "../db/migrations"

	// Create migration instance
	m := foundation.InitiateMigration(migrationPath, cfg.DBSource)

	// Run migrations
	foundation.UpMigration(m)

	log.Println("Database migration completed successfully!")
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

	log.Println("Server exited")
}
