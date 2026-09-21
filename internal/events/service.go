package events

import (
	"context"
	"errors"
)

type Service interface {
	CreateEvent(ctx context.Context, req *CreateEventRequest) (*EventResponse, error)
	GetAllEvents(ctx context.Context) ([]EventResponse, error)
	GetEventByID(ctx context.Context, id string) (*EventResponse, error)
}

type service struct {
	repo      Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo:      repo,
	}
}

func (s *service) CreateEvent(ctx context.Context,req *CreateEventRequest) (*EventResponse,error){
	if req.EndsAt.Before(req.StartsAt) {
		return nil, errors.New("ends_at must be after starts_at")
	}

	event := Event{
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
		StartsAt:    req.StartsAt,
		EndsAt:      req.EndsAt,
		Quota:       req.Quota,
	}

	// Simpan ke 
	if err := s.repo.Create(ctx, &event); err != nil {
		return nil, err
	}

	return &EventResponse{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		Location:    event.Location,
		StartsAt:    event.StartsAt,
		EndsAt:      event.EndsAt,
		Quota:       event.Quota,
		CreatedAt:   event.CreatedAt,
	}, nil

}

func (s *service) GetAllEvents(ctx context.Context)([]EventResponse,error){
	events,err := s.repo.FindAll(ctx)
	if err != nil{
		return nil,err
	}

	var responses []EventResponse
	for _, e := range events {
		responses = append(responses, EventResponse{
			ID:          e.ID,
			Title:       e.Title,
			Description: e.Description,
			Location:    e.Location,
			StartsAt:    e.StartsAt,
			EndsAt:      e.EndsAt,
			Quota:       e.Quota,
			CreatedAt:   e.CreatedAt,
		})
	}
	return responses, nil

}


func (s *service) GetEventByID(ctx context.Context,id string)(*EventResponse,error){
	event,err := s.repo.FindByID(ctx,id)
	if err != nil {
		return nil,err
	}

	return &EventResponse{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		Location:    event.Location,
		StartsAt:    event.StartsAt,
		EndsAt:      event.EndsAt,
		Quota:       event.Quota,
		CreatedAt:   event.CreatedAt,
	}, nil
}



