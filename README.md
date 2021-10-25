# Dockerized React+Golang api skeleton

Golang API folder `_source/api`

React APP folder `web`

# Setup

### General

+ `cp .env.dist .env`
+ Fill up .env

### API

Api has base examples of handlers

`_source/api/main.go:20`

```
/api/hello_world
/api/hello_user/{user}
/api/error
```

### React App

`create-react-app ./web`

# Update

#### Api

+ Change code
+ `docker-compose stop api`
+ `docker-compose up -d --build`

Sometimes it could be a glitch with accessibility to api. In that case:

```
docker-compose stop api
docker-compose stop revproxy
docker-compose up -d
```

# Run

`docker-compose up -d`

Web Application will be available on `localhost:8080`

React app has access to Api via `localhost:8080/api`