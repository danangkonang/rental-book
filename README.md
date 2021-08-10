# Student Library Scholl

## Requirements
- go language
- docker
- docker compose

## How to running this app

```docker
docker-compose up --build -d
```

## Running migration and seeder

We use [gomig](https://github.com/danangkonang/migration-go-cli) to generate migration and seeder

```bash
docker exec -it golang-api ash

./gomig run migration && ./gomig run seeder

exit
```

## Tes

```bash
curl http://localhost:9000
```

## To this stop
```docker
docker-compose down
```