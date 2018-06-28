# Article API

### Get started

Begin by installing the dependencies by running the `make vendor` command. In this case there's only two so alternatively you can run `go get -u github.com/gorilla/pat` and `go get -u github.com/stretchr/testify/assert`.

Once the dependencies are installed `make run` will start the API at `localhost:8080`.

### Endpoints

There are 3 endpoints available.

`POST /articles` receives a JSON blob containing article data and stores it. The article ID is returned to the user.

`GET /articles/{id}` returns the JSON representation of the article with the given ID.

`GET /tags/{tagName}/{date}` returns the JSON representation of tag data for the given tag on the given day.

An article has the following attributes `id`, `title`, `date`, `body` and `tags` like the one below:
```
{
  "title": "an article title",
  "date" : "20090909",
  "body" : "an article body",
  "tags" : ["health", "potatoes", "life"]
}
```

### Assumptions
* Ids are automatically generated in this API meaning that when a user posts an article they should not include the ID field. They generated ID will be returned to the user.
* The `POST` endpoint only supports new articles. It does not allow the user to upsert or update articles.
* The date format used in this API is `YYYYMMDD`.
* Tags are unique in an article e.g. `health` will only appear once in the `tags` array, never twice.

### Description
This Article API is a Golang REST API.

The project was structured to be concise and meaningful. It consists of 4 packages: `api`, `cmd`, `handlers` and `svc`.
* `api` contains the structs used to store the Article and Tag representation.
* `cmd` contains the `main` function.
* `handlers` contains the router and the handlers that it routes to.
* `svc` contains the services to store and retrieve Article and Tag data.

For every new article, the date field is validated by checking whether it has the format `YYYYMMDD`. If it does not, the request is not accepted and an error is returned to the user stating that they have used the incorrect date format. Similarly, if the user specifies the wrong date format on the `GET` tag endpoint, an invalid date error is also returned to the user. To further improve the data validation, new articles being posted to the API should also be checked that their `body`, `title` and `tag` fields are not empty.

Unit tests have been included in `svc/article_test.go` in a table driven format. Unit testing should also be performed in the `handlers` package.

There are two external packages used in this project:
* `bmizerany/pat` was chosen because the project required a simple URL pattern muxer.
* `stretchr/testify` was chosen for the ease of assertions.