package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"time"
)

// Post represents each JSON object returned by the API.
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	// Setup an HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Context with timeout for safety
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create HTTP request
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating request: %v\n", err)
		os.Exit(1)
	}

	// Do the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Check HTTP status
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "unexpected status code: %d\n", resp.StatusCode)
		os.Exit(1)
	}

	// Decode JSON response
	var posts []Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		fmt.Fprintf(os.Stderr, "error decoding JSON: %v\n", err)
		os.Exit(1)
	}

	// Collect unique UserIDs using a map
	unique := make(map[int]struct{})
	for _, p := range posts {
		unique[p.UserID] = struct{}{}
	}

	var ids []int

	// unique UserIDs
	for id := range unique {
		ids = append(ids, id)
	}

	sort.Ints(ids)

	// Print sorted unique UserIDs
	for _, id := range ids {
		fmt.Println(id)
	}


	// Print JSON back to stdout (pretty-printed)
	pretty, err := json.MarshalIndent(posts, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(pretty))
}

