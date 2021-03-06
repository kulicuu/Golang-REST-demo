


#### A Little Preface:

My pro experience is almost entirely NodeJS, though I've spent a lot of spare time learning Rust, Erlang/BEAM/Elixir, and Golang.

Before I'd only been perusing the language features of Go, I'm new to actually creating a real REST app.
I started by researching frameworks, went through Revel and Beego tutorials.  Then I found some very good advice to start with the standard library and only hit the frameworks after knowing the particular requirements demanded of them.

###### So:

I did want live-reload and containerization; for that I've based my implementation from this seed/boilerplate repo and associated article:

Repo: [https://github.com/mikemadisonweb/go-autoreload-example](https://github.com/mikemadisonweb/go-autoreload-example)

Article: [https://mikemadisonweb.github.io/2018/03/06/go-autoreload/](https://mikemadisonweb.github.io/2018/03/06/go-autoreload/).

After that I'll be using this for a ladder:
[https://thenewstack.io/make-a-restful-json-api-go/](https://thenewstack.io/make-a-restful-json-api-go/)

Associated repo: [https://github.com/corylanou/tns-restful-json-api](https://github.com/corylanou/tns-restful-json-api)





### Requirements and How-To-Run:

If Docker-Compose is installed then `docker-compose up` from the project root directory should run it.

I tested it with Postman.


### Status :

I'm pressed for time and pretty new to the Golang ecosystem, accordingly spent a lot of time initially getting the Docker and dev environment set up. Things to do yet:


- CSV import/export
- Testing framework : Usually I'd do the tests before or concurrent with the implementation, but in this case I'll need to research first in order to get a good configuration setup.
- Database:  I'll probably add Redis to the Docker environment. For now the data is held in process memory.

Otherwise, this is a complete -- if spartan -- REST/CRUD API for an address book.   





















______________________________________-
____________________________________

##### Self-reminder / Notes:


Run the app with

```bash
docker-compose up

```

Kill all Docker containers:

```bash
docker stop $(docker ps -a -q)
docker rm $(docker ps -a -q)
```

To get an interactive shell in on the docker container:

`docker exec -ti <container name> /bin/bash`

(`<container name>` is currently `app`)

_________________________



curl -H "Content-Type: application/json" -d '{"name":"Tghimmmmmheyyy"}' http://localhost:8080/persons
