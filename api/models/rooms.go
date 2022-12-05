package models

type Room struct {
	ID            int64   `json:"id"`
	RoomNum       int32   `json:"room_number"`
	Type          string  `json:"type"`
	Description   string  `json:"description"`
	HotelID       int64   `json:"hotel_id"`
	PricePerNight float64 `json:"price_per_night"`
	Status        bool    `json:"status"`
}  

type CreateRoomRequest struct {
	RoomNum       int32   `json:"room_number"`
	Type          string  `json:"type"`
	Description   string  `json:"description"`
	HotelID       int64   `json:"hotel_id"`
	PricePerNight float64 `json:"price_per_night"`
	Status        bool    `json:"status"`
}

type CreateRoomResponse struct {
	ID int64 `json:"id"`
}

type UpdateRoomRequest struct {
	RoomNum       int32   `json:"room_number"`
	Type          string  `json:"type"`
	Description   string  `json:"description"`
	HotelID       int64   `json:"hotel_id"`
	PricePerNight float64 `json:"price_per_night"`
	Status        bool    `json:"status"`
}

type GetAllRoomsParams struct {
	Limit  int32 `json:"limit" binding:"required" default:"10"`
	Page   int32 `json:"page" binding:"required" default:"1"`
	HotelID int64 `json:"hotel_id"`
}
type GetAllRoomsResponse struct {
	Rooms []*Room `json:"rooms"`
	Count int32 `json:"count"`
}
