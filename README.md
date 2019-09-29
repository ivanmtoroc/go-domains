# 🔍🌎 Go Domains

Web application to get information about any domain.

The information is getted from [SSL Labs](https://www.ssllabs.com/), web scraping and [Whois](http://manpages.ubuntu.com/manpages/bionic/man1/whois.1.html).

Development using Golang and Vue.js.

## Dependecies

```bash
# install Whois bin in Ubuntu
sudo apt install whois

# install Go dependecies
go get -u github.com/lib/pq
go get -u github.com/go-chi/chi
go get -u github.com/go-chi/render
go get -u golang.org/x/net/html
```

## Run

1. Clone Github repository:

    ```bash
    git clone https://github.com/ivanmtoroc/go-domains.git $GOPATH/src/github.com/ivanmtoroc/go-domains
    cd $GOPATH/src/github.com/ivanmtoroc/go-domains/
    ```

2. Start CockroachDB node:

    ```bash
    # console session 0
    cockroach start --insecure --listen-addr=localhost # --background
    ```

3. Start backend server:

    ```bash
    # console session 1
    cockroach sql --insecure < models/config/init_db.sql
    go run main.go
    ```

4. Start frontend server:

    ```bash
    # console session 2
    cd frontend/
    npm i
    npm run dev
    ```
