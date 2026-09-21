package booking

import "time"

type BookingResponse struct {
	ID        string    `json:"id" example:"2c943e11-884a-4ecb-99f1-d0831707c220"`
	EventID   string    `json:"event_id" example:"ae57d22b-e67e-43de-adeb-4965a6e3a334"`
	UserID    string    `json:"user_id" example:"5b676171-999a-4d62-8506-c0841707b118"`
	CreatedAt time.Time `json:"created_at" example:"2026-09-21T11:00:00Z"`
}

type BookingSuccessResponse struct {
	Message string          `json:"message" example:"event booked successfully"`
	Data    BookingResponse `json:"data"`
}

type UserBookingsSuccessResponse struct {
	Message string            `json:"message" example:"user bookings retrieved successfully"`
	Data    []BookingResponse `json:"data"`
}

type ErrorResponse struct {
	Error string `json:"error" example:"event is fully booked"`
}
