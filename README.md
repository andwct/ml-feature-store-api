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



