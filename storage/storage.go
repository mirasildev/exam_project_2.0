package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/mirasildev/exam_project_2.0/storage/postgres"
	"github.com/mirasildev/exam_project_2.0/storage/repo"
)

type StorageI interface {
	User() repo.UserStorageI
	Hotel() repo.HotelStorageI
	Room() repo.RoomStorageI
	Booking() repo.BookingStorageI
}

type storagePg struct {
	userRepo    repo.UserStorageI
	hotelRepo   repo.HotelStorageI
	roomRepo    repo.RoomStorageI
	bookingRepo repo.BookingStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		userRepo:    postgres.NewUser(db),
		hotelRepo:   postgres.NewHotel(db),
		roomRepo:    postgres.NewRoom(db),
		bookingRepo: postgres.NewBooking(db),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}

func (s *storagePg) Hotel() repo.HotelStorageI {
	return s.hotelRepo
}

func (s *storagePg) Room() repo.RoomStorageI {
	return s.roomRepo
}

func (s *storagePg) Booking() repo.BookingStorageI {
	return s.bookingRepo
}
