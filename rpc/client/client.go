package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpc/server"
)

func main() {
	c, err := rpc.DialHTTP("tcp", "10.80.5.222:1234")
	if err != nil {
		log.Fatal("dialing....:", err)
	}

	args := &server.Args{7, 8}
	var reply int
	err = c.Call("Arith.Mutiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	quotient := new(server.Quotient)
	divcall := c.Go("Arith.Divide", args, &quotient, nil)

	replyCall := <-divcall.Done

	fmt.Printf("Arith.Divide by go %d/%d=%d\n", args.A, args.B, replyCall)

}
