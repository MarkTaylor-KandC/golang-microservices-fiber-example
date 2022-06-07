# Example Microservice with Fiber

## Introduction

This is an example project that details the steps needed to create a basic HTTP server using [Fiber](https://docs.gofiber.io/). It is assumed that the reader has at least a basic understanding of Golang, Microservice architectures and software development practices. 

In this example we are going to:

* [Set up a new Project](#set-up-a-new-project)
* [Create a basic HTTP Server](#creating-a-basic-http-server)
* [Containerise the Server with Docker](#containerisation-with-docker)
* [Test the server with a browser](#testing-with-a-browser)
* [Create some unit tests to test the server](#testing-with-unit-tests)

By the end of this you will have a service that listens for an HTTP GET request on a given port and returns the String "Hello World!". The service will have a number of unit tests and will act at the basis to build out a more complex service. We will also containerise this service with docker allowing it to easily be deployed on any machine. I wouldn't recommend putting it into production but feel free to disregard that recommendation.

## Set up a new project

To start with we need to set up a new project directory

``` bash
mkdir fiber-example
cd fiber-example
```
Then we need to initalise the go module, in this case I used the github URL for the project as the module name and created `hello.go` to act as the main source file for the service.
```
go mod init github.com/MarkTaylor-KandC/golang-microservices-fiber-example
touch hello.go
```
Now that we have some basic files that we are working with we should probably version control those with Git
```
git init
```

## Creating a basic HTTP Server

This is the same as the example file provided on the Fiber documentation [file](https://docs.gofiber.io/#hello-world)

[Source](./hello.go)

``` go
package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	app.Listen(":3000")
}
```

In order to use fiber, we are going to need to vendor it:

``` bash
go get -u github.com/gofiber/fiber/v2
```

Then we can install and run the service:

``` bash
go install github.com/MarkTaylor-KandC/golang-microservices-fiber-example
go run hello.go   
```

The source code for the server is fairly simple and if you are familar with [Express](https://www.npmjs.com/package/express) this probably all looks very familiar. 

* `app` is the instance of fiber.
* `Get` is the HTTP Request Method that we want to expose.
* `"/"` is the path we want to expose the method on.
* `func(c *fiber.Ctx) error` is the callback function that is invoked when a request is received and matched to the method and path (route). It contains the context which holds details on the HTTP request and response and allows us to send the response. 
* `return c.SendString("Hello World!")` sends a string as an HTTP response object and returns the result of that operation.

## Containerisation with Docker

As this is a simple application it is easy to create a docker file to containerise the application. 

[Source](./dockerfile)

```dockerfile
FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /golang-microservices-fiber-example

EXPOSE 3000

CMD [ "/golang-microservices-fiber-example" ]
```

The docker file builds the application and then runs it, exposing the service on port 3000. 

We can build it with the following command:

```bash
docker build -t golang-microservices-fiber-example .
```

And run it with:

```bash
docker run -p 3000:3000 -t golang-microservices-fiber-example
```

This will then provide the same result as running above but from within a docker container which should be portable. 

## Testing with a browser

We can easily test the application using a browser by navigating to [http://127.0.0.1:3000/](http://127.0.0.1:3000/) with either the program running directly or running in the docker container. This is good for a sanity check but we want to also have something that we can run automatically as part of our build process.

## Testing with Unit tests