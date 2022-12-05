package postgres_test

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/mirasildev/exam_project_2.0/storage/repo"
	"github.com/stretchr/testify/require"
)

func createRoom(t *testing.T) int64 {
	id := createHotel(t)
	room, err := strg.Room().Create(&repo.Room{
		RoomNum:       1,
		Type:          repo.RoomTypeSingle,
		Description:   faker.Sentence(),
		HotelID:       id,
		PricePerNight: 100.00,
		Status:        true,
	})
	require.NoError(t, err)
	require.NotEmpty(t, room)

	return room
}

func deleteRoom(id int64, t *testing.T) {
	err := strg.Room().Delete(id)
	require.NoError(t, err)
}


func TestCreateRoom(t *testing.T) {
	id := createRoom(t)
	require.NotEmpty(t, id)
}

func TestGetRoom(t *testing.T) {
	id := createRoom(t)
	
	room, err := strg.Room().Get(id)
	require.NoError(t, err)
	require.NotEmpty(t, room)
}

func TestGetAllRooms(t *testing.T) {
	
	rooms, err := strg.Room().GetAll(&repo.GetAllRoomsParams{
		Limit: 10,
		Page:  1,
		HotelID: 19,
	})
	require.NoError(t, err)
	require.NotEmpty(t, rooms)

}

func TestUpdateRoom(t *testing.T) {
	hotelID := createHotel(t)
	id := createRoom(t)
	room, err := strg.Room().Update(&repo.Room{
		ID:       id,
		RoomNum: 2,
		Type:      repo.RoomTypeDouble,
		Description: faker.Sentence(),
		HotelID: hotelID,
		PricePerNight: 90.00,
		Status: true,
	})
	require.NoError(t, err)
	require.NotEmpty(t, room)

}

func TestDeleteRoom(t *testing.T) {
	id := createRoom(t)
	deleteRoom(id, t)
}

