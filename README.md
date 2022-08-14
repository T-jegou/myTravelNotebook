# MyTravelBook
## _Write, Store and share your vacation stories_

MyTravelBook is a side-project based on a simple idea, design from scrath an API with go and go-swagger aim to understand the fundamentals of Rest principles and some extra tools to create a complete backend solution.

In order to challenge myself by adding some complexity but meanwhile create a better fault-tolerant tool, I decided to add an ORM (GORM), a logger (logrus), a http middleware (negroni) and some other cools stuff to to increase skills in Golang ecosysteme.

## Features
- Create a travel stories
- Read a travel stories 
- Update a travel stories
- Delete a travel stories

CRUD
## Generate docs and server
MyTravelBook require :
- swagger-merger
- go-swagger
- go 1.14>


```sh
swagger-merger -i swagger/index.yaml -o docs/api_docs/bundle.yaml
swagger validate docs/api_docs/bundle.yaml 
swagger generate server -A myTravelBook -P models.Principal -t ./swagger_gen -f docs/api_docs/bundle.yaml
go get -u -f ./swagger_gen/... 
```

## Run API

Once the documentation has been generated the API via the swagger description, we can launch the it

```sh
go run swagger_gen/cmd/my-travel-notebook-server/main.go --port 55600
```

