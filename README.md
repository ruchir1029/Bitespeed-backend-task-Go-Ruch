# Bitespeed-backend-task-Go-Ruch
This application helps in identifying and reconciling user identities based on phone numbers and emails. It uses BoltDB, a pure Go key-value store, for data storage.

Hi, I'm Ruchir, a 2024 CSE graduate from New Horizon College of Engineering, aspiring to be a full-stack developer. Welcome to the BiteSpeed Identity Reconciliation project!



## Requirements

- Go 1.22.4 or later

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/ruchir1029/Bitespeed-backend-task-Go-Ruch.git
    cd Bitespeed-backend-task-Go-Ruch
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

## Usage

1. Run the application:
    ```sh
    go run main.go db.go handler.go models.go utils.go
    ```

2. The server will start on port 8080.

## API Endpoints

- **POST /identify**
    - Request:
      ```json
      {
          "phone_number": "1234567890",
          "email": "example@example.com"
      }
      ```
    - Response:
      ```json
      {
          "contact": {
              "primary_contact_id": 1,
              "emails": ["example@example.com"],
              "phone_numbers": ["1234567890"],
              "secondary_contact_ids": []
          }
      }
      ```

Thank You !
