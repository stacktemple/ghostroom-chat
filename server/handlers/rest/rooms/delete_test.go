package rooms_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestDeleteRoom_Success(t *testing.T) {
	app, mock, handler := setup(t)

	roomID := "room-123"
	roomName := "myroom"
	guestName := "hong"

	// Mock GetRoomByNameToday
	mock.ExpectQuery("SELECT id, password_hash, need_pass").
		WithArgs(roomName).
		WillReturnRows(sqlmock.NewRows([]string{"id", "password_hash", "need_pass"}).
			AddRow(roomID, "", false))

	// Mock IsGuestOwner
	mock.ExpectQuery("SELECT is_owner FROM room_guests").
		WithArgs(roomID, guestName).
		WillReturnRows(sqlmock.NewRows([]string{"is_owner"}).AddRow(true))

	// Mock DeleteRoomByID
	mock.ExpectExec("DELETE FROM rooms").
		WithArgs(roomID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := httptest.NewRequest(http.MethodDelete, "/api/rooms/"+roomName, nil)

	app.Delete("/api/rooms/:name", func(c *fiber.Ctx) error {
		c.Locals("room_name", roomName)
		c.Locals("guest_name", guestName)
		return handler.DeleteRoom(c)
	})

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDeleteRoom_TokenMismatch(t *testing.T) {
	app, _, handler := setup(t)

	paramRoom := "realroom"
	tokenRoom := "fakeroom"
	guest := "hong"

	app.Delete("/api/rooms/:name", func(c *fiber.Ctx) error {
		c.Locals("room_name", tokenRoom)
		c.Locals("guest_name", guest)
		return handler.DeleteRoom(c)
	})

	req := httptest.NewRequest(http.MethodDelete, "/api/rooms/"+paramRoom, nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestDeleteRoom_NotOwner(t *testing.T) {
	app, mock, handler := setup(t)

	roomID := "room-123"
	roomName := "myroom"
	guestName := "notowner"

	// Mock GetRoomByNameToday
	mock.ExpectQuery("SELECT id, password_hash, need_pass").
		WithArgs(roomName).
		WillReturnRows(sqlmock.NewRows([]string{"id", "password_hash", "need_pass"}).
			AddRow(roomID, "", false))

	// Mock IsGuestOwner = false
	mock.ExpectQuery("SELECT is_owner FROM room_guests").
		WithArgs(roomID, guestName).
		WillReturnRows(sqlmock.NewRows([]string{"is_owner"}).AddRow(false))

	app.Delete("/api/rooms/:name", func(c *fiber.Ctx) error {
		c.Locals("room_name", roomName)
		c.Locals("guest_name", guestName)
		return handler.DeleteRoom(c)
	})

	req := httptest.NewRequest(http.MethodDelete, "/api/rooms/"+roomName, nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusForbidden, resp.StatusCode)
}

func TestDeleteRoom_DBError(t *testing.T) {
	app, mock, handler := setup(t)

	roomID := "room-123"
	roomName := "myroom"
	guestName := "hong"

	// Mock GetRoomByNameToday
	mock.ExpectQuery("SELECT id, password_hash, need_pass").
		WithArgs(roomName).
		WillReturnRows(sqlmock.NewRows([]string{"id", "password_hash", "need_pass"}).
			AddRow(roomID, "", false))

	// Mock IsGuestOwner
	mock.ExpectQuery("SELECT is_owner FROM room_guests").
		WithArgs(roomID, guestName).
		WillReturnRows(sqlmock.NewRows([]string{"is_owner"}).AddRow(true))

	// Mock DeleteRoomByID
	mock.ExpectExec("DELETE FROM rooms").
		WithArgs(roomID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := httptest.NewRequest(http.MethodDelete, "/api/rooms/"+roomName, nil)

	app.Delete("/api/rooms/:name", func(c *fiber.Ctx) error {
		c.Locals("room_name", roomName)
		c.Locals("guest_name", guestName)
		return handler.DeleteRoom(c)
	})

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestDeleteRoom_RoomNotFound(t *testing.T) {
	app, mock, handler := setup(t)

	roomName := "ghost-room"
	guestName := "hong"

	// Mock: room not found → return error
	mock.ExpectQuery("SELECT id, password_hash, need_pass").
		WithArgs(roomName).
		WillReturnError(sql.ErrNoRows)

	app.Delete("/api/rooms/:name", func(c *fiber.Ctx) error {
		c.Locals("room_name", roomName)
		c.Locals("guest_name", guestName)
		return handler.DeleteRoom(c)
	})

	req := httptest.NewRequest(http.MethodDelete, "/api/rooms/"+roomName, nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestDeleteRoom_DBErrorOnDelete(t *testing.T) {
	app, mock, handler := setup(t)

	roomID := "room-err"
	roomName := "badroom"
	guestName := "hong"

	// Mock GetRoomByNameToday
	mock.ExpectQuery("SELECT id, password_hash, need_pass").
		WithArgs(roomName).
		WillReturnRows(sqlmock.NewRows([]string{"id", "password_hash", "need_pass"}).
			AddRow(roomID, "", false))

	// Mock IsGuestOwner = true
	mock.ExpectQuery("SELECT is_owner FROM room_guests").
		WithArgs(roomID, guestName).
		WillReturnRows(sqlmock.NewRows([]string{"is_owner"}).AddRow(true))

	// Mock DeleteRoomByID returns error
	mock.ExpectExec("DELETE FROM rooms").
		WithArgs(roomID).
		WillReturnError(assert.AnError)

	app.Delete("/api/rooms/:name", func(c *fiber.Ctx) error {
		c.Locals("room_name", roomName)
		c.Locals("guest_name", guestName)
		return handler.DeleteRoom(c)
	})

	req := httptest.NewRequest(http.MethodDelete, "/api/rooms/"+roomName, nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
