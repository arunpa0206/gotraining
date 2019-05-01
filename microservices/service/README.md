# Service

This is an example of creating a micro service.

## Contents

- main.go - is the main definition of the service, handler and client
- proto - contains the protobuf definition of the API

## Run the example

Run the service

```shell
go run main.go
```

Run the client

```shell
go run main.go --run_client
```

And that's all there is to it.

Test using the following

 curl http://localhost:8000/people
[{"id":"1","firstname":"John","lastname":"Doe","address":{"city":"City X","state":"State X"}},{"id":"2","firstname":"Koko","lastname":"Doe","address":{"city":"City Z","state":"State Y"}}]
$ curl http://localhost:8000/people?1
[{"id":"1","firstname":"John","lastname":"Doe","address":{"city":"City X","state":"State X"}},{"id":"2","firstname":"Koko","lastname":"Doe","address":{"city":"City Z","state":"State Y"}}]
$ curl http://localhost:8000/people?2
[{"id":"1","firstname":"John","lastname":"Doe","address":{"city":"City X","state":"State X"}},{"id":"2","firstname":"Koko","lastname":"Doe","address":{"city":"City Z","state":"State Y"}}]
$ 

