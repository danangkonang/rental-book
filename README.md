# Rental Book

## Requirements
- go language
- docker
- docker-compose

## How to running this app

1. clone project

```sh
$ git clone https://github.com/danangkonang/rental-book.git

$ cd rental-book
```

3. build app with docker-compose

```docker
$ docker-compose up --build -d
```

3. runing migration and seeder inside container

    We use [gomig](https://github.com/danangkonang/migration-go-cli) to generate migration and seeder

```bash
# need for fist build

$ docker exec -it golang-api ash

$ ./gomig run migration && ./gomig run seeder

$ exit
```

## Tes

```bash
curl http://localhost:9000
```

## To stop this app
```docker
docker-compose down
```