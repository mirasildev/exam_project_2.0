package postgres_test

import (
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/mirasildev/exam_project_2.0/storage/repo"
	"github.com/stretchr/testify/require"
)

func createHotel(t *testing.T) int64 {
	u := createUser(t)
	hotel, err := strg.Hotel().Create(&repo.Hotel{
		Name:        faker.FirstName(),
		Description: faker.Sentence(),
		Address:     faker.MacAddress(),
		ImageUrl:    faker.Sentence(),
		NumOfRooms:  2,
		UserID:      u.ID,
		Images: []*repo.HotelImage{
			{
				ImageUrl:       "url1",
				SequenceNumber: 1,
			},
		},
	})

	require.NoError(t, err)
	require.NotEmpty(t, hotel)

	return hotel
}

func deleteHotel(id int64, t *testing.T) {
	err := strg.Hotel().Delete(id)
	require.NoError(t, err)
}

func TestUpdateHotel(t *testing.T) {
	id := createHotel(t)
	hotel, err := strg.Hotel().Update(&repo.Hotel{
		ID: id,
		Name:        faker.ChineseFirstName(),
		Description: faker.Sentence(),
		Address:     faker.MacAddress(),
		ImageUrl:    "test1",
		NumOfRooms:  12,
		Images: []*repo.HotelImage{
			{
				HotelID:        id,
				ImageUrl:       faker.Word(),
				SequenceNumber: 1,
			},
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, hotel)

	deleteHotel(hotel.ID, t)
}
func TestCreateHotel(t *testing.T) {
	id := createHotel(t)
	deleteHotel(id, t)
}

func TestGetHotel(t *testing.T) {
	id := createHotel(t)

	hotel, err := strg.Hotel().Get(id)
	require.NoError(t, err)
	require.NotEmpty(t, hotel)
}

func TestGetAllHotels(t *testing.T) {

	Notes, err := strg.Hotel().GetAll(&repo.GetAllHotelsParams{
		Limit: 10,
		Page:  1,
	})
	require.NoError(t, err)
	require.NotEmpty(t, Notes)

}

func TestDeleteHotel(t *testing.T) {
	id := createHotel(t)
	deleteHotel(id, t)
}