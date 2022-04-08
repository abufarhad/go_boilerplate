# Golang Backend Boilerplate

# Run Locally
Install dependencies

```go
go mod vendor
```

### Data Migration using `gooes`
> To create initial table go to this directory`go_boilerplate/infra/conn/migration`  

> Run this command
> `goose create table_name go  `
### Seedding database

```go
go run main.go seed
```

### Use the below command to truncate then seed database  

```go
go run main.go seed --truncate=true

or

go run main.go seed -t=true
```

## Start the server Locally

```go
go run main.go serve
```

# Start the server using Docker

```go
make development
```


