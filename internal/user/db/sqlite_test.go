package db

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/VoC925/go-testify/internal"
	"github.com/VoC925/go-testify/internal/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	userTest = &user.User{
		Name:          "Test_user_name",
		Login:         "loginTest",
		PasswordHash:  "s4vse6v16v4ew6v",
		CreatedAt:     "created",
		LastChangedAt: "changed",
	}
)

func TestAddGetDeleteOK(t *testing.T) {
	db, err := sql.Open("sqlite", "user_test.db")
	require.NoError(t, err)
	defer db.Close()

	store := New(db)
	// add
	id, err := store.Add(userTest)
	require.NoError(t, err)
	assert.Greater(t, id, 0)
	// get
	userActual, err := store.GetByLogin(userTest.Login)
	require.NoError(t, err)
	assert.True(t, reflect.DeepEqual(userTest, userActual))
	// delete
	err = store.Delete(userActual.Login)
	require.NoError(t, err)
	// get
	userActualDel, err := store.GetByLogin(userTest.Login)
	require.ErrorIs(t, err, internal.ErrNotExistUser)
	assert.Nil(t, userActualDel)
}

func TestAddUpdateOK(t *testing.T) {
	db, err := sql.Open("sqlite", "user_test.db")
	require.NoError(t, err)
	defer db.Close()

	store := New(db)
	// add
	id, err := store.Add(userTest)
	require.NoError(t, err)
	assert.Greater(t, id, 0)
	// get
	userActual, err := store.GetByLogin(userTest.Login)
	require.NoError(t, err)
	assert.True(t, reflect.DeepEqual(userTest, userActual))
	// update
	newLogin := "newLogin"
	err = store.UpdateLogin(userActual.Login, newLogin)
	require.NoError(t, err)
	// get
	userActualUp, err := store.GetByLogin(newLogin)
	require.NoError(t, err)
	assert.NotNil(t, userActualUp)
}
