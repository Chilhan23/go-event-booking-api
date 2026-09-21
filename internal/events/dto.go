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


type EventResponse struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Location string `json:"location"`
	StartsAt time.Time `json:"starts_at"`
	EndsAt time.Time `json:"ends_at"`
	Quota int `json:"quota"`
	CreatedAt time.Time `json:"created_at"`
}