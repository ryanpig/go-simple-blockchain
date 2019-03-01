## Table of content
- [Table of content](#table-of-content)
- [Goal](#goal)
- [Status](#status)
- [DEMO](#demo)
- [TODO](#todo)
- [Dependencies installation](#dependencies-installation)
- [Change listening port by modifying .env file](#change-listening-port-by-modifying-env-file)


## Goal
- Inspired by [Reference](#reference), I decided to build a simple and small blockchain written in GO, that more features can be added later on.  

## Status
- Create a small blockchain to record Github project history (Note: Now the blockchain is created from reading test file, I will switch it by using github API)
- Generate new block w/ SHA-256 hash 
- Build a webserver providing REST API that allows user to view blocks from client browser. 

## DEMO 
- [AWS EC2](http://ec2-35-180-136-222.eu-west-3.compute.amazonaws.com:8080/#) 

Note: It's just a small prototype now.

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
```
## Change listening port by modifying .env file
- `vim .env` 

## Reference 
- [Code your own blockchain in less than 200 lines of Go!](https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc)
