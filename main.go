package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv"
    "github.com/permitio/permit-golang/pkg/permit"
    "github.com/permitio/permit-golang/pkg/config"
    "github.com/permitio/permit-golang/pkg/enforcement"
)

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Get API key from environment variable
    apiKey := os.Getenv("PERMIT_API_KEY")
    if apiKey == "" {
        log.Fatal("PERMIT_API_KEY environment variable is not set")
    }

    // Initialize the Permit client
    permitConfig := config.NewConfigBuilder(apiKey).Build()
    Permit := permit.New(permitConfig)

    // Create a user
    user := enforcement.UserBuilder("abc123").Build()

    // Get user permissions
    permissions, err := Permit.GetUserPermissions(user, "default")
    if err != nil {
        log.Fatalf("Failed to get user permissions: %v", err)
    }

    // Convert permissions to JSON
    jsonData, err := json.MarshalIndent(permissions, "", "  ")
    if err != nil {
        log.Fatalf("Failed to marshal permissions to JSON: %v", err)
    }

    // Print the permissions in JSON format
    fmt.Println(string(jsonData))
} 