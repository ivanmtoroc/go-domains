# Go Domains

> Development using Golang and Vue.js

Web application to get information of a any domain.

## Install dependecies

```bash
# Whois bin
sudo apt install whois # In Ubuntu

# Golang dependecies
go get -u github.com/lib/pq
go get -u github.com/go-chi/chi
go get -u github.com/go-chi/render
go get -u golang.org/x/net/html
```

## Run project

1. Clone Github repository:

    ```bash
    git clone https://github.com/ivanmtoroc/go-domains.git $GOPATH/src/github.com/ivanmtoroc/go-domains
    cd $GOPATH/src/github.com/ivanmtoroc/go-domains/
    ```

2. Start CockroachDB node:

    ```bash
    # console session 01
    cockroach start --insecure --listen-addr=localhost # --background
    ```

3. Start backend server:

    ```bash
    # console session 02
    cockroach sql --insecure < models/config/init_db.sql
    go run main.go
    ```

4. Start frontend server:

    ```bash
    # console session 03
    cd frontend/
    npm i
    npm run dev
    ```
