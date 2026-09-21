package booking

import "time"

type Booking struct {
    ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    EventID   string    `json:"event_id" gorm:"type:uuid;not null"`
    UserID    string    `json:"user_id" gorm:"type:uuid;not null"`
    CreatedAt time.Time `json:"created_at"`
}
