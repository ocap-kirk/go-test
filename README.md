# Permit.io Go SDK Example

This is a simple Go application that demonstrates how to use the Permit.io Go SDK to check user permissions.

## Features

- Initialize the Permit.io client
- Create a user
- Check user permissions
- Output permissions in JSON format

## Prerequisites

- Go 1.23 or later
- A Permit.io account and API key

## Setup

1. Clone the repository:
```bash
git clone https://github.com/yourusername/go-test.git
cd go-test
```

2. Install dependencies:
```bash
go mod download
```

3. Set up your Permit.io API key:
   - Replace `YOUR_API_KEY_HERE` in `main.go` with your actual Permit.io API key
   - Make sure your Permit.io PDP (Policy Decision Point) is running and accessible

## Running the Application

```bash
go run main.go
```

## Configuration

The application uses the following configuration:
- API Key: Set in `main.go`
- PDP URL: Configured in the Permit.io dashboard
- User ID: Set in `main.go`
- Resource ID: Set in `main.go`

## License

MIT 