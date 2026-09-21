package booking

import (
	"context"
)

type Service interface {
	BookEvent(ctx context.Context, eventID,userID string) (*BookingResponse, error)
	GetUserBooking(ctx context.Context,userID string) ([]BookingResponse, error)
}

type service struct {
	repo      Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo:      repo,
	}
}

func (s *service) BookEvent(ctx context.Context, eventID,userID string) (*BookingResponse, error) {
	booking,err := s.repo.BookEvent(ctx,eventID,userID)
	if err != nil {
		return nil,err
	}

	return &BookingResponse{
		ID:        booking.ID,
		EventID:   booking.EventID,
		UserID:    booking.UserID,
		CreatedAt: booking.CreatedAt,
	}, nil
}


func (s *service) GetUserBooking(ctx context.Context,userID string) ([]BookingResponse, error){
	bookings, err := s.repo.FindUserBooking(ctx, userID)
	if err != nil {
		return nil, err
	}
	var responses []BookingResponse
	for _, b := range bookings {
		responses = append(responses, BookingResponse{
			ID:        b.ID,
			EventID:   b.EventID,
			UserID:    b.UserID,
			CreatedAt: b.CreatedAt,
		})
	}
	return responses, nil
}