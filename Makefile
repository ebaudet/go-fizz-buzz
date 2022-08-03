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

help:
	@printf $(YELLOW)$(BOLD)"HELP\n--------------------\n"$(NORMAL)
	@printf $(NORMAL)"-> make "$(BOLD)"server"$(NORMAL)": launch the server\n"
	@printf $(NORMAL)"-> make "$(BOLD)"test"$(NORMAL)": run all the tests\n"
	@printf $(NORMAL)"-> make "$(BOLD)"test_nocache"$(NORMAL)": run all the tests without caching\n"
	@printf $(NORMAL)"-> make "$(BOLD)"help | usage"$(NORMAL)": show the help\n"
	@printf $(YELLOW)$(BOLD)"--------------------\n"$(NORMAL)

usage: help

.PHONY: server test test_nocache help usage
