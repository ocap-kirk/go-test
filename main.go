package main

import (
    "encoding/json"
    "fmt"
    "log"
    "github.com/permitio/permit-golang/pkg/permit"
    "github.com/permitio/permit-golang/pkg/config"
    "github.com/permitio/permit-golang/pkg/enforcement"
)

func main() {
    // Initialize the Permit client
    permitConfig := config.NewConfigBuilder("YOUR_API_KEY_HERE").Build()
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