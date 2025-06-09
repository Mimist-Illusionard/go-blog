# Go-blog description
A simple blog written on go

This project consists of three main components:

frontend — the React-based client

backend — the Go-based API server

postgres — a PostgreSQL database

Make sure you have the following installed:
- Docker
- Docker Compose

## 📁 Project Structure
```
├── backend/
│   ├── cmd/
|   ├── config/
│   ├── internal/
│   ├── utils/
|   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
├── frontend/
│   ├── src/
│   ├── public/
│   ├── Dockerfile/
│   └── package.json
├── docker-compose.yml
└── README.md
```
## ⚙️ How to Run
1. Clone the repository:
```
git clone https://github.com/Mimist-Illusionard/go-blog.git
cd go-blog
```
2. Start the project:
```
docker-compose up --build
```
This command will build and run all three containers:

go-blog-frontend will be available at: http://localhost:3000

go-blog-backend will be accessible at: http://localhost:9090

go-blog-postgres will run PostgreSQL on port 5432

3. Stopping the project:
```
docker-compose down
```

## Postman collection
https://orange-crescent-5490765.postman.co/workspace/Test.Ayo's-Workspace~c4089bd2-231e-4d63-8d8b-3d0cded6c412/collection/44290956-5a1b022b-0ad0-4823-a6b4-2fb8bea928b5?action=share&creator=44290956
