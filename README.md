# AUTH MOVIE APP

## Prerequisite

- Go 1.14
- migrate/migrate
- docker
- postgres

## How to install

- Go 1.14

  You can install go 1.14 through this following link https://golang.org/dl/

- migrate/migrate 
  ```
  docker pull migrate/migrate
  ```
  


## How to add migration file

### Example Script
```bash
docker run --rm -v "$(pwd)"/db/migrations:/migrations migrate/migrate create -ext sql -dir /migrations create_users_table
```

## How to migrate database

```
docker run --rm -v "$(pwd)"/db/migrations:/migrations --network host migrate/migrate -database 'postgres://postgres:password@localhost:5432/auth?sslmode=disable' -path migrations up
``` 

## How to create mock of interfaces
You have to install mockgen
  ```bash
  mockgen -source internal/domain/user.go -destination internal/domain/mock/user.go
  ```

Don't forget to include your go/bin to $PATH
```
export PATH="$PATH:$GOPATH/bin"
```