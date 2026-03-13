# contact-be

Go bootstrap for a Contact microservice.

## Prerequisites

- Go 1.22+

## Project Structure

```
.
├── cmd/
│   └── contact-service/
│       └── main.go
├── internal/
│   └── contact/
│       └── service.go
├── go.mod
└── README.md
```

## Run Locally

```bash
go run ./cmd/contact-service
```

Server default address: `:8080`

Override with environment variable:

```bash
PORT=9090 go run ./cmd/contact-service
```

## Available Endpoints

- `GET /healthz` returns service health.
- `GET /contacts` lists all contacts from in-memory store.
- `POST /contacts` creates a contact in-memory.

### Example Create Request

```bash
curl -X POST http://localhost:8080/contacts \
	-H "Content-Type: application/json" \
	-d '{"name":"Ada Lovelace","email":"ada@example.com","phone":"+1-555-0001"}'
```

## Notes

- This is an initial bootstrap and currently uses in-memory storage.
- Next iteration can add persistence, validation, auth, and OpenAPI docs.