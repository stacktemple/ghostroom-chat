package middleware

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/stacktemple/realtime-chat/server/auth"
)

const (
	ctxRoomName   = "room_name"
	ctxGuestName  = "guest_name"
	ctxIssuedDate = "issued_date"
)

func JWTGuard(jwtSecret string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Missing token")
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := auth.ParseToken(jwtSecret, tokenStr)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}

		roomName, ok1 := claims["room_name"].(string)
		guestName, ok2 := claims["guest_name"].(string)
		issuedDate, ok3 := claims["issued_date"].(string)
		if !ok1 || !ok2 || !ok3 {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid claims")
		}

		loc, _ := time.LoadLocation("Asia/Bangkok")
		if issuedDate != time.Now().In(loc).Format("2006-01-02") {
			return fiber.NewError(fiber.StatusUnauthorized, "Token expired")
		}

		c.Locals(ctxRoomName, roomName)
		c.Locals(ctxGuestName, guestName)
		c.Locals(ctxIssuedDate, issuedDate)

		return c.Next()
	}
}
