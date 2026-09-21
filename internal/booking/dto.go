package booking

import "time"

type BookingResponse struct {
    ID        string    `json:"id"`
    EventID   string    `json:"event_id"`
    UserID    string    `json:"user_id"`
    CreatedAt time.Time `json:"created_at"`
}
