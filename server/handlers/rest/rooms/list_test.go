package rooms_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestListTodayRooms_Success(t *testing.T) {
	app, mock, handler := setup(t)

	// Mock result: 2 rooms
	mock.ExpectQuery("SELECT name, need_pass, created_at").
		WillReturnRows(
			sqlmock.NewRows([]string{"name", "need_pass", "created_at"}).
				AddRow("Room A", false, time.Now()).
				AddRow("Room B", true, time.Now()),
		)

	app.Get("/api/rooms/today", handler.ListTodayRooms)

	req := httptest.NewRequest(http.MethodGet, "/api/rooms/today", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestListTodayRooms_DBError(t *testing.T) {
	app, mock, handler := setup(t)

	// Mock query error
	mock.ExpectQuery("SELECT name, need_pass, created_at").
		WillReturnError(assert.AnError)

	app.Get("/api/rooms/today", handler.ListTodayRooms)

	req := httptest.NewRequest(http.MethodGet, "/api/rooms/today", nil)
	resp, err := app.Test(req)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
