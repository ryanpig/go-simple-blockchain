## Table of content

## Description

Build a simple blockchain w/ simple functionality in GO

## Goal
- Create a new blockchain by using slices of blocks 
- Generate new block w/ SHA-256 hash (package: "crypto/sha-256")
- Solve conflict blockchains by choosing longer length
- Build a webserver providing REST API that allows user to view blocks and add new block client browser. (package: "mux" , "net/http")

Other:
- Logging function (package: "log")
- Print struct in a elegant text format (package: "Spew", method:Spew.Dump)


## TODO 
- prrof-of-work 
- network broadcasting
