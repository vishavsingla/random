# Image Processing Service

A high-performance microservice designed to process large volumes of store-based image data using GPU simulation. The service efficiently handles thousands of images, calculating perimeters and managing processing states through a robust job system.

## 🚀 Features

- **Bulk Image Processing**: Handle thousands of images in a single job
- **Store-based Validation**: Validate store IDs against master data
- **Job Management**: Track processing status and handle errors gracefully
- **GPU Simulation**: Realistic processing time simulation for testing
- **RESTful API**: Simple and intuitive API endpoints
- **Docker Support**: Easy deployment with containerization

## 📋 Prerequisites

- Go 1.20 or higher
- Docker (optional)

## 🛠️ Installation

### Local Setup

1. Clone the repository:
```bash
git clone https://github.com/vishavsingla/random.git
cd random
```

2. Install dependencies:
```bash
go mod tidy
```

3. Run the application:
```bash
go run main.go
```

### Docker Setup

1. Build the image:
```bash
docker build -t myproject .
```

2. Run the container:
```bash
docker run -p 8080:8080 myproject
```

## 🔌 API Endpoints

### Submit Job
```http
POST /api/submit
```

**Request Body:**
```json
{
   "count": 2,
   "visits": [
      {
         "store_id": "S00339218",
         "image_url": [
            "https://www.gstatic.com/webp/gallery/2.jpg",
            "https://www.gstatic.com/webp/gallery/3.jpg"
         ],
         "visit_time": "2024-11-15T10:00:00Z"
      },
      {
         "store_id": "S01408764",
         "image_url": [
            "https://www.gstatic.com/webp/gallery/3.jpg"
         ],
         "visit_time": "2024-11-15T10:10:00Z"
      }
   ]
}
```

**Success Response:**
```json
{
   "job_id": 123
}
```

### Check Job Status
```http
GET /api/status?jobid=123
```

**Success Response:**
```json
{
   "status": "completed",
   "job_id": 123
}
```

## 🧪 Testing

Run the test suite:
```bash
go test ./main_test.go
```

## 🏗️ Architecture

The service follows a microservices architecture with these key components:

- **API Layer**: Handles HTTP requests and routing
- **Store Management**: Validates store data against master records
- **Job Processing**: Manages worker pools and image processing
- **Status Tracking**: Monitors job progress and maintains state

## 📚 Tech Stack

- **Backend**: Go 1.20+
- **Libraries**:
  - `net/http`: HTTP server implementation
  - `encoding/json`: JSON processing
  - `os`: System operations
  - `csv`: CSV file processing
  - `time`: Processing simulation

## 🚀 Future Improvements

- Worker pool implementation for parallel processing
- Database integration for result persistence
- Authentication and authorization
- Enhanced error handling and logging
- Message queue integration (RabbitMQ/Kafka)
- File upload support
- Rate limiting and back-pressure mechanisms

## 🤝 Contributing

Contributions are welcome! Please feel free to submit pull requests.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
