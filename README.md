# ML Feature Store API 🧠🚀

A lightweight, in-memory feature store API built with Go and Gin. Designed for fast prototyping of machine learning pipelines, online inference, or ML infrastructure experiments.

## 📦 Features

- 🧠 Create and retrieve ML features via REST API
- 🗂 In-memory storage (no external DB required)
- 🔒 Thread-safe with `sync.RWMutex`
- ⚡ Fast and easy to deploy

## 🏗️ API Endpoints

| Method | Endpoint                            | Description                      |
|--------|-------------------------------------|----------------------------------|
| POST   | `/features`                         | Create or update a feature       |
| GET    | `/features/:entity_id/:name`        | Retrieve a specific feature      |

## 📁 Project Structure


| File/Folder     | Description                     |
|-----------------|---------------------------------|
| `main.go`       | Entry point                    |
| `go.mod`        | Go module configuration       |
| `models/`       | Feature model definition       |
| └─ `feature.go` | Struct definitions for Feature |
| `storage/`      | In-memory store with mutex     |
| └─ `store.go`   | Store implementation           |
| `handlers/`     | API logic (set/get feature)    |
| └─ `feature_handler.go` | Handlers for feature APIs |
| `routers/`      | Route setup                   |
| └─ `routers.go` | Router initialization          |


### Example Feature Payload (POST)

```json
{
  "entity_id": "2025-06-01",
  "name": "temperature",
  "value": 28.5
}
```

## 🧠 3 Key Takeaways from Building This Project
1. 🔒 Understanding Lock()/Unlock() in Concurrent Access
By implementing an in-memory store with sync.RWMutex, I learned the importance of using Lock() and Unlock() (or RLock()/RUnlock()) to protect shared data like a Go map.

Without locks:

Concurrent reads/writes can cause data races

Go maps will panic on simultaneous writes

With proper locking:

RLock() allows safe concurrent reads

Lock() serializes writes

defer Unlock() ensures clean, reliable code

This is essential for building safe, concurrent systems.

2. 🚀 Practical Scenarios for Using an In-Memory Feature Store
This kind of API can be used in many useful ways:

🧠 Online inference cache for low-latency model serving

🧪 Rapid prototyping of ML pipelines or REST APIs

🧰 Local development/testing without needing Redis or a database

📦 Temporary store in data pipelines

🌐 Edge ML applications with memory-constrained environments

3. 🧱 Learning Go + Gin Project Structure (MVC Style)
This project taught me how to structure a Go-based API server:

models/ → data schema

handlers/ → API logic

routers/ → route wiring

storage/ → business logic (store/retrieve)

main.go → entry point

Following this modular structure prepares the codebase for future improvements like adding Redis, databases, or Docker.

