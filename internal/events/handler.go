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


func (h *Handler) Create(c *gin.Context){
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


func (h * Handler) GetAll(c *gin.Context){
	res,err := h.service.GetAllEvents(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message": "events retrieved successfully",
		"data" : res,
	})
}

func (h *Handler) GetByID(c *gin.Context){
	EventID := c.Param("id")
	

	Event,err := h.service.GetEventByID(c.Request.Context(),EventID)
	if err != nil{
		c.JSON(http.StatusNotFound,gin.H{
			"error" : "Event Not Found",
		})
		return
	}

	c.JSON(http.StatusOK,gin.H{
		"message": "event retrieved successfully",
		"data" : Event,
	})

}