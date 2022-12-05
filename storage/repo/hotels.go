package repo

type Hotel struct {
	ID          int64
	Name        string
	Description string
	Address     string
	ImageUrl    string
	NumOfRooms  int32
	UserID      int64
	Images      []*HotelImage
}

type HotelImage struct {
	ID             int64  `json:"id"`
	HotelID        int64  `json:"hotel_id"`
	ImageUrl       string `json:"image_url"`
	SequenceNumber int32  `json:"sequence_number"`
}

type GetAllHotelsParams struct {
	Limit  int32
	Page   int32
	Search string
}

type GetAllHotelsResult struct {
	Hotels []*Hotel
	Count  int32
}

type HotelStorageI interface {
	Create(h *Hotel) (int64, error)
	Get(id int64) (*Hotel, error)
	GetAll(params *GetAllHotelsParams) (*GetAllHotelsResult, error)
	Update(h *Hotel) (*Hotel, error)
	Delete(id int64) error
}
