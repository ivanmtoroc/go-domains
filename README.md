# 🔍🌎 Go Domains

Web application to get information about any domain.

The information is getted from [SSL Labs](https://www.ssllabs.com/), web scraping and [Whois](http://manpages.ubuntu.com/manpages/bionic/man1/whois.1.html).

Development using Golang and Vue.js. 🐹💚

## Install dependecies

```bash
# Install Whois bin in Ubuntu
sudo apt install -y whois

# Install Go dependecies
go get -u github.com/lib/pq
go get -u golang.org/x/net/html
go get -u github.com/go-chi/chi
go get -u github.com/go-chi/render
```

Also, [install CockroachDB](https://www.cockroachlabs.com/docs/v20.1/install-cockroachdb.html). 🐞

## Run project

1. Clone GitHub repository:

    ```bash
    git clone https://github.com/ivanmtoroc/go-domains.git $GOPATH/src/github.com/ivanmtoroc/go-domains
    cd $GOPATH/src/github.com/ivanmtoroc/go-domains/
    ```

2. Init CockroachDB cluster:

    ```bash
    chmod 764 init_database.sh
    ./init_database.sh
    ```

3. Start backend server:

    ```bash
    # Console session 1
    go run main.go
    ```

4. Start frontend server:

   ```bash
   # Console session 2
   cd frontend/
   npm i
   npm run dev
   ```
