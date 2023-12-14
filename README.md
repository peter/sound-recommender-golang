# Sound Recommender Golang

A port of [sound-recommender](https://github.com/peter/sound-recommender) from Python to Golang.

## Libraries

* [fiber - web framework](https://docs.gofiber.io/)
* [gorm - ORM](https://gorm.io/index.html)
* [validator - input validation](https://github.com/go-playground/validator)
* [air - hot reload](https://github.com/cosmtrek/air)

## TODO

* How should Gorm errors be handled?

## Invoking the API with curl

```sh
export BASE_URL=http://localhost:8080

# Admin sounds create - Stairway to Heaven / Led Zeppelin
curl -s -H "Content-Type: application/json" -X POST -d '{"data":[{"title":"Stairway to Heaven","genres":["rock"],"meta"{"credits":[{"name":"Led Zeppelin","role":"ARTIST"}]},"bpm":82}]}' $BASE_URL/admin/sounds | jq

# Admin sounds create - Halo / Beyonce
curl -s -H "Content-Type: application/json" -X POST -d '{"data":[{"title":"Halo","genres":["dance pop"],"credits":[{"name":"Beyonce","role":"ARTIST"}],"bpm":80}]}' $BASE_URL/admin/sounds | jq

# Admin sounds create - Blank Space / Taylor Swift
curl -s -H "Content-Type: application/json" -X POST -d '{"data":[{"title":"Blank Space","genres":["pop"],"credits":[{"name":"Taylor Swift","role":"ARTIST"}],"bpm":96}]}' $BASE_URL/admin/sounds | jq

# Admin sounds create - Hips Don't Lie / Shakira Featuring Wyclef Jean
curl -s -H "Content-Type: application/json" -X POST -d '{"data":[{"title":"Hips Dont Lie","genres":["latin pop","reggaeton"],"credits":[{"name":"Shakira Featuring Wyclef Jean","role":"ARTIST"}],"bpm":100}]}' $BASE_URL/admin/sounds | jq

# Admin sounds create - Wrecking Ball / Miley Cyrus
curl -s -H "Content-Type: application/json" -X POST -d '{"data":[{"title":"Wrecking Ball","genres":["pop"],"credits":[{"name":"Miley Cyrus","role":"ARTIST"}],"bpm":120}]}' $BASE_URL/admin/sounds | jq

# Admin sounds create - Livin' La Vida Loca / Ricky Martin
curl -s -H "Content-Type: application/json" -X POST -d '{"data":[{"title":"Livin La Vida Loca","genres":["latin pop","dance"],"credits":[{"name":"Ricky Martin","role":"ARTIST"}],"bpm":178}]}' $BASE_URL/admin/sounds | jq

# Admin sounds create - Single Ladies (Put A Ring On It) / Beyonce
curl -s -H "Content-Type: application/json" -X POST -d '{"data":[{"title":"Single Ladies (Put A Ring On It)","genres":["dance pop","r&b"],"credits":[{"name":"Beyonce","role":"ARTIST"}],"bpm":97}]}' $BASE_URL/admin/sounds | jq

# Admin sounds create - Master of Puppets / Metallica
curl -s -H "Content-Type: application/json" -X POST -d '{"data":[{"title":"Master of Puppets","genres":["thrash metal"],"credits":[{"name":"Metallica","role":"ARTIST"}],"bpm":220}]}' $BASE_URL/admin/sounds | jq

# Get sounds
curl -s $BASE_URL/sounds/1 | jq

# List sounds
curl -s $BASE_URL/sounds | jq

# List sounds by metallica
curl -s $BASE_URL/sounds?query=metallica | jq

# Admin sounds update
curl -i -H "Content-Type: application/json" -X PUT -d '{"title":"Stairway to Hell","genres":["death metal"],"credits":[{"name":"Jakob Marklund","role":"ARTIST"}]}' $BASE_URL/admin/sounds/1

# Create playlist
curl -H "Content-Type: application/json" -X POST -d '{"data":[{"title":"Greatest of all time", "sounds":[1]}]}' $BASE_URL/playlists

# Get playlist
curl -s $BASE_URL/playlists/1 | jq

# List playlist
curl -s $BASE_URL/playlists | jq

# Update playlist
curl -i -H "Content-Type: application/json" -X PUT -d '{"title":"Greatest of all time!!!", "sounds":[1]}' $BASE_URL/playlists/1

# Get playlist
curl -s $BASE_URL/playlists/1 | jq

# Get recommendation for playlist
# Note that you get similariy scores in the list of recommendations
curl -s $BASE_URL/sounds/recommended?playlistId=1 | jq

# Get recommendation for sound
curl -s $BASE_URL/sounds/recommended?soundId=1 | jq

# Get recommendation for sound without using pgvector (does similarity sort in memory based on all sounds)
curl -s "$BASE_URL/sounds/recommended?strategy=memory&soundId=1" | jq

# Delete playlist
curl -i -X DELETE $BASE_URL/playlists/1

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
* [Build a RESTful CRUD API with Golang (Gorm/Gin/Postgres article + github code)](https://github.com/wpcodevo/golang-gorm-postgres)
* [Tutorial: Developing a RESTful API with Go and Gin (go.dev)](https://go.dev/doc/tutorial/web-service-gin)
* [Go REST Guide. Gin Framework (JetBrains)](https://www.jetbrains.com/guide/go/tutorials/rest_api_series/gin/)
* [Build a RESTful CRUD API with Golang Gin and Gorm](https://lemoncode21.medium.com/build-a-restful-crud-api-with-golang-gin-and-gorm-e1e976ef5b9f)
* [Optimizing GoLang APIs with Gin, New Relic, and Swagger](https://blog.stackademic.com/optimizing-golang-apis-with-gin-new-relic-and-swagger-a-comprehensive-guide-d60cea368fbe)
* [Performing CRUD Operations in PostgreSQL with Go (without an ORM)](https://edwinsiby.medium.com/performing-crud-operations-in-postgresql-with-go-42657761125c)

Frameworks and libraries:

* [Gin Web Framework](https://github.com/gin-gonic/gin)
* [GORM (ORM)](https://github.com/go-gorm/gorm)
