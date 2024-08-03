package handlers

import (
	"github.com/go-redis/redis/v8"
	"net/http"
	"strconv"
)

// CountHandler returns the approximate number of unique users who have called the /ping API
func CountHandler(redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Get the approximate count of unique users
		count, err := redisClient.PFCount(ctx, uniqueUsersKey).Result()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Convert count to string
		countStr := strconv.FormatInt(count, 10)

		// Write the count to the response
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(countStr))
		if err != nil {
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	}
}
