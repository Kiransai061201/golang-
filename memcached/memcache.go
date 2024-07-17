package main

import (
	"fmt"
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	// Initialize a new memcache client
	mc := memcache.New("localhost:11211")

	// Define the key used for caching
	key := "facebook_link_cache_key"

	// Try to fetch the data from memcache
	item, err := mc.Get(key)
	if err != nil {
		log.Fatalf("Error fetching data from memcache: %v", err)
	}

	// Print the cached data
	fmt.Println("Cached data:")
	fmt.Println(string(item.Value))
}
