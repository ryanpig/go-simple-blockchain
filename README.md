[![Build Status](https://travis-ci.org/ryanpig/go-simple-blockchain.svg?branch=master)](https://travis-ci.org/ryanpig/go-simple-blockchain)

## Table of content
- [Table of content](#table-of-content)
- [Goal](#goal)
- [Features](#features)
- [DEMO](#demo)
- [TODO](#todo)
- [Dependencies installation](#dependencies-installation)
- [Authentification](#authentification)
- [Change listening port by modifying .env file](#change-listening-port-by-modifying-env-file)


## Goal
- Inspired by [Reference](#reference), I decided to build a simple and small blockchain written in GO, which stores Github project information in this blockchain.   

## Features
- The blockchain maintains project information from Github 
- Use Github APIv.4(GraphQL) to retrieve project data
- Use SHA-256 to hash project data
- Build a webserver providing REST API that allows user to view blocks from client browser. 
- Run in the docker container w/ environment varailbe  (

## DEMO 
- [AWS EC2](http://52.47.189.34:8080/) 

## TODO
- Allow user to write a new block to blockchain from POST request
- Solve conflict blockchains by choosing longer length
- Add mining algorithm, e.g. prrof-of-work 
- Network broadcasting , p2p

## Dependencies installation
- all in one by executing `. dependencies_install.sh`
- manually install by commands below 
```
go get github.com/davecgh/go-spew/spew
go get github.com/gorilla/mux
go get github.com/joho/godotenv
go get -u github.com/shurcooL/githubv4
go get -u golang.org/x/oauth2
```

## Authentification
- Create github token, and save envrionment variable `GITHUB_TOKEN=xxxxx` w/ the token

## Change listening port by modifying .env file
- `vim .env` 

## Reference 
- [Code your own blockchain in less than 200 lines of Go!](https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc)
- [Github API graphQL explorer](https://developer.github.com/v4/explorer/)
- [GIthub API v4 official](https://github.com/shurcooL/githubv4)
