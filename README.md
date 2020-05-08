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
go run main.go
```

This server will be run in the port `2020`.

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

To get the details of one person you can use the route:
```
/persons/{id}
```
