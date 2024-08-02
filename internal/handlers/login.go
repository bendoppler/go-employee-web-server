package handlers

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/go-redis/redis/v8"
	"net/http"
	"strconv"
	"time"
)

const sessionTTL = time.Hour * 24

func LoginHandler(redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Unable to parse form", http.StatusBadRequest)
				return
			}

			username := r.FormValue("username")
			password := r.FormValue("password")

			if !validateUser(username, password) {
				http.Error(w, "Invalid credentials", http.StatusUnauthorized)
				return
			}
			sessionID := generateSessionID(username, password)
			err = storeSession(username, sessionID, redisClient)
			if err != nil {
				http.Error(w, "Error storing session", http.StatusInternalServerError)
				return
			}
			http.SetCookie(
				w, &http.Cookie{
					Name:    "session_id",
					Value:   sessionID,
					Path:    "/",
					Expires: time.Now().Add(sessionTTL),
				},
			)
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		}
		renderTemplate(w, "login", nil)
	}
}

func validateUser(username, password string) bool {
	// TOFIX: - Will validate login after sing-up flow is finished
	return true
}

func generateSessionID(username, password string) string {
	timestamp := time.Now().Unix()
	data := username + password + strconv.FormatInt(timestamp, 10)

	// Create a SHA-256 hash of the combined data
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

func storeSession(username, sessionID string, redisClient *redis.Client) error {
	ctx := context.Background()
	// Store session ID in Redis with the username as the key
	return redisClient.Set(ctx, username, sessionID, 24*time.Hour).Err()
}
