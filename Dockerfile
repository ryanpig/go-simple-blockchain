# Start from golang v1.11 base image
FROM golang:1.11

# Add Maintainer info.
LABEL maintainer="ryanpig <ryanpig@gmail.com>"

# set working directory inside the container
WORKDIR $GOPATH/src/github.com/go-simple-blockchain

# Copy everythin from currrent directory to the PWD inside the container
COPY . . 

# Download dependencies
Run go get -u github.com/davecgh/go-spew/spew
Run go get -u github.com/gorilla/mux
Run go get -u github.com/joho/godotenv
Run go get -u github.com/shurcooL/githubv4
Run go get -u golang.org/x/oauth2

# exposes port 8080 to outside world
EXPOSE 8080

# Run the executable
CMD ["go", "run", "main.go"]
