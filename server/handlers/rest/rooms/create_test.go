package rooms_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/assert"
)

func TestCreateRoom_Success(t *testing.T) {
	app, mock, _ := setup(t)

	// Mock RoomExistsToday
	mock.ExpectQuery("SELECT EXISTS").
		WithArgs("test-room").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	// Mock CreateRoom (returning room ID)
	mock.ExpectQuery("INSERT INTO rooms").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("room-uuid"))

	// Mock AddGuest
	mock.ExpectExec("INSERT INTO room_guests").
		WithArgs("room-uuid", "hong", true).
		WillReturnResult(sqlmock.NewResult(1, 1))

	body := `{"name": "test-room", "guest_name": "hong", "password": "abcd1234"}`
	req := httptest.NewRequest(http.MethodPost, "/api/rooms", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}

func TestCreateRoom_Conflict(t *testing.T) {
	app, mock, _ := setup(t)

	mock.ExpectQuery("SELECT EXISTS").
		WithArgs("test-room").
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(true))

	body := `{"name": "test-room", "guest_name": "hong"}`
	req := httptest.NewRequest(http.MethodPost, "/api/rooms", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusConflict, resp.StatusCode)
}

func TestCreateRoom_MissingFields(t *testing.T) {
	app, _, _ := setup(t)

	body := `{"name": "", "guest_name": ""}`
	req := httptest.NewRequest(http.MethodPost, "/api/rooms", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateRoom_DBErrorOnExists(t *testing.T) {
	app, mock, _ := setup(t)

	mock.ExpectQuery("SELECT EXISTS").
		WithArgs("test-room").
		WillReturnError(assert.AnError) // simulate DB error

	body := `{"name": "test-room", "guest_name": "hong"}`
	req := httptest.NewRequest(http.MethodPost, "/api/rooms", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}

func TestCreateRoom_BadJSON(t *testing.T) {
	app, _, _ := setup(t)

	body := `invalid-json`
	req := httptest.NewRequest(http.MethodPost, "/api/rooms", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestCreateRoom_ShortPassword(t *testing.T) {
	app, _, _ := setup(t)

	body := `{"name": "room", "guest_name": "hong", "password": "12"}`
	req := httptest.NewRequest(http.MethodPost, "/api/rooms", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}
