package api

import (
	"encoding/json"
	"net/http"
	"username-check-api/internal/bloom"
	"username-check-api/internal/db"
	"username-check-api/internal/redis"
)

type UsernameRequest struct {
	Username string `json:"username"`
}

type UsernameResponse struct {
	Available bool   `json:"available"`
	Message   string `json:"message,omitempty"`
}

func CheckUsernameHandler(w http.ResponseWriter, r *http.Request) {
	var req UsernameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Step 1: First check the Bloom filter
	if bloom.CheckUsername(req.Username) {
		// If Bloom filter says the username might exist, check Redis first for fast lookup
		if exists := redis.CheckUsername(req.Username); exists {
			json.NewEncoder(w).Encode(UsernameResponse{Available: false})
			return
		}

		// Step 2: If Redis misses, check MongoDB as fallback
		exists := db.UsernameExists(req.Username)
		json.NewEncoder(w).Encode(UsernameResponse{Available: !exists})
		return
	}

	// Step 3: If Bloom filter says it does not exist, check Redis and MongoDB
	if exists := redis.CheckUsername(req.Username); exists {
		json.NewEncoder(w).Encode(UsernameResponse{Available: false})
		return
	}

	exists := db.UsernameExists(req.Username)
	json.NewEncoder(w).Encode(UsernameResponse{Available: !exists})
}

func CreateUsernameHandler(w http.ResponseWriter, r *http.Request) {
	var req UsernameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Step 1: Check the username availability (combined strategy)
	if bloom.CheckUsername(req.Username) {
		// If Bloom filter says the username might exist, check Redis first
		if redis.CheckUsername(req.Username) {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}

		// Check MongoDB as fallback
		if db.UsernameExists(req.Username) {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}
	} else {
		// If Bloom filter misses, check Redis and MongoDB
		if redis.CheckUsername(req.Username) {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}

		if db.UsernameExists(req.Username) {
			http.Error(w, "Username already exists", http.StatusConflict)
			return
		}
	}

	// Step 2: Insert into MongoDB
	if err := db.InsertUsername(req.Username); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Step 3: Update Bloom Filter
	bloom.AddUsername(req.Username)

	// Step 4: Cache the username in Redis for future fast lookups
	redis.SetUsername(req.Username)

	json.NewEncoder(w).Encode(UsernameResponse{Available: true, Message: "Username registered successfully"})

}
