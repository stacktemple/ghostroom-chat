package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCheck(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	mock.ExpectQuery("SELECT true").
		WillReturnRows(sqlmock.NewRows([]string{"?"}).AddRow(true))

	app := fiber.New()
	healthHandler := HealthHandler{
		AppName: "StackTemple",
		DB:      sqlxDB,
	}
	app.Get("/api/health", healthHandler.Check)

	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
