Customer Review Email System:
=============================
A simple test task implementation from iTechArt.
Technical description in [google doc](https://docs.google.com/document/d/1RVXLkRCXvY1LuEW1ZurSFfEt0fC4A63Zgu6qXRkG1kE/edit).

Requirements:
-------------

[Docker](https://www.docker.com/)  
[Dep](https://golang.github.io/dep/docs/installation.html) - Golang package manager

Usage:
------

UI is ready on [http://35.192.27.127/](http://35.192.27.127/) for some testing.

Local development:
-----------------

1. Run installation
```bash
make install
```

2. Run containers:
```bash
docker-compose up -d
```

3. Ensure all containers were run successfully:
```bash
docker-compose ps
```

4. Run migrations:
```bash
docker-compose exec cr-golang-grpc-server make migrate-up
```

Go to [http://localhost:3000](http://localhost:3000) to see how everything works.


