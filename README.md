# 👻 GhostRoom Chat

A safe, anonymous space to chat freely.
Create open or password-protected rooms and start the conversation.

👉 **Try It Now:** [https://ghostroom.stacktemple.com](https://ghostroom.stacktemple.com)

GhostRoom Chat is an open-source, real-time chatroom system  
crafted for **anonymous gatherings and hidden conversations**.

Anyone can **step into a room** or **create one without revealing their identity**—no login, no sign-up.

Each **browser session becomes a unique guest**, verified by a **token valid for a single day**.  
At **midnight (GMT+7)**, all rooms, guests, and messages are **cleansed**, as if they were never there.

This design embraces **temporary presence** and **vanishing dialogue**,  
perfect for fleeting talks, secret strategies, or casual discussions with no strings attached.

## Key Features

- Real-time messaging over WebSocket
- Guest-only access without registration
- One guest identity per browser session, valid for one day
- Create open or password-protected rooms
- Automatic data reset at midnight (GMT+7)
- Stateless architecture powered by JWT
- PostgreSQL as the only persistence layer (no Redis)
- Fully dockerized for both development and production

## Technology Stack

- **Frontend:** React, Vite, TypeScript, Tailwind CSS, Nginx
- **Backend:** Go, Fiber, WebSocket, JWT, sqlx
- **Database:** PostgreSQL 16
- **Deployment:** Docker, Traefik (for reverse proxy and HTTPS)
- **Dev Tools:** Goose for SQL migration management

## Project Structure

- [/client](./client) for Frontend
- [/server](./server/) for Backend

## Try in Your Local

1. Add environment variables for server and client.

### Example: `server/.env`

```bash
PORT=3000
JWT_SECRET=temple-secret
DATABASE_URL=postgres://st-user:st-pass@ghostroom-db:5432/chat-db
```

### Example: client/.env

```bash
VITE_API_URL=http://localhost:3000/api
VITE_WS_BASE_URL=ws://localhost:3000/ws
```

2. Start all services with Docker Compose.

```bash
   docker compose up -d --build
```

3. Run database migration (requires [Goose](https://github.com/pressly/goose) installed):

```bash
    cd server
    make migrate-up
```

### Access Points

- **Client (Web App)** - http://localhost:3001
- **Server API (Health Check)** - http://localhost:3000/api/health

## Production Deployment

See [docker-compose.prod.yaml](./docker-compose.prod.yaml) for production deployment.
Requires **Traefik** with external Docker network configured as `stacktemple_network`.

## Author

Built with ❤️ by [Stacktemple](https://github.com/stacktemple).

Author: Thitiphum Chaikarnjanakit (stacktemple@gmail.com)
