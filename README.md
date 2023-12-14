# Sound Recommender Golang

A port of [sound-recommender](https://github.com/peter/sound-recommender) from Python to Golang.

## Libraries

* [fiber - web framework](https://docs.gofiber.io/)
* [gorm - ORM](https://gorm.io/index.html)
* [validator - input validation](https://github.com/go-playground/validator)
* [gojsonschema - input validation](https://github.com/xeipuuv/gojsonschema)
* [air - hot reload](https://github.com/cosmtrek/air)

## TODO

* How should Gorm errors be handled?

## Invoking the API with curl

```sh
export BASE_URL=http://localhost:8080

# Admin sounds create - Stairway to Heaven / Led Zeppelin
curl -s -H "Content-Type: application/json" -X POST -d '{"title":"","genres":["rock"],"meta":{"credits":[{"name":"Led Zeppelin","role":"ARTIST"}]},"bpm":82}' $BASE_URL/admin/sounds | jq

# Get sound by ID
curl -s $BASE_URL/sounds/172 | jq

# List sounds
curl -s $BASE_URL/sounds | jq

# Admin sounds update
curl -i -H "Content-Type: application/json" -X PUT -d '{"title":"Stairway to Hell","genres":["death metal"],"credits":[{"name":"Jakob Marklund","role":"ARTIST"}]}' $BASE_URL/admin/sounds/1

# Admin sounds delete
curl -i -X DELETE $BASE_URL/admin/sounds/1
```

## Resources

Learning Go:

* [My own notes on learning Go](https://github.com/peter/learning/tree/master/languages/golang)

* [Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started)
* [Official Go Learning Portal](https://go.dev/learn/)
* [Learn Go in Y minutes](https://learnxinyminutes.com/docs/go/)
* [The Golang Handbook – A Beginner's Guide to Learning Go](https://www.freecodecamp.org/news/learn-golang-handbook/)
* [Go by Example](https://gobyexample.com/)
* [Go Standard library Documentation](https://pkg.go.dev/std)
* [Learn Go Programming - Golang Tutorial for Beginners](https://www.youtube.com/watch?v=YS4e4q9oBaU&t=537s)

Building REST APIs with Go:

* [Build a REST API in Go using Fiber + GORM](https://www.youtube.com/watch?v=dpx6hpr-wE8&t=1773s)
* [Schema Validation in Go using Validator and Fiber](https://dev.to/franciscomendes10866/how-to-validate-data-in-golang-1f87)
* [Build a RESTful CRUD API with Golang (Gorm/Gin/Postgres article + github code)](https://github.com/wpcodevo/golang-gorm-postgres)
* [Tutorial: Developing a RESTful API with Go and Gin (go.dev)](https://go.dev/doc/tutorial/web-service-gin)
* [Go REST Guide. Gin Framework (JetBrains)](https://www.jetbrains.com/guide/go/tutorials/rest_api_series/gin/)
* [Build a RESTful CRUD API with Golang Gin and Gorm](https://lemoncode21.medium.com/build-a-restful-crud-api-with-golang-gin-and-gorm-e1e976ef5b9f)
* [Optimizing GoLang APIs with Gin, New Relic, and Swagger](https://blog.stackademic.com/optimizing-golang-apis-with-gin-new-relic-and-swagger-a-comprehensive-guide-d60cea368fbe)
* [Performing CRUD Operations in PostgreSQL with Go (without an ORM)](https://edwinsiby.medium.com/performing-crud-operations-in-postgresql-with-go-42657761125c)

Frameworks and libraries:

* [Gin Web Framework](https://github.com/gin-gonic/gin)
* [GORM (ORM)](https://github.com/go-gorm/gorm)
