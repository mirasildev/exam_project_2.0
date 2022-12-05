package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mirasildev/exam_project_2.0/storage/repo"
)

type bookingRepo struct {
	db *sqlx.DB
}

func NewBooking(db *sqlx.DB) repo.BookingStorageI {
	return &bookingRepo{
		db: db,
	}
}

func (bg *bookingRepo) Create(b *repo.Booking) (*repo.Booking, error) {
	query := `
		INSERT INTO bookings(
			arrival,
			checkout,
			room_id,
			room_number,
			user_id,
			booked_at
		) VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id, booked_at
	`

	row := bg.db.QueryRow(
		query,
		b.Arrival,
		b.Checkout,
		b.RoomID,
		b.RoomNumber,
		b.UserID,
		time.Now(),
	)

	err := row.Scan(&b.ID, &b.Booked_at)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (bg *bookingRepo) Get(id int64) (*repo.Booking, error) {
	var result repo.Booking

	query := `
		SELECT
			id,
			arrival,
			checkout,
			room_id,
			room_number,
			user_id,
			booked_at
		FROM bookings
		WHERE id=$1
	`

	row := bg.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.Arrival,
		&result.Checkout,
		&result.RoomID,
		&result.RoomNumber,
		&result.UserID,
		&result.Booked_at,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (bg *bookingRepo) GetAll(params *repo.GetAllBookingsParams) (*repo.GetAllBookingsResult, error) {
	result := repo.GetAllBookingsResult{
		Bookings: make([]*repo.Booking, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	query := `
		SELECT
			id,
			arrival,
			checkout,
			room_id,
			room_number,
			user_id,
			booked_at
		FROM bookings
		WHERE room_id=$1` + `
		ORDER BY booked_at desc
		` + limit

	rows, err := bg.db.Query(query, params.RoomID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var b repo.Booking

		err := rows.Scan(
			&b.ID,
			&b.Arrival,
			&b.Checkout,
			&b.RoomID,
			&b.RoomNumber,
			&b.UserID,
			&b.Booked_at,
		)
		if err != nil {
			return nil, err
		}

		result.Bookings = append(result.Bookings, &b)
	}

	queryCount := `SELECT count(1) FROM bookings `
	err = bg.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (bg *bookingRepo) Update(b *repo.Booking) (*repo.Booking, error) {
	query := `
		UPDATE bookings SET
			arrival=$1,
			checkout=$2,
			room_id=$3,
			room_number=$4,
			user_id=$5,
			booked_at=$6
		WHERE id=$7
		RETURNING id, arrival, checkout, room_id, 
		room_number, user_id, booked_at
	`
	var result repo.Booking
	err := bg.db.QueryRow(
		query,
		b.Arrival,
		b.Checkout,
		b.RoomID,
		b.RoomNumber,
		b.UserID,
		time.Now(),
		b.ID,
	).Scan(
		&result.ID,
		&result.Arrival,
		&result.Checkout,
		&result.RoomID,
		&result.RoomNumber,
		&result.UserID,
		&result.Booked_at,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (bg *bookingRepo) Delete(id int64) error {

	query := "DELETE FROM bookings WHERE id=$1"
	result, err := bg.db.Exec(query, id)
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
