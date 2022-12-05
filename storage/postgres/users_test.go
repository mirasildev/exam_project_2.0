package postgres_test

import (
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/mirasildev/exam_project_2.0/storage/repo"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T) *repo.User {
	phone := faker.Phonenumber()
	user, err := strg.User().Create(&repo.User{
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		PhoneNumber: &phone,
		Email:       faker.Email(),
		Password:    faker.Password(),
		Type:        repo.UserTypeUser,
		CreatedAt:   time.Now(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}

func TestDeleteUser(t *testing.T) {
	u := createUser(t)
	err := strg.User().Delete(u.ID)
	require.NoError(t, err)
}

func updateUser(t *testing.T) *repo.User {
	u := createUser(t)
	user, err := strg.User().Update(&repo.User{
		ID:          u.ID,
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		Email:       faker.Email(),
		Type:        repo.UserTypeUser,
		Password:    faker.Password(),
	})
	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}
func TestCreateUser(t *testing.T) {
	user := createUser(t)
	require.NotEmpty(t, user)
}

func TestGetUser(t *testing.T) {
	c := createUser(t)

	user, err := strg.User().Get(c.ID)
	require.NoError(t, err)
	require.NotEmpty(t, user)
}

func TestUpdateUser(t *testing.T) {
	user := updateUser(t)
	require.NotEmpty(t, user)
}

func TestGetAllUser(t *testing.T) {

	users, err := strg.User().GetAll(&repo.GetAllUsersParams{
		Limit: 10,
		Page:  1,
	})
	require.NoError(t, err)
	require.NotEmpty(t, users)

}
