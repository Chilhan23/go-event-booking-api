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


