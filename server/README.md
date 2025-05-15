# GhostRoom Server

Backend of GhostRoom Chat built with **Go**, **Fiber**, **WebSocket**, and **PostgreSQL**.

## Environment Variables

Create a `.env` file:

```env
PORT=3000
JWT_SECRET=temple-secret
DATABASE_URL=postgres://st-user:st-pass@localhost:6000/chat-db
```

## Local Development

1. Install dependencies (Go modules):

```bash
    go mod tidy
```

2. Make sure PostgreSQL is running and accessible via `DATABASE_URL`.
3. Run database migration (requires [Goose](https://github.com/pressly/goose) ):

```bash
    make migrate-up
```

4. Start the server:

```bash
    go run .
```

## Available Endpoints

- **REST API** - _accessible_ via `/api`
- **WebSocket** - accessible via `/ws`

By default, the server runs at: http://localhost:3000
