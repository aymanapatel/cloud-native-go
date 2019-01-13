package main

import (
	"fmt"

	proto "github.com/"
	micro "github.com/micro/go-micro"
)


	// The Greeter API
	type Greeter struct{}


	// Hello is a Greeter API method
	func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse){
		rsp.Greeting = "Hello" + req.Name
	}