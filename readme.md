# Fuel Tracker

## Run

```
go run main.go
```

## Test

```
curl -X POST http://localhost:8080/api/v1/fuel-logs \
  -H "Content-Type: application/json" \
  -d '{
    "date": "2025-10-01T00:00:00Z",
    "price_per_liter": 12100,
    "total_paid": 30000,
    "liters_filled": 2.48,
    "km_start": 17544,
    "km_end": 17639,
    "location": "3P Studio Alam",
    "notes": ""
  }'

curl http://localhost:8080/api/v1/fuel-logs
curl http://localhost:8080/api/v1/fuel-logs/1
curl -X DELETE http://localhost:8080/api/v1/fuel-logs/1
```