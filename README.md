# backend_with_go
This is a little backend created with Go language

## First steps
If you want to learn more about Go before read all this backend you can go to the folder
```
./first_steps
```

NOTE: This backend was made with go 1.14.2

## Backend
You will se the API REST created with Go this is a simple CRUD of a persons system, you will see in the folder:
```
./api_rest
```

To run the server you need to run the followin command into the `api_rest` folder:
```
go run *.go
```

This server will be run in the port `2020`.

### Connection with GO
We will use the driver of `mgo`, the official documentation is:

http://niemeyer.net/mgo

To install the package you can install with the command:
```
go get gopkg.in/mgo.v2
```

Also you need the package:
```
go get gopkg.in/mgo.v2/bson
```
This last package help to work with the binaries documents of mongo

PSDTA: In mongodb you need the db `go_test_db` and the collection `persons`.

### Routes
I used Gorilla Mux to make the routes of the API. You need to install and you can do it using run the following command
```
go get -u github.com/gorilla/mux
```

The documentation of the package is here
```
https://github.com/gorilla/mux
```

To know if the server is up you will use the route:
```
/health
```

To get all the persons saved in the collection you can use the route:
```
/persons
```

To get the details of one person you can use the route (using GET method):
```
/persons/{id}
```

To save a new person to the collection you can use the POST route:
```
/person
```

POST Body:
```
{
	"id": "4e",
	"name": "Aaron",
	"last_name": "Moran",
	"age": 23,
	"civil_state": 1
}
```

To update the information of a person you can use the PUT route:
```
/person/{id}
```

Body:
```
{
	"id": "4e",
	"name": "Aaron",
	"last_name": "Moran",
	"age": 23,
	"civil_state": 1
}
```

To remove the information of a person you can use the DELELTE route:
```
/person/{id}
```