Customer Review Email System:
=============================
A simple test task implementation from iTechArt.
Technical description in [google doc](https://docs.google.com/document/d/1RVXLkRCXvY1LuEW1ZurSFfEt0fC4A63Zgu6qXRkG1kE/edit).

Requirements:
-------------

[Docker](https://www.docker.com/)

Local development:
-----------------

1. Copy all .env.example files into .env
```bash
cd api/.docker.golang
cp .env.example .env
 
cd grpc-server/.docker/golang 
cp .env.example .env
 
cd grpc-server/.docker/mysql
cp .env.example .env

cd tone-analyzer/.docker/golang
cp .env.example .env

cd tone-analyzer/.docker/rabbitmq
cp .env.example .env
```

2. Build images:
```bash
docker-compose build
```

3. Run containers:
```bash
docker-compose up -d
```

4. Ensure all containers were run successfully:
```bash
docker-compose ps
```

Go to [http://localhost:3000](http://localhost:3000) to see how everything works.


