package hello

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testing"
)

func TestHelloClient_Say(t *testing.T) {
	c, e := grpc.Dial("127.0.0.1:81", grpc.WithInsecure())
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	defer c.Close()
	cc := NewHelloClient(c)
	resp, e := cc.Say(context.Background(), &Msg{Val: "123"})
	if e != nil {
		fmt.Println(e.Error())
		return
	}
	fmt.Println(resp.Val)
}