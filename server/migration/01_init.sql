-- +goose Up
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE IF NOT EXISTS rooms (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    password_hash TEXT,
    need_pass BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ  DEFAULT now()
);

CREATE TABLE IF NOT EXISTS room_guests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    room_id UUID NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    guest_name TEXT NOT NULL,
    is_owner BOOLEAN DEFAULT FALSE,
    joined_at TIMESTAMPTZ DEFAULT now(),
    UNIQUE (room_id, guest_name)
);

CREATE TABLE IF NOT EXISTS messages (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    room_id UUID NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    guest_name TEXT NOT NULL,
    content TEXT NOT NULL,
    type TEXT NOT NULL DEFAULT 'text',
    sent_at TIMESTAMPTZ DEFAULT now()
);


CREATE INDEX IF NOT EXISTS idx_rooms_name ON rooms (name);
CREATE INDEX IF NOT EXISTS idx_messages_room_id_sent_at ON messages (room_id, sent_at DESC);


-- +goose Down
DROP INDEX IF EXISTS idx_messages_room_id_sent_at;
DROP INDEX IF EXISTS idx_rooms_name;
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS room_guests;
DROP TABLE IF EXISTS rooms;
DROP EXTENSION IF EXISTS "pgcrypto";
