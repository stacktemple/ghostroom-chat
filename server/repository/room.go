package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type RoomRepository struct {
	DB *sqlx.DB
}

func NewRoomRepository(db *sqlx.DB) *RoomRepository {
	return &RoomRepository{DB: db}
}

func (r *RoomRepository) RoomExistsToday(name string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM rooms WHERE name = $1 AND created_at::date = CURRENT_DATE)`
	err := r.DB.Get(&exists, query, name)
	return exists, err
}

func (r *RoomRepository) CreateRoom(name string, passwordHash *string, needPass bool) (string, error) {
	var roomID string
	query := `INSERT INTO rooms (name, password_hash, need_pass) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.Get(&roomID, query, name, passwordHash, needPass)
	return roomID, err
}

func (r *RoomRepository) AddGuest(roomID string, guestName string, isOwner bool) error {
	query := `INSERT INTO room_guests (room_id, guest_name, is_owner) VALUES ($1, $2, $3)`
	_, err := r.DB.Exec(query, roomID, guestName, isOwner)
	return err
}

type RoomInfo struct {
	ID        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	NeedPass  bool      `db:"need_pass" json:"need_pass"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

func (r *RoomRepository) GetTodayRooms() ([]RoomInfo, error) {
	var rooms []RoomInfo
	query := `SELECT id, name, need_pass, created_at FROM rooms WHERE created_at::date = CURRENT_DATE ORDER BY created_at DESC`
	err := r.DB.Select(&rooms, query)
	return rooms, err
}

type RoomDetail struct {
	ID           string `db:"id"`
	PasswordHash string `db:"password_hash"`
	NeedPass     bool   `db:"need_pass"`
}

func (r *RoomRepository) GetRoomByNameToday(name string) (*RoomDetail, error) {
	var room RoomDetail
	query := `SELECT id, password_hash, need_pass FROM rooms WHERE name = $1 AND created_at::date = CURRENT_DATE`
	err := r.DB.Get(&room, query, name)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepository) GuestExistsToday(roomID string, guestName string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM room_guests WHERE room_id = $1 AND guest_name = $2)`
	err := r.DB.Get(&exists, query, roomID, guestName)
	return exists, err
}

func (r *RoomRepository) IsGuestOwner(roomID, guestName string) (bool, error) {
	var isOwner bool
	query := `SELECT is_owner FROM room_guests WHERE room_id = $1 AND guest_name = $2`
	err := r.DB.Get(&isOwner, query, roomID, guestName)
	return isOwner, err
}

func (r *RoomRepository) DeleteRoomByID(roomID string) error {
	query := `DELETE FROM rooms WHERE id = $1`
	_, err := r.DB.Exec(query, roomID)
	return err
}
