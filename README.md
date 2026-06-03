# Fraud Alert Monitor

Fraud Alert Monitor is a backend service written in Go for monitoring voice call traffic and detecting potentially suspicious calls.

## Features

* Store call information in PostgreSQL
* REST API for managing calls
* Fraud detection engine
* Automatic fraud alert generation
* Alert monitoring API
* Statistics API

## Tech Stack

* Go
* PostgreSQL
* REST API
* Git
* GitHub

## Architecture

HTTP API

↓

Handlers

↓

Fraud Detection Service

↓

Repository Layer

↓

PostgreSQL

## API Endpoints

### Health Check

GET /health

Response:

```json
{
  "status": "ok"
}
```

### Create Call

POST /calls

Request:

```json
{
  "phone_from": "77001234567",
  "phone_to": "49123456789",
  "country": "Germany",
  "duration": 120,
  "status": "SUCCESS"
}
```

### Get Calls

GET /calls

Returns all stored calls.

### Get Alerts

GET /alerts

Returns all fraud alerts.

### Get Statistics

GET /stats

Example response:

```json
{
  "total_calls": 15,
  "total_alerts": 4,
  "failed_calls": 2
}
```

## Fraud Detection Rules

The system automatically generates alerts for:

* Calls shorter than 10 seconds
* Calls with FAILED status

## Database

Main tables:

### calls

Stores call records.

### fraud_alerts

Stores generated fraud alerts.

## Running the Project

1. Install PostgreSQL
2. Create database:

```sql
CREATE DATABASE fraud_monitor;
```

3. Configure connection settings in:

```text
database/db.go
```

4. Start the application:

```bash
go run cmd/main.go
```

5. Open:

```text
http://localhost:8080/health
```

## Future Improvements

* Docker support
* Dependency Injection
* CI/CD
* Advanced fraud detection rules
* Dashboard UI
* Authentication and authorization

## Author

Dias Zhorkhabaev
