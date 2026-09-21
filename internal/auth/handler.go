package auth

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}


// Register handles HTTP user registration requests
// @Summary Register a new user
// @Description Register a new user account with username, email, and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body RegisterRequest true "User Registration Credentials"
// @Success 201 {object} RegisterSuccessResponse "Registration successful"
// @Failure 400 {object} ErrorResponse "Validation error or username/email taken"
// @Router /auth/register [post]
func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.service.Register(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"data":    res,
	})
}

// Login handles HTTP user authentication requests
// @Summary Login user
// @Description Authenticate user and return signed JWT bearer token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "User Login Credentials"
// @Success 200 {object} LoginSuccessResponse "Login successful with JWT token"
// @Failure 400 {object} ErrorResponse "Invalid request body"
// @Failure 401 {object} ErrorResponse "Invalid credentials"
// @Router /auth/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.service.Login(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"data":    res,
	})
}