#!/bin/sh
# print struct
go get github.com/davecgh/go-spew/spew
# routing w/ different function handler
go get github.com/gorilla/mux
# read .env file 
go get github.com/joho/godotenv
# github api v4 (GraphQL) , requires v.18
go get -u github.com/shurcooL/githubv4
# authentification
go get -u golang.org/x/oauth2
