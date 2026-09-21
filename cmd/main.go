package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "example.com/event-app/docs"
	"example.com/event-app/internal/auth"
	"example.com/event-app/internal/booking"
	"example.com/event-app/internal/config"
	"example.com/event-app/internal/database"
	"example.com/event-app/internal/events"
	"example.com/event-app/internal/middleware"
)

// @title Event Booking REST API
// @version 1.0
// @description High-performance Event Management & Concurrency-Safe Ticket Booking REST API in Go.
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
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

    // 1. Wiring Auth Layer
    authRepo := auth.NewRepository(db)
    authService := auth.NewService(authRepo, cfg.JWTSecret)
    authHandler := auth.NewHandler(authService)

    // 2. Wiring Events Layer
    eventRepo := events.NewRepository(db)
    eventService := events.NewService(eventRepo)
    eventHandler := events.NewHandler(eventService)

    // 3. Wiring Booking Layer
    bookingRepo := booking.NewRepository(db)
    bookingService := booking.NewService(bookingRepo)
    bookingHandler := booking.NewHandler(bookingService)

    // 4. Public Auth Routes
    authRoutes := router.Group("/auth")
    {
        authRoutes.POST("/register", authHandler.Register)
        authRoutes.POST("/login", authHandler.Login)
    }

    // 5. Public Event Routes
    router.GET("/events", eventHandler.GetAll)
    router.GET("/events/:id", eventHandler.GetByID)

    // 6. Protected Routes (Requires JWT Authentication)
    protected := router.Group("/")
    protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
    {
        protected.POST("/events/create", eventHandler.Create)
        protected.POST("/events/:id/book", bookingHandler.Book)
        protected.GET("/me/bookings", bookingHandler.GetMyBookings)
    }

    // 7. Swagger Documentation & Root Route
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    router.GET("/", func(c *gin.Context) {
        c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
    })

    // 8. Start HTTP server
    router.Run(":" + cfg.Port)
}







