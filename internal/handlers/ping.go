package handlers

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"net/http"
	"time"
)

var (
	lockKey         = "ping_lock"
	lockTimeout     = 10 * time.Second
	rateLimitWindow = 60 * time.Second
	maxRequests     = 2
)

func PingHandler(redisClient *redis.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := r.Header.Get("Username")
		ctx := r.Context()

		// Acquire the lock
		acquired, err := acquireLock(ctx, redisClient, username)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if !acquired {
			http.Error(w, "Resource is currently locked", http.StatusTooManyRequests)
			return
		}
		defer releaseLock(ctx, redisClient, username)

		// Rate limit check
		allowed, err := checkRateLimit(ctx, redisClient, username)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if !allowed {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		// Update total call count
		err = updateTotalCallCount(ctx, redisClient, username)
		if err != nil {
			http.Error(w, "Error updating call count", http.StatusInternalServerError)
			return
		}

		// Simulate a long-running operation
		time.Sleep(5 * time.Second)
	}
}

// acquireLock attempts to acquire a lock for the given user
func acquireLock(ctx context.Context, client *redis.Client, username string) (bool, error) {
	result, err := client.SetNX(ctx, lockKey, username, lockTimeout).Result()
	if err != nil {
		return false, err
	}
	return result, nil
}

// releaseLock releases the lock for the given user
func releaseLock(ctx context.Context, client *redis.Client, username string) {
	currentHolder, err := client.Get(ctx, lockKey).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return
	}
	if currentHolder == username {
		client.Del(ctx, lockKey)
	}
}

// checkRateLimit checks if the user has exceeded the rate limit
func checkRateLimit(ctx context.Context, client *redis.Client, username string) (bool, error) {
	key := "rate_limit:" + username
	count, err := client.Incr(ctx, key).Result()
	if err != nil {
		return false, err
	}
	if count == 1 {
		// Set the expiration for the rate limit window
		client.Expire(ctx, key, rateLimitWindow)
	}
	return count <= int64(maxRequests), nil
}

// Increment total call count for the user
func updateTotalCallCount(ctx context.Context, client *redis.Client, username string) error {
	totalKey := "user:ping:total:" + username
	_, err := client.Incr(ctx, totalKey).Result()
	return err
}
