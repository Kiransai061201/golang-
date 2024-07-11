package main

import (
    "fmt"
    "io"
    "net/http"

    "github.com/bradfitz/gomemcache/memcache"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        url := "https://moviesdatabase.p.rapidapi.com/titles"
        cacheKey := "moviesDatabaseTitles"

        // Initialize the Memcached client
        mc := memcache.New("127.0.0.1:11211")

        // Try to get the data from cache
        item, err := mc.Get(cacheKey)
        if err == memcache.ErrCacheMiss {
            // Cache miss, make the API call
            req, _ := http.NewRequest("GET", url, nil)
            req.Header.Add("x-rapidapi-key", "51156b7684mshc07d1beddf99fdbp1cbc1bjsncf984ec67efc")
            req.Header.Add("x-rapidapi-host", "moviesdatabase.p.rapidapi.com")

            res, err := http.DefaultClient.Do(req)
            if err != nil {
                fmt.Println("Error making API call:", err)
                return c.SendStatus(500)
            }
            defer res.Body.Close()
            body, err := io.ReadAll(res.Body)
            if err != nil {
                fmt.Println("Error reading response body:", err)
                return c.SendStatus(500)
            }

            // Save the response to cache with a 1-hour expiration
            err = mc.Set(&memcache.Item{Key: cacheKey, Value: body, Expiration: 3600})
            if err != nil {
                fmt.Println("Error setting cache:", err)
            }

            // Return the response
            return c.Send(body)
        } else if err != nil {
            // Some other error occurred
            fmt.Println("Error accessing cache:", err)
            return c.SendStatus(500)
        } else {
            // Cache hit, use the cached data
            fmt.Println("Cache hit!")
            return c.Send(item.Value)
        }
    })

    app.Listen(":3000")
}
