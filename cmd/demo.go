package main

import (
	"context"
	"fmt"
	cli "github.com/ojbkgo/grpc-demo/cli"
	"google.golang.org/grpc"
)

func main() {
	c, e := grpc.Dial("demo-server:80", grpc.WithInsecure())
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	defer c.Close()
	cc := cli.NewHelloClient(c)
	resp, e := cc.Say(context.Background(), &cli.Msg{Val: "123"})
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(resp.Val)
}
