package repo

const (
	RoomTypeSingle = "single"
	RoomTypeDouble = "double"
	RoomTypeFamily = "family"
)

type Room struct {
	ID            int64
	RoomNum       int32
	Type          string
	Description   string
	HotelID       int64
	PricePerNight float64
	Status        bool
}

type GetAllRoomsParams struct {
	Limit   int32
	Page    int32
	HotelID int64
}
type GetAllRoomsResult struct {
	Rooms []*Room
	Count int32
}

type RoomStorageI interface {
	Create(room *Room) (int64, error)
	Get(id int64) (*Room, error)
	GetAll(params *GetAllRoomsParams) (*GetAllRoomsResult, error)
	Update(room *Room) (*Room, error)
	Delete(id int64) error
}
