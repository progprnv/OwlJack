package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// DisplayBanner shows the OWLJACK banner
func DisplayBanner() {
	banner := `
       ________        
     /          \      
    /            \     
   /   OWLJACK    \    
  /________________\   
    \   /\  /\   /     
     ) ( •  • ) (      
    \    ----    /     
     \    ^^    /       
      \________/       
   OWLJACK - Your Clickjacking Tool
`
	fmt.Println(banner)
}

// TestClickjacking checks for X-Frame-Options in the response headers
func TestClickjacking(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[ERROR] Failed to fetch URL: %s\n", url)
		return
	}
	defer resp.Body.Close()

	xFrameOptions := resp.Header.Get("X-Frame-Options")
	if xFrameOptions == "" {
		fmt.Printf("[VULNERABLE] %s: No X-Frame-Options header found.\n", url)
	} else {
		fmt.Printf("[SECURE] %s: X-Frame-Options = %s\n", url, xFrameOptions)
	}
}

func main() {
	// Display the banner
	DisplayBanner()

	// Ask user for the file containing the list of URLs
	fmt.Print("Enter the path to the URL list file: ")
	var filePath string
	fmt.Scanln(&filePath)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("[ERROR] Unable to open file: %s\n", err)
		return
	}
	defer file.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	fmt.Println("\nTesting URLs for Clickjacking vulnerability...")
	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url == "" {
			continue
		}
		TestClickjacking(url)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("[ERROR] Reading file: %s\n", err)
	}
}
