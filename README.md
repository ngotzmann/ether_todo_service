# Example Service

This service is an example service in golang in onion architecture and DDD.
The project has some features which will go in own libraries on boilerplate and real projects.

Following libraries emerged of this example:
* gorror
* config_reader
* jwt_user_management
* TODO: Add more

## Used third party libraries
* As web framework [echo](https://echo.labstack.com/) will be used
  * Sessions comes from echo, they use [gorilla](https://github.com/gorilla/sessions) 
  * For jwt is [jwt-go](github.com/dgrijalva/jwt-go) used
* Configurations will read with the fantastic [viper](https://github.com/spf13/viper)
* For errors, I use my own library [gorror](https://github.com/ngotzmann/gorror)
* [Logrus](https://github.com/sirupsen/logrus) will be used for logging
* TODO: Database/ORM?

## Structure
The `main.go` contains the hole setup of the service. 
There will echo configured, the routes will be set...

### Architecture
As software architecture the onion architecture was choose.
This allows a very modularity structure, the 'entity' module in onion architecture accommodated the domain with his business logic and application logic.

The entity module contains following files:
* `user.go` contains the business logic, it contains the struct which will used for json and the database.
  * There are just a few dependencies to other libraries.
* `repository.go` is and interface for datastores implementations.
* `service.go` contains application logic which is in the domain context, like call repo...

At the usecase module is just one file the `usecase.go`:
* Here are application logic which can use other domains and libraries.
* At this specific service the usecase file use components of the echo framework for jwt and session integration.

The controller package represent the interface module, it contains the API and the implementation of the repository.

![onion architecture](https://miro.medium.com/max/1400/1*B7LkQDyDqLN3rRSrNYkETA.jpeg)

[declaration](https://medium.com/@hatajoe/clean-architecture-in-go-4030f11ec1b1)

## How to use
Configure this service in `config/config-env.yml`.
