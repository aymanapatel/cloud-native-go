# Implement Sync RPC Calls With Binary Protocols


## Installation


1. [protoc/ Protocol Buffer installation](https://github.com/protocolbuffers/protobuf/blob/master/src/README.md)
    This doesn't include protoc for golang.

2. [Installation](https://grpc.io/docs/quickstart/go.html) of protoc for golang.    


3. Set environment variable for $GOPATH and add to $PATH

    export GOPATH=$HOME/go
    export PATH=$PATH:$GOROOT/bin:$GOPATH/bin


    
4. Install gRPC
`go get -u google.golang.org/grpc`