package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mirasildev/exam_project_2.0/storage/repo"
)

type roomRepo struct {
	db *sqlx.DB
}

func NewRoom(db *sqlx.DB) repo.RoomStorageI {
	return &roomRepo{
		db: db,
	}
}

func (rm *roomRepo) Create(room *repo.Room) (int64, error) {
	query := `
		INSERT INTO rooms(
			room_number,
			type,
			description,
			hotel_id,
			price_per_night,
			status
		) VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id
	`

	row := rm.db.QueryRow(
		query,
		room.RoomNum,
		room.Type,
		room.Description,
		room.HotelID,
		room.PricePerNight,
		room.Status,
	)

	err := row.Scan(&room.ID)
	if err != nil {
		return 0, err
	}

	return room.ID, nil
}

func (rm *roomRepo) Get(id int64) (*repo.Room, error) {
	var result repo.Room

	query := `
		SELECT
			id,
			room_number,
			type,
			description,
			hotel_id,
			price_per_night,
			status
		FROM rooms
		WHERE id=$1
	`

	row := rm.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.RoomNum,
		&result.Type,
		&result.Description,
		&result.HotelID,
		&result.PricePerNight,
		&result.Status,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (rm *roomRepo) GetAll(params *repo.GetAllRoomsParams) (*repo.GetAllRoomsResult, error) {

	result := repo.GetAllRoomsResult{
		Rooms: make([]*repo.Room, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	query := `
		SELECT
			id,
			room_number,
			type,
			description,
			hotel_id,
			price_per_night,
			status
		FROM rooms
		WHERE hotel_id=$1
		` + limit

	rows, err := rm.db.Query(query, params.HotelID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var r repo.Room

		err := rows.Scan(
			&r.ID,
			&r.RoomNum,
			&r.Type,
			&r.Description,
			&r.HotelID,
			&r.PricePerNight,
			&r.Status,
		)
		if err != nil {
			return nil, err
		}

		result.Rooms = append(result.Rooms, &r)
	}

	queryCount := `SELECT count(1) FROM rooms `
	err = rm.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (rm *roomRepo) Update(room *repo.Room) (*repo.Room, error) {
	query := `
		UPDATE rooms SET
			room_number=$1,
			type=$2,
			description=$3,
			hotel_id=$4,
			price_per_night=$5,
			status=$6
		WHERE id=$7
		RETURNING id, room_number, type, description, hotel_id, price_per_night, status
	`
	var result repo.Room
	err := rm.db.QueryRow(
		query,
		room.RoomNum,
		room.Type,
		room.Description,
		room.HotelID,
		room.PricePerNight,
		room.Status,
		room.ID,
	).Scan(
		&result.ID,
		&result.RoomNum,
		&result.Type,
		&result.Description,
		&result.HotelID,
		&result.PricePerNight,
		&result.Status,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (rm *roomRepo) Delete(id int64) error {

	query := "DELETE FROM rooms WHERE id=$1"
	result, err := rm.db.Exec(query, id)
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
