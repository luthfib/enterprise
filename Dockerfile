# Base build image
FROM golang:latest
 
ADD . /go/src/github.com/luthfi/enterprise
WORKDIR /go/src/github.com/luthfi/enterprise
# Force the go compiler to use modules
ENV GO111MODULE=on
 
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .
 
#This is the ‘magic’ step that will download all the dependencies that are specified in 
# the go.mod and go.sum file.
RUN go mod download
 
# Here we copy the rest of the source code
COPY . .

# RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' /enterprise

# Document that the service listens on port 8080.
EXPOSE 8080
CMD ["go", "run", "main.go"]


