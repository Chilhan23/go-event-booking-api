package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const baseURL = "http://localhost:8080"

type RegisterResponse struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

type LoginResponse struct {
	Data struct {
		Token string `json:"token"`
	} `json:"data"`
}

type EventResponse struct {
	Data struct {
		ID string `json:"id"`
	} `json:"data"`
}

func main() {
	fmt.Println("Starting concurrent booking simulation...")

	// 1. Create an admin user to set up the event
	adminToken := registerAndLogin("admin_organizer", "admin@example.com")

	// 2. Create event with quota of 5
	eventID := createEvent(adminToken, 5)
	fmt.Printf("Event created with quota: 5 (ID: %s)\n", eventID)

	// 3. Register 20 distinct users
	totalUsers := 20
	userTokens := make([]string, totalUsers)
	for i := 0; i < totalUsers; i++ {
		u := fmt.Sprintf("user_%d_%d", time.Now().UnixNano(), i)
		userTokens[i] = registerAndLogin(u, u+"@example.com")
	}
	fmt.Printf("%d participant users prepared.\n", totalUsers)

	// 4. Setup synchronization barriers
	var wg sync.WaitGroup
	startSignal := make(chan struct{})
	resultChan := make(chan int, totalUsers)

	fmt.Println("\nFiring 20 concurrent booking requests simultaneously...")

	for i := 0; i < totalUsers; i++ {
		wg.Add(1)
		go func(token string) {
			defer wg.Done()

			<-startSignal // Wait for trigger signal

			req, _ := http.NewRequest("POST", fmt.Sprintf("%s/events/%s/book", baseURL, eventID), nil)
			req.Header.Set("Authorization", "Bearer "+token)

			client := &http.Client{Timeout: 10 * time.Second}
			resp, err := client.Do(req)
			if err != nil {
				resultChan <- 500
				return
			}
			defer resp.Body.Close()

			resultChan <- resp.StatusCode
		}(userTokens[i])
	}

	// 5. Trigger all goroutines at once
	close(startSignal)

	wg.Wait()
	close(resultChan)

	// 6. Aggregate results
	successCount := 0
	failedCount := 0

	for status := range resultChan {
		if status == http.StatusCreated {
			successCount++
		} else {
			failedCount++
		}
	}

	fmt.Println("\nTest Result : ")
	fmt.Printf("Total Requests: %d\n", totalUsers)
	fmt.Printf("Successful Bookings (201 Created): %d\n", successCount)
	fmt.Printf("Rejected Requests (Quota Exhausted): %d\n", failedCount)
	fmt.Println("")

	if successCount == 5 {
		fmt.Println("SUCCESS: PostgreSQL Row-Level Lock successfully prevented race condition.")
	} else {
		fmt.Println("FAILURE: Overbooking or data anomaly detected.")
	}
}

// Helper: Register & Login
func registerAndLogin(username, email string) string {
	regBody, _ := json.Marshal(map[string]string{
		"username": username, "email": email, "password": "password123",
	})
	http.Post(baseURL+"/auth/register", "application/json", bytes.NewBuffer(regBody))

	loginBody, _ := json.Marshal(map[string]string{
		"username": username, "password": "password123",
	})
	resp, _ := http.Post(baseURL+"/auth/login", "application/json", bytes.NewBuffer(loginBody))
	var res LoginResponse
	json.NewDecoder(resp.Body).Decode(&res)
	return res.Data.Token
}

// Helper: Create Event
func createEvent(token string, quota int) string {
	body, _ := json.Marshal(map[string]interface{}{
		"title":       "Exclusive Concert 2026",
		"description": "Super high-demand event",
		"location":    "Jakarta Stadium",
		"starts_at":   time.Now().Add(24 * time.Hour),
		"ends_at":     time.Now().Add(48 * time.Hour),
		"quota":       quota,
	})

	req, _ := http.NewRequest("POST", baseURL+"/events/create", bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, _ := client.Do(req)

	var res EventResponse
	json.NewDecoder(resp.Body).Decode(&res)
	return res.Data.ID
}
