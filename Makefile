# ********************************************************** #
#                                                            #
#    Makefile                                                #
#                                                            #
#    by: ebaudet <emilien.baudet@gmail.com>                  #
#                                                            #
# ********************************************************** #
NORMAL = "\x1B[0m"
YELLOW = "\x1B[33m"
BOLD = "\e[1m"

server:
	go run main.go

test:
	go test -v -cover ./...

test_nocache:
	go clean -testcache
	make test

postgres:
	# docker network create fizzbuzz-network
	docker run --name pg_fizzbuzz --network fizzbuzz-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it pg_fizzbuzz createdb --username=root --owner=root fizzbuzz

dropdb:
	docker exec -it pg_fizzbuzz dropdb fizzbuzz

help:
	@printf $(YELLOW)$(BOLD)"HELP\n--------------------\n"$(NORMAL)
	@printf $(NORMAL)"-> make "$(BOLD)"server"$(NORMAL)": launches the server\n"
	@printf $(NORMAL)"-> make "$(BOLD)"test"$(NORMAL)": runs all the tests\n"
	@printf $(NORMAL)"-> make "$(BOLD)"test_nocache"$(NORMAL)": runs all the tests without caching\n"
	@printf $(NORMAL)"-> make "$(BOLD)"postgres"$(NORMAL)": runs the docker postgres14 container\n"
	@printf $(NORMAL)"-> make "$(BOLD)"createdb"$(NORMAL)": creates the database on docker postres container\n"
	@printf $(NORMAL)"-> make "$(BOLD)"dropdb"$(NORMAL)": drop the database on docker postres container\n"
	@printf $(NORMAL)"-> make "$(BOLD)"help | usage"$(NORMAL)": shows the help\n"
	@printf $(YELLOW)$(BOLD)"--------------------\n"$(NORMAL)

usage: help

.PHONY: server test test_nocache help usage postgres createdb dropdb 
