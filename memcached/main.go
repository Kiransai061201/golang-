package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	// Initialize a new memcache client
	mc := memcache.New("localhost:11211")

	// Define the key for memcache
	key := "facebook_link_cache_key"

	// Try to fetch the data from memcache first
	item, err := mc.Get(key)
	if err == nil {
		// If data exists in cache, use it
		fmt.Println("Fetching data from cache...")
		fmt.Println("Cached result:", string(item.Value))
		return
	}

	// If data is not found in cache, make the API call
	url := "https://facebook17.p.rapidapi.com/api/facebook/links"
	payload := strings.NewReader("{\"url\":\"https://www.facebook.com/groups/1024490957622648/permalink/1396382447100162/\"}")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("x-rapidapi-key", "51156b7684mshc07d1beddf99fdbp1cbc1bjsncf984ec67efc")
	req.Header.Add("x-rapidapi-host", "facebook17.p.rapidapi.com")
	req.Header.Add("Content-Type", "application/json")

	// Make the HTTP request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return
	}
	defer res.Body.Close()

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print HTTP response status
	fmt.Println(res.Status)

	// Print API response body
	fmt.Println(string(body))

	// Store the API response in memcache for future use, with an expiration time of 1 hour
	mc.Set(&memcache.Item{
		Key:        key,
		Value:      body,
		Expiration: 3600, // 1 hour in seconds
	})

	fmt.Println("Data stored in cache.")
}
