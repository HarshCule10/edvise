# Backend for Edvise

This is the backend service for the Edvise application. It provides APIs for managing assignments, file uploads, and other functionalities. The backend is built using Go and integrates with MongoDB and Google Cloud Storage.

##THIS IS A WIP
Currently CRUD is not fully implemented. Only Create and Uplaod works.
Other functions will be functional soon.
Backend can be tested using CURL
Example :
   ```bash
   curl -X POST -F "file=@your/file/path/testfile.pdf" \
-F "student_id=123" \
-F "title=Math Assignment" \
-F "description=Algebra Equations" \
http://localhost:8080/api/assignment/upload
   ```


## Project Structure
backend/
├── cmd/
│   └── main.go         # Main entry point
├── config/
│   ├── config.go       # MongoDB and GCS connection
│   └── gcs.go          # Google Cloud Storage setup
├── internal/
│   ├── handlers/
│   │   └── assignment_handler.go  # Handlers for CRUD operations
│   ├── models/
│   │   └── assignment.go         # Assignment model definition
│   └── services/
│       └── assignment_service.go # Business logic for MongoDB
├── routes/
│   └── routes.go       # API routes definition
└── README.md           # Project documentation


## Prerequisites

- Go 1.19 or later
- MongoDB
- Google Cloud Storage bucket
- `.env` file with the following variables:
  ```env
  MONGO_URI=mongodb://localhost:27017
  MONGO_DB_NAME=teacher_assistant_db
  GCS_BUCKET_NAME=eduvise-bucket
  GCS_CREDENTIALS_PATH=config/eduvise-bucket.json
  ```

## Setup

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd backend
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up the `.env` file with the required environment variables.

4. Start MongoDB and ensure it is running.

5. Run the backend service:
   ```bash
   go run cmd/main.go
   ```

## API Endpoints

### Health Check
- **GET** `/api/health`
  - Response: `API is running... 🚀`

### Assignment Management
- **POST** `/api/assignment/upload`
  - Upload an assignment file and metadata.
- **GET** `/api/assignment/get`
  - Retrieve a specific assignment (to be implemented).
- **GET** `/api/assignment/list`
  - List all assignments (to be implemented).
- **PUT** `/api/assignment/update`
  - Update an assignment (to be implemented).
- **DELETE** `/api/assignment/delete`
  - Delete an assignment (to be implemented).

## Technologies Used

- **Go**: Backend programming language.
- **MongoDB**: Database for storing assignment metadata.
- **Google Cloud Storage**: File storage for assignment uploads.


