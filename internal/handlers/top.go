package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"go-employee-web-server/internal/models"
	"net/http"
	"sort"
)

// TopHandler returns the top 10 users who have called the /ping API the most
func TopHandler(redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		// Fetch all user call counts
		callCounts, err := getAllUserCallCounts(ctx, redisClient)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		// Sort users by call count
		topUsers := getTopUsers(callCounts, 10)

		// Render the top users as a JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(topUsers); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
		}
	}
}

// getAllUserCallCounts fetches all user call counts from Redis
func getAllUserCallCounts(ctx context.Context, client *redis.Client) (map[string]int64, error) {
	pattern := "user:ping:total:*"
	keys, err := client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, err
	}

	callCounts := make(map[string]int64)
	for _, key := range keys {
		username := key[len("user:ping:total:"):]
		count, err := client.Get(ctx, key).Int64()
		if err != nil {
			if errors.Is(err, redis.Nil) {
				count = 0
			} else {
				return nil, err
			}
		}
		callCounts[username] = count
	}
	return callCounts, nil
}

// getTopUsers sorts users by call count and returns the top N
func getTopUsers(callCounts map[string]int64, topN int) []models.UserCallCount {
	var userCallCounts []models.UserCallCount
	for username, count := range callCounts {
		userCallCounts = append(userCallCounts, models.UserCallCount{Username: username, CallCount: count})
	}

	sort.Slice(
		userCallCounts, func(i, j int) bool {
			return userCallCounts[i].CallCount > userCallCounts[j].CallCount
		},
	)

	if len(userCallCounts) > topN {
		userCallCounts = userCallCounts[:topN]
	}

	return userCallCounts
}
