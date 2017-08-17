# go-todo-example

## Development
To run locally, you will need to set the following environment variables:
```bash
DB_USERNAME
DB_PASSWORD
DB_HOST
DB_PORT
```
Run the server:
```bash
cd cmd/server
go build
./server
```

## Testing
Set up your database:
```bash
docker run --name test-mysql -e MYSQL_ROOT_PASSWORD=my-secret -e MYSQL_DATABASE=todo -d -p 3306:3306 -v schema.sql:/docker-entrypoint-initdb.d/schema.sql mysql
```
Test mysql package:
```bash
export DB_USERNAME=root
export DB_PASSWORD=my-secret
export DB_HOST=127.0.0.1
export DB_PORT=3306

cd mysql
go test
```

## API
This API provides the following endpoints:
```
POST    /items
DELETE  /items/{itemID}
GET     /items/{itemID}
PUT     /items/{itemID}
```
