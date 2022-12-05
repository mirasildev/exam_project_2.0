package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/mirasildev/exam_project_2.0/storage/repo"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{
		db: db,
	}
}

func (ur *userRepo) Create(user *repo.User) (*repo.User, error) {
	query := `
		INSERT INTO users(
			first_name,
			last_name,
			phone_number,
			email,
			password,
			type
		) VALUES($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`

	row := ur.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.Email,
		user.Password,
		user.Type,
	)

	err := row.Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepo) Get(id int64) (*repo.User, error) {
	var result repo.User

	query := `
		SELECT
			id,
			first_name,
			last_name,
			phone_number,
			email,
			password,
			type,
			created_at,
			updated_at
		FROM users
		WHERE id=$1 AND deleted_at IS NULL
	`

	row := ur.db.QueryRow(query, id)
	err := row.Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.PhoneNumber,
		&result.Email,
		&result.Password,
		&result.Type,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *userRepo) GetAll(params *repo.GetAllUsersParams) (*repo.GetAllUsersResult, error) {
	result := repo.GetAllUsersResult{
		Users: make([]*repo.User, 0),
	}

	offset := (params.Page - 1) * params.Limit

	limit := fmt.Sprintf(" LIMIT %d OFFSET %d ", params.Limit, offset)

	filter := " "
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter += fmt.Sprintf(`
			WHERE first_name ILIKE '%s' OR last_name ILIKE '%s' OR email ILIKE '%s' 
			OR phone_number ILIKE '%s' AND deleted_at IS NULL `,
			str, str, str, str,
		)
	}

	query := `
		SELECT
			id,
			first_name,
			last_name,
			phone_number,
			email,
			password,
			type,
			created_at,
			updated_at
		FROM users
		` + filter + `
		ORDER BY created_at desc
		` + limit

	rows, err := ur.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u repo.User

		err := rows.Scan(
			&u.ID,
			&u.FirstName,
			&u.LastName,
			&u.PhoneNumber,
			&u.Email,
			&u.Password,
			&u.Type,
			&u.CreatedAt,
			&u.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		result.Users = append(result.Users, &u)
	}

	queryCount := `SELECT count(1) FROM users ` + filter
	err = ur.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *userRepo) Update(user *repo.User) (*repo.User, error) {
	query := `
		UPDATE users SET
			first_name=$1,
			last_name=$2,
			phone_number=$3,
			email=$4,
			password=$5,
			type=$6,
			updated_at=$7
		WHERE id=$8
		RETURNING id, first_name, last_name, phone_number, email, 
		password, type, created_at, updated_at
	`
	var result repo.User
	err := ur.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.Email,
		user.Password,
		user.Type,
		time.Now(),
		user.ID,
	).Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.PhoneNumber,
		&result.Email,
		&result.Password,
		&result.Type,
		&result.CreatedAt,
		&result.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *userRepo) Delete(id int64) error {

	queryDeleteBookings := "DELETE FROM bookings WHERE room_id=$1"
	_, err := ur.db.Exec(queryDeleteBookings, id)
	if err != nil {
		return err
	}

	queryDeleteImages := "DELETE FROM hotel_images WHERE hotel_id=$1"
	_, err = ur.db.Exec(queryDeleteImages, id)
	if err != nil {
		return err
	}

	queryDeleteRooms := "DELETE FROM rooms WHERE hotel_id=$1"
	_, err = ur.db.Exec(queryDeleteRooms, id)
	if err != nil {
		return err
	}

	queryDeleteHotels := "DELETE FROM hotels WHERE user_id=$1"
	_, err = ur.db.Exec(queryDeleteHotels, id)
	if err != nil {
		return err
	}

	query := "DELETE FROM users WHERE id=$1"
	result, err := ur.db.Exec(query, id)
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

func (ur *userRepo) GetByEmail(email string) (*repo.User, error) {
	var result repo.User

	query := `
			SELECT 
				id,
				first_name,
				last_name,
				phone_number,
				email,
				password,
				type,
				created_at
			FROM users
			WHERE email=$1
	`

	row := ur.db.QueryRow(query, email)
	err := row.Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.PhoneNumber,
		&result.Email,
		&result.Password,
		&result.Type,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (ur *userRepo) UpdatePassword(req *repo.UpdatePassword) error {
	query := `UPDATE users SET password=$1 WHERE id=$2`

	_, err := ur.db.Exec(query, req.Password, req.UserID)
	if err != nil {
		return err
	}

	return nil
}
