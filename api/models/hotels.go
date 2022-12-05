package models

type Hotel struct {
	ID          int64         `json:"id"`
	Name        string        `json:"hotel_name"`
	Description string        `json:"description"`
	Address     string        `json:"address"`
	ImageUrl    string        `json:"image_url"`
	NumOfRooms  int32         `json:"num_of_rooms"`
	UserID      int64         `json:"user_id"`
	Images      []*HotelImage `json:"images"`
}

type HotelImage struct {
	ImageUrl       string `json:"image_url"`
	SequenceNumber int32  `json:"sequence_number"`
}

type CreateHotelRequest struct {
	Name        string `json:"hotel_name" binding:"required,min=3,max=40"`
	Description string `json:"description" binding:"required"`
	Address     string `json:"address" binding:"required"`
	ImageUrl    string `json:"image_url"`
	NumOfRooms  int32  `json:"num_of_rooms"`
	// UserID      int64         `json:"user_id"`
	Images []*HotelImage `json:"images"`
}

type CreateHotelResponse struct {
	ID int64 `json:"id"`
}

type UpdateHotelRequest struct {
	Name        string `json:"hotel_name" binding:"required,min=3,max=40"`
	Description string `json:"description" binding:"required"`
	Address     string `json:"address" binding:"required"`
	ImageUrl    string `json:"image_url"`
	NumOfRooms  int32  `json:"num_of_rooms"`
	// UserID      int64         `json:"user_id"`
	Images []*HotelImage `json:"images"`
}

type GetAllHotelsParams struct {
	Limit       int32  `json:"limit" binding:"required" default:"10"`
	Page        int32  `json:"page" binding:"required" default:"1"`
	Description string `json:"description"`
}

type GetAllHotelsResponse struct {
	Hotels []*Hotel `json:"hotels"`
	Count  int32    `json:"count"`
}


