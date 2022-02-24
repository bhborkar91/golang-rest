# Running locally

- (optional) If you don't have a mongodb instance, you can start one with

```
docker-compose -f docker-compose-db.yml up -d
```

- If you have nodemon installed, you can start the app server with live reload using

```
nodemon
```

- If you don't have nodemon installed, start the server using

```
export MONGO_URL="mongodb://localhost:37017/golang-rest"
go run .

```

# Running using docker compose

TBD

# Running using skaffold

TBD