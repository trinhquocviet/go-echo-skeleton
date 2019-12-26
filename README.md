
# Introduction

Requires before start:
- [Golang version 1.13.1 or lastest](https://golang.org/dl/)
- [Echo Framework](https://echo.labstack.com/)


# Overview Documents

1. DB Diagram [https://dbdiagram.io/](https://dbdiagram.io/)


# Table of Contents
1. [Files & Directories structure](#files--directories-structure)
    - [Project structure](#project-structure)
    - [API package structure](#api-package-structure)
2. [Usage](#usage)

---
## Files & Directories structure

Describe for all files and directories structure using in this project.

### Project structure
```
.
├── api
├── constants
├── internal           # Internal package, only share to internal project
│   ├── helper         # Contain Helper functions
│   ├── message        # The error/response messages
│   ├── service        # internal - background service
├── migrations         # Contain SQL scripts
├── migrations_test    # Contain SQL scripts include mock-data for testing
├── model
├── routes
├── scripts            # All bashscript to help to run on production
├── .env.sample        # Global config for db and docker-compose
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── README.md
└── ...
```

###  API package structure

- One `api` package can have many handlers.
- One `handler` contains controller - model - repository - route.
	- Controller will handle request, response and validate.
	- Model contains all structs/interface and mapping.
	- Repository handles the request to get data from database.
  - Route contains config for routes of handler.

```
api
├── controller_base.go     # The controller inside handler should inherit this
├── repository_base.go     # The repository inside handler should inherit this
├── handler
│   ├── controller.go
│   ├── model.go
│   ├── repository.go
│   ├── route.go           # Config routes for the handler
└── ...
```

## Usage

**NOTE**: clone this project and put follow this path `$GOPATH/src/github.com/trinhquocviet/go-echo-skeleton`, all codebase should place inside this folder.

## Install & start

01 - Run the command to create `.env` file
```bash
cp .env.sample .env
```

02 - Run the command to install the missing libs
```bash
go install
```

03 - Start project
```bash
export $(grep -v '^#' .env | xargs) && go run main.go
```

### Run testing

01 - Run the command below to run the test
```bash
export $(grep -v '^#' .env | xargs) && go test ./... -count=1
```