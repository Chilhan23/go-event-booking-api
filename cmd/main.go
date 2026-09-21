package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"example.com/event-app/internal/config"
	"example.com/event-app/internal/database"
	"example.com/event-app/internal/auth"
    "example.com/event-app/internal/events"
    "example.com/event-app/internal/middleware"

)


func main() {
    // 1. Load application configuration
    cfg, err := config.Load()
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // 2. Initialize database connection
    db, err := database.Connect(cfg.DatabaseURL)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    log.Println("Database connected successfully")

    // 3. Setup router & dependencies
    router := gin.Default()

    authRepo := auth.NewRepository(db)
    authService := auth.NewService(authRepo, cfg.JWTSecret)
    authHandler := auth.NewHandler(authService)

    // 4. Register routes
    authRoutes := router.Group("/auth")
    {
        authRoutes.POST("/register", authHandler.Register)
        authRoutes.POST("/login", authHandler.Login)
    }

	// 1. Wiring Layer Events
    eventRepo := events.NewRepository(db)
    eventService := events.NewService(eventRepo)
    eventHandler := events.NewHandler(eventService)

    // 2. Public Event Routes (Melihat Event)
    router.GET("/events", eventHandler.GetAll)
    router.GET("/events/:id", eventHandler.GetByID)

    // 3. Protected Routes (Membuat Event Wajib Login)
    protected := router.Group("/")
    protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
    {
        protected.POST("/events/create", eventHandler.Create)
    }


    // 5. Start HTTP server
    router.Run(":" + cfg.Port)
}







