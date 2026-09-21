package events

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


// Create handles creating a new event
// @Summary Create a new event
// @Description Create a new event with schedule and seat quota (Requires JWT Authentication)
// @Tags Events
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateEventRequest true "Event Details"
// @Success 201 {object} map[string]interface{} "Event created successfully"
// @Failure 400 {object} map[string]string "Validation error or invalid schedule"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /events/create [post]
func (h *Handler) Create(c *gin.Context) {
	var req CreateEventRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := h.service.CreateEvent(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "event created successfully",
		"data":    res,
	})
}

// GetAll retrieves all available events
// @Summary Get all events
// @Description Retrieve a list of all available events
// @Tags Events
// @Produce json
// @Success 200 {object} map[string]interface{} "List of events"
// @Failure 400 {object} map[string]string "Database error"
// @Router /events [get]
func (h *Handler) GetAll(c *gin.Context) {
	res, err := h.service.GetAllEvents(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "events retrieved successfully",
		"data":    res,
	})
}

// GetByID retrieves a single event by its UUID
// @Summary Get event by ID
// @Description Retrieve detailed information for a single event
// @Tags Events
// @Produce json
// @Param id path string true "Event UUID"
// @Success 200 {object} map[string]interface{} "Event details"
// @Failure 404 {object} map[string]string "Event not found"
// @Router /events/{id} [get]
func (h *Handler) GetByID(c *gin.Context) {
	eventID := c.Param("id")

	event, err := h.service.GetEventByID(c.Request.Context(), eventID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "event not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "event retrieved successfully",
		"data":    event,
	})
}