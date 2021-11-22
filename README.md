# Blog System

## Setup
Execute the `start.sh` bash script at project to start the application server.

## Information
1. Setting custom server port:

By default, the server is started on port 8080 on the host system.
To change this, simply update the application service in `docker-compose.yml` file and then rerun the `start.sh` script.
```
article-blog:
    depends_on:
      - mysqldb
    build:
      context: .
      dockerfile: Dockerfile
    image: article-blog:1.0
    restart: always
    ports:
      - target: "8080"
#       Mention port to be exposed on the host system
        published: "8080"      <- put wanted server port here
    environment:
#     Environment variable to the mysql service::provides SQL DB
      DB_CONFIG_HOST: "mysqldb"
```

Since the entire server is started in a Docker container, the only pre-requisite is the installation of docker on the host system along with docker-compose, which usually is installed with Docker.
This makes the application **independent of the host OS**.

2. API Collection:

The API collection, `API-Collection.json`, is provided at the project root. One may import this to get the list of API endpoints with default configurations to run on localhost, once the application is ready.