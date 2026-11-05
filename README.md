# Building-job-Queue

A job queue system written in Go and built step by step. Clients submit jobs over HTTP. The plan is for background workers to pick those jobs up and process them, retrying any that fail.

> **Status:** early stage. The HTTP API accepts and validates jobs, but they are not stored or processed yet.

## Tech stack

- Go 1.26
- Standard library only (`net/http`, `encoding/json`), with no external dependencies

## Project structure

```
.
├── main.go        # HTTP server and route handlers
├── jobs/
│   └── jobs.go    # Job and response types
├── go.mod         # Module: jobqueue
└── README.md
```

## Getting started

### Prerequisites

- [Go](https://go.dev/dl/) 1.26 or newer

### Run the server

```bash
go run .
```

The server listens on `http://localhost:8080`.

## API

### `POST /createjob`

Submits a new job.

**Request body**

| Field         | Type   | Description                                   |
|---------------|--------|-----------------------------------------------|
| `type`        | string | Kind of job, for example `send_email`         |
| `payload`     | JSON   | Any JSON data the job needs                   |
| `max_retries` | int    | How many times to retry the job if it fails   |

**Example**

```bash
curl -i -X POST localhost:8080/createjob \
  -H "Content-Type: application/json" \
  -d '{"type":"send_email","payload":{"to":"user@example.com"},"max_retries":3}'
```

**Responses**

| Status                     | When                                               | Body                   |
|----------------------------|----------------------------------------------------|------------------------|
| `202 Accepted`             | The job was parsed successfully                    | `{"status":"queued"}`  |
| `400 Bad Request`          | The body is invalid JSON or larger than 1 MB       | `Unable to parse json` |
| `405 Method Not Allowed`   | Any method other than `POST`                       | n/a                    |

## Roadmap

- [x] Basic HTTP server
- [x] `POST /createjob` endpoint
- [ ] In-memory queue (Go channel)
- [ ] Worker pool to process jobs
- [ ] Retry logic using `max_retries`
- [ ] Job IDs and a `GET /jobs/{id}` status endpoint
- [ ] Persistent storage (for example Redis or Postgres)
- [ ] Graceful shutdown

## Progress log

### Day 1

- Set up the Go module (`jobqueue`) and the basic package layout
- Added the `jobs` package with the `Jobs` and `JobResponse` types
- Built the HTTP server with a `POST /createjob` route that parses and validates incoming jobs
