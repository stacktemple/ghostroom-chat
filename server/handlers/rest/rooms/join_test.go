package rooms_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stacktemple/realtime-chat/server/auth"
	"github.com/stretchr/testify/assert"
)

func TestJoinRoom_Success(t *testing.T) {
	app, mock, handler := setup(t)

	roomID := "room-123"
	roomName := "test-room"
	guestName := "hong"
	password := "abcd1234"
	passwordHash, _ := auth.HashPassword(password)

	// Mock GetRoomByNameToday
	mock.ExpectQuery("SELECT id, password_hash, need_pass").
		WithArgs(roomName).
		WillReturnRows(sqlmock.NewRows([]string{"id", "password_hash", "need_pass"}).
			AddRow(roomID, passwordHash, true))

	// Mock GuestExistsToday = false
	mock.ExpectQuery("SELECT EXISTS").
		WithArgs(roomID, guestName).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	// Mock AddGuest
	mock.ExpectExec("INSERT INTO room_guests").
		WithArgs(roomID, guestName, false).
		WillReturnResult(sqlmock.NewResult(1, 1))

	app.Post("/api/rooms/join", handler.JoinRoom)

	body := `{"name":"test-room", "guest_name":"hong", "password":"abcd1234"}`
	req := httptest.NewRequest(http.MethodPost, "/api/rooms/join", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestJoinRoom_GuestAlreadyExists(t *testing.T) {
	app, mock, handler := setup(t)

	roomID := "room-123"
	roomName := "myroom"
	guestName := "hong"

	// Mock room info without password
	mock.ExpectQuery("SELECT id, password_hash, need_pass").
		WithArgs(roomName).
		WillReturnRows(sqlmock.NewRows([]string{"id", "password_hash", "need_pass"}).
			AddRow(roomID, "", false))

	// Mock guest already exists
	mock.ExpectQuery("SELECT EXISTS").
		WithArgs(roomID, guestName).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	app.Post("/api/rooms/join", handler.JoinRoom)

	body := `{"name":"myroom", "guest_name":"hong"}`
	req := httptest.NewRequest(http.MethodPost, "/api/rooms/join", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusConflict, resp.StatusCode)
}

func TestJoinRoom_IncorrectPassword(t *testing.T) {
	app, mock, handler := setup(t)

	roomID := "room-123"
	roomName := "locked-room"
	correctHash, _ := auth.HashPassword("correct-pass")

	// Mock: room exists with password
	mock.ExpectQuery("SELECT id, password_hash, need_pass").
		WithArgs(roomName).
		WillReturnRows(sqlmock.NewRows([]string{"id", "password_hash", "need_pass"}).
			AddRow(roomID, correctHash, true))

	app.Post("/api/rooms/join", handler.JoinRoom)

	body := `{"name":"locked-room", "guest_name":"hong", "password":"wrong123"}`
	req := httptest.NewRequest(http.MethodPost, "/api/rooms/join", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
