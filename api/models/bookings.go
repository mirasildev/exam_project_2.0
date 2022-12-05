package models

import "time"

type Booking struct {
	ID         int64     `json:"id"`
	Arrival    time.Time `json:"arrival"`
	Checkout   time.Time `json:"checkout"`
	RoomID     int64     `json:"room_id"`
	RoomNumber int64     `json:"number_rooms"`
	UserID     int64     `json:"user_id"`
	Booked_at  time.Time `json:"booked_at"`
}

type CreateBookingRequest struct {
	Arrival    time.Time `json:"arrival" binding:"required"`
	Checkout   time.Time `json:"checkout" binding:"required"`
	RoomID     int64     `json:"room_id"`
	RoomNumber int64     `json:"number_rooms"`
}

type CreateBookingResponse struct {
	ID       int64     `json:"id"`
	BookedAt time.Time `json:"booked_at"`
}

type UpdateBookingRequest struct {
	Arrival    time.Time `json:"arrival" binding:"required"`
	Checkout   time.Time `json:"checkout" binding:"required"`
	RoomID     int64     `json:"room_id"`
	RoomNumber int64     `json:"number_rooms"`
}

type GetAllBookingsParams struct {
	Limit  int32 `json:"limit" binding:"required" default:"10"`
	Page   int32 `json:"page" binding:"required" default:"1"`
	RoomID int64 `json:"room_id" binding:"required"`
}
type GetAllBookingsResponse struct {
	Bookings []*Booking `json:"bookings"`
	Count    int32 `json:"count"`
}
