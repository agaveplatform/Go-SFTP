package main

import (
	"../SftpPb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("hello ima a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := sftppb.NewSFTPClient(conn)
	doUnary(c)

}

func doUnary(c sftppb.SFTPClient) {
	fmt.Println("Starting to do uniary rpc client: ")
	req := &sftppb.CopyLocalToRemoteRequest{
		Sftp: &sftppb.Sftp{
			Username: "foo",
			SystemId: "192.168.1.14",
			FileName: "test12.txt",
			HostKey:  "",
			PassWord: "123",
			HostPort: ":2225",
		},
	}
	res, err := c.CopyLocalToRemoteService(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Println("Response from SFTP: " + res.Result)

}
