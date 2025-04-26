package bloom

import (
	"log"
	"username-check-api/internal/db"

	"github.com/bits-and-blooms/bloom/v3"
)

var filter *bloom.BloomFilter

func InitBloom() {
	filter = bloom.NewWithEstimates(1000000, 0.01)

	preloadUsernames()
}

func preloadUsernames() {
	usernames, err := db.FetchAllUsernames()
	if err != nil {
		log.Println("Failed to preload usernames into Bloom filter:", err)
		return
	}

	for _, username := range usernames {
		filter.AddString(username)
	}
	log.Println("Preloaded", len(usernames), "usernames into Bloom filter.")
}

func AddUsername(username string) {
	filter.AddString(username)
}

func CheckUsername(username string) bool {
	return filter.TestString(username)
}
