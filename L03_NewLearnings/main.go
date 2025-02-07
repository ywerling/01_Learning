package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Define the URL
	url := "https://stackoverflow.com/questions/37351022/golang-mixed-assignment-and-declaration#37358972"

	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close() // Ensure response body is closed

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print response
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}
