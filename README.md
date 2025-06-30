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

### Example Feature Payload (POST)

```json
{
  "entity_id": "2025-06-01",
  "name": "temperature",
  "value": 28.5
}
