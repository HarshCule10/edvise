
---

# 🎓 Backend for Edvise

This is the backend service for the **Edvise** application. It provides APIs for managing assignments, file uploads, and other essential functionalities. The backend is built using **Go** and integrates with **MongoDB** for storing metadata and **Google Cloud Storage (GCS)** for file uploads.

---

## 🚧 **Project Status: WIP**

✅ **Implemented:**
- File upload and metadata storage (MongoDB + GCS).
- Basic API routing and health check.

⚠️ **Pending:**
- Full CRUD functionality (Update, Delete, List).
- Teacher review and feedback APIs.

---

## 📂 **Project Structure**

```
backend/
├── cmd/
│   └── main.go                  # Main entry point
├── config/
│   ├── config.go                # MongoDB and GCS connection
│   └── gcs.go                   # Google Cloud Storage setup
├── internal/
│   ├── handlers/
│   │   └── assignment_handler.go  # Handlers for CRUD operations
│   ├── models/
│   │   └── assignment.go         # Assignment model definition
│   └── services/
│       └── assignment_service.go # Business logic for MongoDB
├── routes/
│   └── routes.go                # API routes definition
└── README.md                    # Project documentation
```

---

## ⚡️ **Prerequisites**

Ensure the following are installed and configured:

- **Go 1.19+** – [Download Go](https://go.dev/dl/)
- **MongoDB** – Running locally or in Docker
- **Google Cloud Storage (GCS)** – Bucket configured for storing uploaded files
- `.env` file with the following variables:
```bash
MONGO_URI=mongodb://localhost:27017
MONGO_DB_NAME=teacher_assistant_db
GCS_BUCKET_NAME=eduvise-bucket
GCS_CREDENTIALS_PATH=config/eduvise-bucket.json
```

---

## 🛠️ **Setup Instructions**

1. **Clone the repository:**
```bash
git clone <repository-url>
cd backend
```

2. **Install dependencies:**
```bash
go mod tidy
```

3. **Set up environment variables:**
- Create a `.env` file in the root folder and add the required variables.
- Ensure the `eduvise-bucket.json` file (Google Cloud credentials) is located at `backend/config/`.

4. **Start MongoDB:**
```bash
sudo systemctl start mongod
```

5. **Run the backend service:**
```bash
go run cmd/main.go
```

---

## 🔥 **API Endpoints**

### 💡 Health Check
- **GET** `/api/health`
  - Response: 
  ```json
  {
    "status": "API is running... 🚀"
  }
  ```

---

### 📚 Assignment Management

- **POST** `/api/assignment/upload`
  - Uploads an assignment file and metadata.
- **GET** `/api/assignment/get` (🚧 Not yet implemented)
  - Retrieves a specific assignment.
- **GET** `/api/assignment/list` (🚧 Not yet implemented)
  - Lists all assignments.
- **PUT** `/api/assignment/update` (🚧 Not yet implemented)
  - Updates an assignment.
- **DELETE** `/api/assignment/delete` (🚧 Not yet implemented)
  - Deletes an assignment.

---

## 📤 **Testing File Upload with `curl`**

To test file uploads, use the following `curl` command:

```bash
curl -X POST -F "file=@/path/to/testfile.pdf" \
  -F "student_id=123" \
  -F "title=Math Assignment" \
  -F "description=Algebra Equations" \
  http://localhost:8080/api/assignment/upload
```

✅ **Successful Response:**
```json
{
  "message": "File uploaded successfully!",
  "file_url": "https://storage.googleapis.com/eduvise-bucket/123-testfile.pdf"
}
```

❗️ **Error Response:**
```json
{
  "error": "Failed to upload file"
}
```

---

## 📦 **Technologies Used**

- 🟦 **Go** – Backend and API development
- 🟢 **MongoDB** – Database for storing assignment metadata
- ☁️ **Google Cloud Storage (GCS)** – File storage for uploaded assignments
- 📚 **gorilla/mux** – Router for handling API requests

---

## 🚀 **Next Steps**
- Implement CRUD for assignments.
- Add teacher review and approval functionality.
- Secure APIs and handle authentication/authorization.

---

✅ All set! You’re ready to run and test the backend for Edvise! 🎉
