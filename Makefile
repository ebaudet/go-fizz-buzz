# ********************************************************** #
#                                                            #
#    Makefile                                                #
#                                                            #
#    by: ebaudet <emilien.baudet@gmail.com>                  #
#                                                            #
# ********************************************************** #

server:
	go run main.go

test:
	go test -v -cover ./...

test_nocache:
	go clean -testcache
	make test

.PHONY: server test test_nocache
