
inspired by @gbrayhan (https://github.com/gbrayhan/microservices-go)

Project structure

```text
├── Dockerfile
├── LICENSE
├── README.md
├── config.json
├── config.json.example
├── docker
│   └── scripts
│       └── schema.sql
├── docker-compose.yml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
├── github.com/minlebay/pausalac.iml
├── main.go
├── pausalac.iml
├── src
│   ├── application
│   │   ├── security
│   │   │   └── jwt
│   │   │       └── JWT.go
│   │   ├── services
│   │   │   └── Sendgrid.go
│   │   └── usecases
│   │       ├── auth
│   │       │   ├── auth.go
│   │       │   ├── mappers.go
│   │       │   └── structures.go
│   │       └── user
│   │           ├── mappers.go
│   │           ├── structures.go
│   │           └── user.go
│   ├── domain
│   │   ├── Types.go
│   │   ├── errors
│   │   │   ├── Errors.go
│   │   │   └── Gorm.go
│   │   └── user
│   │       └── user.go
│   └── infrastructure
│       ├── repository
│       │   ├── Repository.go
│       │   ├── Utils.go
│       │   ├── config
│       │   │   ├── DataBase.go
│       │   │   └── DataBaseSQL.go
│       │   └── user
│       │       ├── mappers.go
│       │       ├── structures.go
│       │       └── user.go
│       └── rest
│           ├── adapter
│           │   ├── auth.go
│           │   └── user.go
│           ├── controllers
│           │   ├── BindTools.go
│           │   ├── GeneralResponses.go
│           │   ├── Utils.go
│           │   ├── auth
│           │   │   ├── Auth.go
│           │   │   └── Structures.go
│           │   ├── errors
│           │   │   └── Errors.go
│           │   └── user
│           │       ├── Mapper.go
│           │       ├── Requests.go
│           │       ├── Responses.go
│           │       ├── User.go
│           │       └── Validation.go
│           ├── middlewares
│           │   ├── Headers.go
│           │   ├── Interceptor.go
│           │   └── RequiresLogin.go
│           └── routes
│               ├── auth.go
│               ├── routes.go
│               └── user.go
└── utils
├── CreatePDF.go
└── Tools.go
```