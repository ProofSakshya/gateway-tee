# gateway-tee

The `gateway-tee` project is a Trusted Execution Environment (TEE) gateway written in Go. It provides a secure API service built using the [`gin-gonic/gin`](https://gin-gonic.com/) web framework and is intended to interface with AWS Nitro Enclaves.

## Prerequisites
- **Go**: Version 1.21 or later.

## Initiation & Setup

1. **Clone the repository**:
   ```bash
   git clone <repository_url>
   cd gateway-tee
   ```

2. **Download dependencies**:
   Installs Go packages, specifically `gin` (and AWS Nitro Enclaves SDK once enabled).
   ```bash
   go mod download
   ```

3. **Run the Gateway API**:
   _Note: Provide specific entry location, usually within the `cmd/` directory as typical with standard layout Go projects._
   ```bash
   go run ./cmd/main.go
   # OR
   # go run main.go 
   ```

## Structure
- `/cmd/`: Main entry point(s) for applications.
- `/internal/`: Private application-specific logic.

## Roadmap & Features (MVP)
- TEE Integration: Will utilize `aws-nitro-enclaves-sdk-go` (currently stubbed out for MVP).
- API Layer: REST interface via Gin.
