package booking

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

// Book handles ticket reservation for an event
// @Summary Book an event ticket
// @Description Reserve a ticket for an event with row-level lock concurrency safety (Requires JWT Authentication)
// @Tags Bookings
// @Produce json
// @Security BearerAuth
// @Param id path string true "Event UUID"
// @Success 201 {object} BookingSuccessResponse "Event booked successfully"
// @Failure 400 {object} ErrorResponse "Event fully booked or already booked by user"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Router /events/{id}/book [post]
func (h *Handler) Book(c *gin.Context) {
    eventID := c.Param("id")
    userID := c.GetString("user_id")

    res, err := h.service.BookEvent(c.Request.Context(), eventID, userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "event booked successfully",
        "data":    res,
    })
}

// GetMyBookings retrieves all bookings for the logged-in user
// @Summary Get user booking history
// @Description Retrieve list of all tickets booked by the currently authenticated user
// @Tags Bookings
// @Produce json
// @Security BearerAuth
// @Success 200 {object} UserBookingsSuccessResponse "User booking history"
// @Failure 400 {object} ErrorResponse "Database error"
// @Failure 401 {object} ErrorResponse "Unauthorized"
// @Router /me/bookings [get]
func (h *Handler) GetMyBookings(c *gin.Context) {
    userID := c.GetString("user_id")

    res, err := h.service.GetUserBooking(c.Request.Context(), userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "user bookings retrieved successfully",
        "data":    res,
    })
}


