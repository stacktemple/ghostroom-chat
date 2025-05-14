package rooms_test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCheckToken_Success(t *testing.T) {
	app, mock, handler := setup(t)

	roomID := "room-123"
	roomName := "test-room"
	guestName := "hong"
	issuedDate := time.Now().Format("2006-01-02") // วันนี้

	// Mock room found
	mock.ExpectQuery("SELECT id, password_hash, need_pass").
		WithArgs(roomName).
		WillReturnRows(sqlmock.NewRows([]string{"id", "password_hash", "need_pass"}).
			AddRow(roomID, "", false))

	// Mock guest exists
	mock.ExpectQuery("SELECT EXISTS").
		WithArgs(roomID, guestName).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	app.Get("/api/rooms/verify-token", func(c *fiber.Ctx) error {
		c.Locals("room_name", roomName)
		c.Locals("guest_name", guestName)
		c.Locals("issued_date", issuedDate)
		return handler.CheckToken(c)
	})

	req := httptest.NewRequest(http.MethodGet, "/api/rooms/verify-token", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestCheckToken_ExpiredDate(t *testing.T) {
	app, _, handler := setup(t)

	issuedDate := time.Now().AddDate(0, 0, -1).Format("2006-01-02") // yesterday
	roomName := "test-room"
	guestName := "hong"

	app.Get("/api/rooms/verify-token", func(c *fiber.Ctx) error {
		c.Locals("room_name", roomName)
		c.Locals("guest_name", guestName)
		c.Locals("issued_date", issuedDate)
		return handler.CheckToken(c)
	})

	req := httptest.NewRequest(http.MethodGet, "/api/rooms/verify-token", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}

func TestCheckToken_RoomNotFound(t *testing.T) {
	app, mock, handler := setup(t)

	roomName := "deleted-room"
	guestName := "hong"
	issuedDate := time.Now().Format("2006-01-02")

	mock.ExpectQuery("SELECT id, password_hash, need_pass").
		WithArgs(roomName).
		WillReturnError(sql.ErrNoRows) // Room not found

	app.Get("/api/rooms/verify-token", func(c *fiber.Ctx) error {
		c.Locals("room_name", roomName)
		c.Locals("guest_name", guestName)
		c.Locals("issued_date", issuedDate)
		return handler.CheckToken(c)
	})

	req := httptest.NewRequest(http.MethodGet, "/api/rooms/verify-token", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode)
}
