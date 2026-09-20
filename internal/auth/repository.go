package auth

import (
	"context"

	"gorm.io/gorm"
)

type repository struct {
    db *gorm.DB
}

type Repository interface {
    Create(ctx context.Context, user *User) error
    FindByEmail(ctx context.Context, email string) (*User, error)
    FindByUsername(ctx context.Context, username string) (*User, error)
    FindByID(ctx context.Context, id string) (*User, error)
}



func NewRepository(db *gorm.DB) Repository {
    return &repository{
        db: db,
    }
}


func (r *repository) Create(ctx context.Context, user *User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) FindByUsername(ctx context.Context,username string) (*User,error){
	var user User
	err := r.db.WithContext(ctx).Where("username = ?",username).First(&user).Error
	if err != nil{
		return nil,err
	}

	return &user,nil
}

func (r *repository) FindByID(ctx context.Context,id string) (*User,error){
	var user User
	err := r.db.WithContext(ctx).Where("id = ?",id).First(&user).Error
	if err != nil{
		return nil,err
	}

	return &user,nil
}