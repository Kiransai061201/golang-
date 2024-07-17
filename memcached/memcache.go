package main

import (
    "fmt"
    "log"

    "github.com/bradfitz/gomemcache/memcache"
)

func main() {
    // Initialize a new memcache client
    mc := memcache.New("localhost:11212")

    // Define the key used for caching
    key := "test_key"
    value := "Hello, Memcached!"

    // Set a value in memcache
    err := mc.Set(&memcache.Item{Key: key, Value: []byte(value)})
    if err != nil {
        log.Fatalf("Error setting value in memcache: %v", err)
    }

    // Try to fetch the data from memcache
    item, err := mc.Get(key)
    if err != nil {
        log.Fatalf("Error fetching data from memcache: %v", err)
    }

    // Print the cached data
    fmt.Println("Cached data:")
    fmt.Println(string(item.Value))
}
