# logger
Microservice architecture app that contains 2 microservices. One spams the other with pings and the other logs them.

1. Go into the logger folder
```bash
cd logger-service
```

2. Make a .env file here, take this as an example one:
```bash
POSTGRES_USER=admin
POSTGRES_PASSWORD=admin
POSTGRES_DB=db
PORT=5432
```

3. Run the containers
```bash
docker-compose up
```

4. Go into the pinger service
```bash
cd ../pinger-service
```

5. Run that too
```bash
go run main.go
```

6. Go to `localhost:1234/pings`
