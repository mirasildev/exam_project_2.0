package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mirasildev/exam_project_2.0/storage/repo"
)

type hotelRepo struct {
	db *sqlx.DB
}

func NewHotel(db *sqlx.DB) repo.HotelStorageI {
	return &hotelRepo{
		db: db,
	}
}

func (ht *hotelRepo) Create(h *repo.Hotel) (int64, error) {
	query := `
		INSERT INTO hotels(
			hotel_name,
			description,
			address, 
			image_url,
			num_of_rooms,
			user_id
		) VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	row := ht.db.QueryRow(
		query,
		h.Name,
		h.Description,
		h.Address,
		h.ImageUrl,
		h.NumOfRooms,
		h.UserID,
	)
	var hotelID int64
	err := row.Scan(&hotelID)
	if err != nil {
		return 0, err
	}

	queryInsertImage := `
		INSERT INTO hotel_images (
			hotel_id,
			image_url,
			sequence_number
		) VALUES ($1, $2, $3)
	`

	for _, image := range h.Images {
		_, err := ht.db.Exec(
			queryInsertImage,
			hotelID,
			image.ImageUrl,
			image.SequenceNumber,
		)
		if err != nil {
			return 0, err
		}
	}

	return hotelID, nil  
}

func (ht *hotelRepo) Get(id int64) (*repo.Hotel, error) {
	var result repo.Hotel

	query := `
		SELECT
			id,
			hotel_name,
			description,
			address, 
			image_url,
			num_of_rooms,
			user_id
		FROM hotels
		WHERE id=$1
	`

	row := ht.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.Name,
		&result.Description,
		&result.Address,
		&result.ImageUrl,
		&result.NumOfRooms,
		&result.UserID,
	)
	if err != nil {
		return nil, err
	}

	queryImages := `
		SELECT 
			id,
			image_url,
			sequence_number
		FROM hotel_images
		WHERE hotel_id=$1
	`

	rows, err := ht.db.Query(queryImages, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var image repo.HotelImage

		err := rows.Scan(
			&image.ID,
			&image.ImageUrl,
			&image.SequenceNumber,
		)
		if err != nil {
			return nil, err
		}
		result.Images = append(result.Images, &image)
	}

	return &result, nil
}

func (ht *hotelRepo) GetAll(params *repo.GetAllHotelsParams) (*repo.GetAllHotelsResult, error) {
	result := repo.GetAllHotelsResult{
		Hotels: make([]*repo.Hotel, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := "WHERE true"
	if params.Search != "" {
		filter += " AND description ilike '%" + params.Search + "%' "
	}

	query := `
		SELECT
			id,
			hotel_name,
			description,
			address, 
			image_url,
			num_of_rooms,
			user_id
		FROM hotels
		` + filter + limit

	rows, err := ht.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var h repo.Hotel

		err := rows.Scan(
			&h.ID,
			&h.Name,
			&h.Address,
			&h.Description,
			&h.ImageUrl,
			&h.NumOfRooms,
			&h.UserID,
		)
		if err != nil {
			return nil, err
		}

		result.Hotels = append(result.Hotels, &h)
	}

	queryCount := `SELECT count(1) FROM hotels ` + filter
	err = ht.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ht *hotelRepo) Update(h *repo.Hotel) (*repo.Hotel, error) {
	query := `
		UPDATE hotels SET
			hotel_name=$1,
			description=$2,
			address=$3, 
			image_url=$4,
			num_of_rooms=$5
		WHERE id=$6
		RETURNING id, hotel_name, description, address, 
		image_url, num_of_rooms, user_id
	`
	var result repo.Hotel
	err := ht.db.QueryRow(
		query,
		h.Name,
		h.Description,
		h.Address,
		h.ImageUrl,
		h.NumOfRooms,
		h.ID,
	).Scan(
		&result.ID,
		&result.Name,
		&result.Description,
		&result.Address,
		&result.ImageUrl,
		&result.NumOfRooms,
		&result.UserID,
	)
	if err != nil {
		return nil, err
	}

	queryDeleteImages := `DELETE FROM hotel_images WHERE hotel_id=$1`
	_, err = ht.db.Exec(queryDeleteImages, result.ID)
	if err != nil {
		return nil ,err
	}

	queryInsertImage := `
		INSERT INTO product_image (
			hotel_id,
			image_url,
			sequence_number
		) VALUES ($1, $2, $3)
	`

	for _, image := range result.Images {
		_, err := ht.db.Exec(
			queryInsertImage,
			result.ID,
			image.ImageUrl,
			image.SequenceNumber,
		)
		if err != nil {
			return nil, err
		}
	}


	return &result, nil
}

func (ht *hotelRepo) Delete(id int64) error {

	queryRooms := "DELETE FROM rooms WHERE hotel_id=$1"
	_, err := ht.db.Exec(queryRooms, id)
	if err != nil {
		return err
	}

	queryImages := "DELETE FROM hotel_images WHERE hotel_id=$1"
	_, err = ht.db.Exec(queryImages, id)
	if err != nil {
		return err
	}

	query := "DELETE FROM hotels WHERE id=$1"
	result, err := ht.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsCount == 0 {
		return sql.ErrNoRows
	}

	return nil
}