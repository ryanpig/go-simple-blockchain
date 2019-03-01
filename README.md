## Table of content

## Description

Build a simple blockchain w/ simple functionality in GO

## Goal
- Create a new blockchain by using slices of blocks 
- Generate new block w/ SHA-256 hash 
- Solve conflict blockchains by choosing longer length
- Build a webserver providing REST API that allows user to view blocks and add new block client browser. 

Other:
- Logging function 
- Print struct in a elegant text format by using `spew` package

## Dependencies installation
- all in one by executing `. dependencies_install.sh`
- manually install w/ 
```
go get github.com/davecgh/go-spew/spew
go get github.com/gorilla/mux
go get github.com/joho/godotenv
```


## TODO 
- prrof-of-work 
- network broadcasting
