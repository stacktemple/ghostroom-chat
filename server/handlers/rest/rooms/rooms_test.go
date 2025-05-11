package rooms_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/stacktemple/realtime-chat/server/handlers/rest/rooms"
	"github.com/stacktemple/realtime-chat/server/repository"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (*fiber.App, sqlmock.Sqlmock, *rooms.RoomHandler) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewRoomRepository(sqlxDB)

	h := &rooms.RoomHandler{
		JWTSecret: "test-secret",
		Repo:      repo,
	}

	app := fiber.New()
	app.Post("/api/rooms", h.CreateRoom)
	return app, mock, h
}
