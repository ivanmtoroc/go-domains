# Go Domains
Rest API that allows us to obtain information about a server and to know if the configurations.

## Dependecies
```
go get github.com/go-chi/chi
go get github.com/go-chi/render
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/postgres
go get golang.org/x/net/html
```

## Run
```
# Console tab 01:
cockroach start --insecure --listen-addr=localhost --background
cockroach sql --insecure < models/init_db.sql
go run main.go

# Console tab 02:
cd frontend/
npm i
npm run dev
```
