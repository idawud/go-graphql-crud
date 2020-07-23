# go-graphql-crud
simple crud graphql server in go

## Build & Run Dockerfile
Build:  `docker build -t idawud/gql-server .`

Run:  `docker run --rm idawud/gql-server`

### Run with curl
Run: `curl -XPOST -v http://localhost:8080/graphql -H 'Content-Type: application/json' -d 'query { movies{ id, title,minutes }}'`