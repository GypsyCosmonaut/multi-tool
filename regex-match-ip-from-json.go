package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"sort"
	"time"
)

// ---- Helpers: Random IP Generators ----

// generatePrivateIP returns a random RFC1918 private IPv4 address.
func generatePrivateIP() string {
	r := rand.Intn(3)
	switch r {
	case 0: // 10.0.0.0/8
		return fmt.Sprintf("10.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256))
	case 1: // 172.16.0.0/12
		return fmt.Sprintf("172.%d.%d.%d", 16+rand.Intn(16), rand.Intn(256), rand.Intn(256))
	default: // 192.168.0.0/16
		return fmt.Sprintf("192.168.%d.%d", rand.Intn(256), rand.Intn(256))
	}
}

// generatePublicIP generates a random public IPv4 address in safe, non-private ranges.
func generatePublicIP() string {
	for {
		ip := fmt.Sprintf("%d.%d.%d.%d",
			rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256),
		)

		// Skip private ranges
		if isPrivate(ip) {
			continue
		}

		return ip
	}
}

// Check if an IP is private.
func isPrivate(ip string) bool {
	privateCIDRs := []string{
		"10.", "192.168.",
	}

	for _, p := range privateCIDRs {
		if ip[:len(p)] == p {
			return true
		}
	}

	// Check 172.16.0.0 – 172.31.255.255
	if ip[:3] == "172" {
		var a, b, c, d int
		fmt.Sscanf(ip, "%d.%d.%d.%d", &a, &b, &c, &d)
		if a == 172 && b >= 16 && b <= 31 {
			return true
		}
	}

	return false
}

func main() {
	rand.Seed(time.Now().UnixNano())

	filename := "ips.json"

	// ---- 1. Create JSON structure ----
	type IPData struct {
		PrivateIPs []string `json:"private_ips"`
		PublicIPs  []string `json:"public_ips"`
	}

	data := IPData{
		PrivateIPs: make([]string, 5),
		PublicIPs:  make([]string, 5),
	}

	for i := 0; i < 5; i++ {
		data.PrivateIPs[i] = generatePrivateIP()
		data.PublicIPs[i] = generatePublicIP()
	}

	// Marshal JSON
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "json marshal error: %v\n", err)
		os.Exit(1)
	}

	// ---- 2. Write to file ----
	if err := os.WriteFile(filename, jsonBytes, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "file write error: %v\n", err)
		os.Exit(1)
	}

	// ---- 3. Read file back ----
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file read error: %v\n", err)
		os.Exit(1)
	}

	// ---- 4. Extract all IPs via RegEx ----
	ipRegex := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b`)
	ips := ipRegex.FindAllString(string(content), -1)

	// Remove duplicates
	unique := make(map[string]struct{})
	for _, ip := range ips {
		unique[ip] = struct{}{}
	}

	// Convert to list and sort
	var sortedIPs []string
	for ip := range unique {
		sortedIPs = append(sortedIPs, ip)
	}
	sort.Strings(sortedIPs)

	// Print extracted unique IPs
	for _, ip := range sortedIPs {
		fmt.Println(ip)
	}

	// ---- 5. Delete the file ----
	if err := os.Remove(filename); err != nil {
		fmt.Fprintf(os.Stderr, "file delete error: %v\n", err)
		os.Exit(1)
	}
}

