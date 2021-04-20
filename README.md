# go-graphql-server

## Setup project

https://gqlgen.com/getting-started/

```shell
go mod init github.com/[username]/go-graphql-server
go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init
```

## Start of GraphQL Server

```shell
go run server.go
```

## Add magic comment to resolver.go

```golang
//go:generate go run github.com/99designs/gqlgen

package graph
```

```shell
go generate ./...
```
