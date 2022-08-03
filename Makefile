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

new_migration:
	@if [[ -z "${NAME}" ]]; then\
		printf "Name of the migration: ";\
		read NAME;\
		migrate create -ext sql -dir db/migration -seq $${NAME};\
	else\
		migrate create -ext sql -dir db/migration -seq $${NAME};\
	fi

postgres:
	# docker network create fizzbuzz-network
	docker run --name pg_fizzbuzz --network fizzbuzz-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it pg_fizzbuzz createdb --username=root --owner=root fizzbuzz

dropdb:
	docker exec -it pg_fizzbuzz dropdb fizzbuzz

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fizzbuzz?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fizzbuzz?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fizzbuzz?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fizzbuzz?sslmode=disable" -verbose down 1

help:
	@printf $(YELLOW)$(BOLD)"HELP\n--------------------\n"$(NORMAL)
	@printf $(NORMAL)"-> make "$(BOLD)"server"$(NORMAL)": launches the server\n"
	@printf $(NORMAL)"-> make "$(BOLD)"test"$(NORMAL)": runs all the tests\n"
	@printf $(NORMAL)"-> make "$(BOLD)"test_nocache"$(NORMAL)": runs all the tests without caching\n"
	@printf $(NORMAL)"-> make "$(BOLD)"new_migration [-e NAME=<migration_name>]"$(NORMAL)": creates a new migration file\n"
	@printf $(NORMAL)"-> make "$(BOLD)"migrateup"$(NORMAL)": migrates all the pending migrations\n"
	@printf $(NORMAL)"-> make "$(BOLD)"migratedown"$(NORMAL)": rolldowns all the migrations\n"
	@printf $(NORMAL)"-> make "$(BOLD)"migrateup1"$(NORMAL)": migrates the next migration\n"
	@printf $(NORMAL)"-> make "$(BOLD)"migratedown1"$(NORMAL)": rolldowns the last migration\n"
	@printf $(NORMAL)"-> make "$(BOLD)"postgres"$(NORMAL)": runs the docker postgres14 container\n"
	@printf $(NORMAL)"-> make "$(BOLD)"createdb"$(NORMAL)": creates the database on docker postres container\n"
	@printf $(NORMAL)"-> make "$(BOLD)"dropdb"$(NORMAL)": drop the database on docker postres container\n"
	@printf $(NORMAL)"-> make "$(BOLD)"help | usage"$(NORMAL)": shows the help\n"
	@printf $(YELLOW)$(BOLD)"--------------------\n"$(NORMAL)

usage: help

.PHONY: server test test_nocache help usage postgres createdb dropdb new_migration migrateup migrateup1 migratedown migratedown1
