# ml-feature-store-api
Using Go Gin

ml-feature-store-api/
├── main.go               👈 Entry point: start server, register routes
├── handlers/             👈 Handle HTTP requests/responses
│   └── feature.go
├── models/               👈 Data types & business logic (structs, validation, etc.)
│   └── feature.go
├── storage/              👈 Feature storage logic (in-memory, Redis, etc.)
│   └── memory.go
├── go.mod
└── README.md
