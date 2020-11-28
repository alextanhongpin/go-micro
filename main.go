package main

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/micro/v3/service"
	proto "github.com/micro/services/helloworld/proto"
)

func main() {
	srv := service.New()
	client := proto.NewHelloworldService("helloworld", srv.Client())
	rsp, err := client.Call(context.Background(), &proto.Request{
		Name: "john",
	})
	if err != nil {
		fmt.Println("Error calling helloworld:", err)
		return
	}
	fmt.Println("Response: ", rsp.Msg)

	// Let's delay the process for exiting for reasons you will see below.
	time.Sleep(time.Second * 5)
}
