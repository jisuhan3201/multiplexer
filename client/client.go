package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jisuhan3201/multiplexer/multiplexerpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hi i am client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot get client : %v", err)
	}
	defer cc.Close()
	c := multiplexerpb.NewGLogServierClient(cc)
	doUnary(c)
}

func doUnary(c multiplexerpb.GLogServierClient) {
	fmt.Println("Starting to multiplexer doUnary RPC...")
	req := &multiplexerpb.GLogRequest{
		Id:   "1234",
		Type: "GLog",
	}
	res, err := c.GLog(context.Background(), req)
	if err != nil {
		log.Fatalf("cannot multiplexing : %v\n", err)
	}

	log.Printf("Response from Caculator : %v", res.Result)
}
