package events

import "time"

type CreateEventRequest struct {
	Title string `json:"title" binding:"required,min=3"`
	Description string `json:"description"`
	Location string `json:"location" binding:"required,min=3"`
	StartsAt time.Time `json:"starts_at" binding:"required"`
	EndsAt time.Time `json:"ends_at" binding:"required"`
	Quota int `json:"quota" binding:"required,min=1"`

}


type EventResponse struct {
	ID          string    `json:"id" example:"ae57d22b-e67e-43de-adeb-4965a6e3a334"`
	Title       string    `json:"title" example:"Go Concurrency & Backend Masterclass"`
	Description string    `json:"description" example:"Deep dive into goroutines, channels, and row-level locking."`
	Location    string    `json:"location" example:"Jakarta Convention Center"`
	StartsAt    time.Time `json:"starts_at" example:"2026-10-01T09:00:00Z"`
	EndsAt      time.Time `json:"ends_at" example:"2026-10-01T17:00:00Z"`
	Quota       int       `json:"quota" example:"100"`
	CreatedAt   time.Time `json:"created_at" example:"2026-09-21T09:31:17Z"`
}

type EventSuccessResponse struct {
	Message string        `json:"message" example:"event retrieved successfully"`
	Data    EventResponse `json:"data"`
}

type EventsListResponse struct {
	Message string          `json:"message" example:"events retrieved successfully"`
	Data    []EventResponse `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"event not found"`
}