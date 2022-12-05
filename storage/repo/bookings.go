package repo

import "time"

type Booking struct {
	ID         int64
	Arrival    time.Time
	Checkout   time.Time
	RoomID     int64
	RoomNumber int64
	UserID     int64
	Booked_at  time.Time
}

type GetAllBookingsParams struct {
	Limit  int32
	Page   int32
	Search string
	RoomID int64
}
type GetAllBookingsResult struct {
	Bookings []*Booking
	Count   int32
}

type BookingStorageI interface {
	Create(b *Booking) (*Booking, error)
	Delete(id int64) error
	Update(b *Booking) (*Booking, error)
	Get(id int64) (*Booking, error)
	GetAll(params *GetAllBookingsParams) (*GetAllBookingsResult, error)
}
