package postgres_test

import (
	"testing"
	"time"

	"github.com/mirasildev/exam_project_2.0/storage/repo"
	"github.com/stretchr/testify/require"
)

func createBooking(t *testing.T) int64 {
	user := createUser(t)
	roomID := createRoom(t)
	booking, err := strg.Booking().Create(&repo.Booking{
		Arrival: time.Now(),
		Checkout: time.Now().Add(time.Hour * 24),
		RoomID: roomID,
		RoomNumber: 1,
		UserID: user.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, booking)

	return booking.ID
}

func deleteBooking(id int64, t *testing.T) {
	err := strg.Booking().Delete(id)
	require.NoError(t, err)
}


func TestCreateBooking(t *testing.T) {
	id := createBooking(t)
	require.NotEmpty(t, id)

	deleteBooking(id, t)

}

func TestGetBooking(t *testing.T) {
	id := createBooking(t)
	
	booking, err := strg.Booking().Get(id)
	require.NoError(t, err)
	require.NotEmpty(t, booking)
}

func TestGetAllBookings(t *testing.T) {
	
	booking, err := strg.Booking().GetAll(&repo.GetAllBookingsParams{
		Limit: 10,
		Page:  1,
		RoomID: 13,
	})
	require.NoError(t, err)
	require.NotEmpty(t, booking)

}

func TestUpdateBooking(t *testing.T) {
	user := createUser(t)
	roomID := createRoom(t)
	id := createBooking(t)
	booking, err := strg.Booking().Update(&repo.Booking{
		ID:       id,
		Arrival: time.Now(),
		Checkout: time.Now().Add(time.Hour * 48),
		RoomID: roomID,
		RoomNumber: 13,
		UserID: user.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, booking)

}

func TestDeleteBooking(t *testing.T) {
	id := createBooking(t)
	deleteRoom(id, t)
}

