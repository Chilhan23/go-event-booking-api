package booking

import (
	"context"
	"errors"

	"example.com/event-app/internal/events"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
    db *gorm.DB
}

type Repository interface {
    BookEvent(ctx context.Context, eventID , userID string)(*Booking,error)
	FindUserBooking(ctx context.Context,userID string)([]Booking,error) 
}

func NewRepository(db *gorm.DB) Repository {
    return &repository{
        db: db,
    }
}



func (r *repository) FindUserBooking(ctx context.Context, userID string) ([]Booking, error) {
	var bookings []Booking
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}

func (r *repository) BookEvent(ctx context.Context, eventID, userID string) (*Booking, error) {
	var booking Booking

	// 1. Mulai transaksi database
	err := r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		// 2. Ambil data event DENGAN ROW-LEVEL LOCK (FOR UPDATE)
		var event events.Event
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", eventID).First(&event).Error; err != nil {
			return err
		}

		// 3. Cek apakah kuota masih ada
		if event.RegisteredCount >= event.Quota {
			return errors.New("event is fully booked")
		}

		// 4. Buat record booking baru (Otomatis gagal jika user sudah pernah booking via UNIQUE constraint di Postgres)
		booking = Booking{
			EventID: eventID,
			UserID:  userID,
		}
		if err := tx.Create(&booking).Error; err != nil {
			return err
		}

		// 5. Tambah registered_count event sebanyak +1
		if err := tx.Model(&events.Event{}).Where("id = ?", eventID).Update("registered_count", gorm.Expr("registered_count + ?", 1)).Error; err != nil {
			return err
		}

		// 6. Return nil = Transaksi Sukses & GORM melakukan COMMIT
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &booking, nil
}

