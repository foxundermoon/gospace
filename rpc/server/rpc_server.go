package server

import (
	"errors"
	"fmt"
)

type Args struct {
	A, B int
}
type Arith int
type Quotient struct {
	Quo, Rem int
}

func (t *Arith) Mutiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	fmt.Println("Mutiply @server ")
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	fmt.Println("Divede @server by rpc")
	return nil
}
