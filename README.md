# System Requirement
```golang version >= 1.23.0```

# How to Run
1. git clone this repo
2. copy .env.example and rename it to .env
3. go mod vendor
4. go run .

# Example CURL
```curl --location 'localhost:8000/suggestions?q=Londo&latitude=43.70011&longitude=-79.4163'```

# Example CURL Production
```https://lacak-suggestions-api-production.up.railway.app/suggestions?q=London&latitude=42.98339&longitude=-81.23304```
