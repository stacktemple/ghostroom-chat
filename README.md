## Project Goals

This is a guest-only chatroom that works in real-time and was built with Go and PostgreSQL.

- WebSocket-based real-time messaging
- Guest-only access (no login/signup)
- Room-specific guest identity
- JWT-based room authorization (1 day valid)
- All data automatically resets daily at 00:00 GMT+7 (Thai time).
- Password-protected room support
- PostgreSQL as sole persistence (no Redis, no external broker)
- Scalable via internal Go goroutines and channels

---

## Key Requirements

- Guests must choose a name before entering a room
- Guests can join/create a room with or without a password
- Each guest receives a JWT token tied to room + name + date
- Guest names must be unique **per room per day**
- Messages and room joins are saved in PostgreSQL
- At midnight (Thai time), all rooms and messages are deleted
- After 00:00, guests must recreate rooms and rejoin (previous JWTs become invalid)
