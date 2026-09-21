package events

import (
	"context"

	"gorm.io/gorm"
)


type repository struct {
    db *gorm.DB
}

type Repository interface {
    Create(ctx context.Context, event *Event) error
    FindByID(ctx context.Context, id string) (*Event, error)
	FindAll(ctx context.Context) ([]Event, error)
}



func NewRepository(db *gorm.DB) Repository {
    return &repository{
        db: db,
    }
}

func (r *repository) Create(ctx context.Context, event *Event) error {
	return r.db.WithContext(ctx).Create(event).Error
}

func (r *repository) FindByID(ctx context.Context,id string) (*Event,error){
	var event Event
	err := r.db.WithContext(ctx).Where("id = ?",id).First(&event).Error
	if err != nil{
		return nil,err
	}

	return &event,nil
}


func (r *repository) FindAll(ctx context.Context) ([]Event,error){
	var events []Event
	err := r.db.WithContext(ctx).Find(&events).Error
	if err != nil {
		return nil, err
	}
	return events, nil
}

