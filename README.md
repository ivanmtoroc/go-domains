# Go Domains
Rest API that allows us to obtain information about a server and to know if the configurations.

## Run
To run app:
```
# Start CockroachDB node
cockroach start --insecure --listen-addr=localhost --background
# Create database and user
cockroach sql --insecure < models/init_db.sql
# Execute server
go run main.go

# In new console tab
cd frontend/
npm run dev
```
