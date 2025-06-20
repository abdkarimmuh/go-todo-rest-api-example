# Go Todo REST API Example

A RESTful API example for simple todo application with Go

It is a just simple tutorial or example for making simple RESTful API with Go using **gorilla/mux** (A nice mux library) and **gorm** (An ORM for Go)

## Installation & Run

```bash
# Download this project
go get github.com/abdkarimmuh/go-todo-rest-api-example
```

Before you run the API, you must create a `.env` file containing your database configuration by copying from [.env.example](https://github.com/abdkarimmuh/go-todo-rest-api-example/blob/main/.env.example).

```env
DB_DIALECT=mysql
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=password
DB_NAME=todoapp
```

```bash
# Build and Run
cd go-todo-rest-api-example
go build
./go-todo-rest-api-example

# API Endpoint : http://127.0.0.1:3000
```

## Structure

```text
├── app
│   ├── app.go
│   ├── handler          // Our API core handlers
│   │   ├── common.go    // Common response functions
│   │   ├── projects.go  // APIs for Project model
│   │   └── tasks.go     // APIs for Task model
│   └── model
│       └── model.go     // Models for our application
├── config
│   └── config.go        // Configuration
└── main.go
```

## API

### /projects

* `GET` : Get all projects
* `POST` : Create a new project

#### /projects/:title

* `GET` : Get a project
* `PUT` : Update a project
* `DELETE` : Delete a project

#### /projects/:title/archive

* `PUT` : Archive a project
* `PUT` : Restore a project

#### /projects/:title/tasks

* `GET` : Get all tasks of a project
* `POST` : Create a new task in a project

#### /projects/:title/tasks/:id

* `GET` : Get a task of a project
* `PUT` : Update a task of a project
* `DELETE` : Delete a task of a project

#### /projects/:title/tasks/:id/complete

* `PUT` : Complete a task of a project
* `PUT` : Undo a task of a project

## Todo

* [x] Support basic REST APIs.
* [ ] Support Authentication with user for securing the APIs.
* [ ] Make convenient wrappers for creating API handlers.
* [ ] Write the tests for all APIs.
* [x] Organize the code with packages
* [ ] Make docs with GoDoc
* [ ] Building a deployment process
