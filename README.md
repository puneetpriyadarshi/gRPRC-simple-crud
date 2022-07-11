# go-starter-kit

### Pre-requisites
- [Golang migrate](https://github.com/golang-migrate/migrate)
- [SQL Boiler](https://github.com/volatiletech/sqlboiler)


## Installation

A little intro about the installation. 
```
$ https://github.com/amanjain97/go-starter-kit
$ cd go-starter-kit
$ go mod tidy
$ go run main.go
```

### Execute the migrations 
`make migrate-prepare` install dependencies
`make db-migrateup` executes the sql in `db/migrations` migrations to postgres db

### Create the repository layer
`make repository-gen` will generate the api/repository package using sql boiler and provides 
necessary interface for accessing data from db.

### Connect to the firestore db
`export FIREBASE_JSON=enter your downloaded serviceAccount.json path"` Set environment variable for connect to the firestore db.

### Connect to the google auth
`export GOOGLE_AUTH_JSON=enter your downloaded google_auth.json path"` Set environment variable for connect to the google auth.

### For SSO login
make sure you have `sso_private_key: "save cert file", sso_certificate: "save key file" and sso_service_url: "https://samltest.id/saml/idp"` all the key and value in organization table.  now enter the url in browser `http://localhost:5000/sso/login?name=your user_name`. if the user_name is correct then you are redirect on sso login page.

### What is Cypress ?

Cypress is a next generation  testing tool built for the modern web. We address the key pain points developers and QA engineers face when testing modern applications.

It allows you to write all types of testing : 

End to end tests : End-to-end testing, also known as E2E testing, is a methodology used for ensuring that applications behave as expected and that the flow of data is maintained for all kinds of user tasks and processes.

Integration tests : Integration testing -- also known as integration and testing (I&T) -- is a type of software testing in which the different units, modules or components of a software application are tested as a combined entity.

Unit tests : Unit testing is a software development process in which the smallest testable parts of an application, called units, are individually and independently scrutinized for proper operation. 

In a nutshell , cypress can test anything that runs in a browser.

### e2e testing using cypress : 

Cypress was originally designed to run end-to-end (E2E) tests on anything that runs in a browser. A typical E2E test visits the application in a browser and performs actions via the UI just like a real user would.

For better understanding and to have a near perfect approach regarding how to do e2e testing using cypress ,  we have divided our whole testing process into several steps.

The steps are as follows : 
The first step is to run the server which we are going to test.For that run these  commands in your terminal : 

```export FIREBASE_JSON=service.json
   export GOOGLE_AUTH_JSON=gouth.json
   export APIENV=staging(development for actual db)
   go run cmd/main.go
```
