package events

import "time"

type Event struct {
    ID              string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
    Title           string    `json:"title" gorm:"not null"`
    Description     string    `json:"description"`
    Location        string    `json:"location"`
    StartsAt        time.Time `json:"starts_at" gorm:"not null"`
    EndsAt          time.Time `json:"ends_at" gorm:"not null"`
    Quota           int       `json:"quota" gorm:"not null;default:0"`
    RegisteredCount int       `json:"registered_count" gorm:"not null;default:0"`
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}
