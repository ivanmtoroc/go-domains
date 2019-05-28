# Go Domains
Rest API that allows us to obtain information about a server and to know if the configurations.

## Install dependecies
```
# Whois
sudo apt install whois # In Ubuntu

# Go dependecies
go get -u github.com/lib/pq
go get -u github.com/go-chi/chi
go get -u github.com/go-chi/render
go get -u golang.org/x/net/html
```

## Run project
Clone repo
```
git clone https://github.com/ivanmtoroc/go-domains.git $GOPATH/src/go-domains
cd $GOPATH/src/go-domains/
```

Start CockroachDB node
```
# console session 01
cockroach start --insecure --listen-addr=localhost # --background
```

Start backend server
```
# console session 02
cockroach sql --insecure < models/init_db.sql
go run main.go
```

Start frontend server
```
# console session 03
cd frontend/
npm i
npm run dev
```
