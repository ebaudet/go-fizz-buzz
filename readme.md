# Fizz-Buzz REST server - Golang

> This project is a technical test for leboncoin

## Exercise: Write a simple fizz-buzz REST server.

The original fizz-buzz consists in writing all numbers from 1 to 100, and just replacing all multiples of 3 by `fizz`, all multiples of 5 by `buzz`, and all multiples of 15 by `fizzbuzz`.

The output would look like this: `"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,16,...,buzz."`

----

Your goal is to implement a web server that will expose a REST API endpoint that:
- Accepts five parameters: three integers `int1`, `int2` and `limit`, and two strings `str1` and `str2`.
- Returns a list of strings with numbers from `1` to `limit`, where: all multiples of `int1` are replaced by `str1`, all multiples of `int2` are replaced by `str2`, all multiples of `int1` and `int2` are replaced by `str1str2`.

__The server needs to be:__
- Ready for production
- Easy to maintain by other developers

__Bonus:__ add a `statistics` endpoint allowing users to know what the most frequent request has been. This endpoint should:
- Accept no parameter
- Return the parameters corresponding to the most used request, as well as the number of hits for this request"

## Launch

__with docker:__
```sh
docker compose up
```

__without Docker:__
1st time:
```sh
make postgres
make createdb
make migrateup

make server
```
then:
```sh
make start_postgres
make server
```
ℹ️ To get more infos about the makefile commands

```sh
make help
```

Then it's possible to explore the API with [Postman](https://www.postman.com/) importing the following tests :
`postman/fizzbuzz.postman_collection.json`

## Remove

⚠️ remove the database created:
```sh
make fclean
```

## Tests

You launch the tests with the following command:

```sh
make test
```

## Development

* To update the database schema, create a new migration with `make new_migration` then apply your migration with `make migrateup1` - [see migrate documentation](https://pkg.go.dev/github.com/golang-migrate/migrate)
* For new sql request, edit the `/db/query` part, then generate the code with `make sqlc` - [see sqlc documentation](https://docs.sqlc.dev/en/latest/index.html)
* Update the mock to test database interface with `make mock` - [see mock documentation](https://github.com/golang/mock)
