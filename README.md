# Running

## run up docker postgres
docker run --rm -d --name atlas-psql-demo -p 5432:5432 -e POSTGRES_PASSWORD=123456 postgres

## migrate
go run ./migrate

## start server
go run ./server
