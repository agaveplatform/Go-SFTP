package main

import (
	"../SftpPb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"time"
)

func main() {
	fmt.Println("Starting SFTP client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := sftppb.NewSFTPClient(conn)
	// push file test
	doUnaryPut(c)
	//pull the test file
	doUnaryPull(c)

	//p

}

func doUnaryPut(c sftppb.SFTPClient) {
	fmt.Println("Starting to do uniary rpc client: ")
	startPulltime := time.Now()

	req := &sftppb.CopyLocalToRemoteRequest{
		Sftp: &sftppb.Sftp{
			Username: "foo",
			SystemId: "192.168.1.14",
			FileName: "test2.txt",
			HostKey:  "",
			PassWord: "123",
			HostPort: ":2225",
		},
	}
	res, err := c.CopyLocalToRemoteService(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	secs := time.Since(startPulltime).Seconds()
	log.Println("Response from SFTP: " + res.Result)
	fmt.Println("Function Time: " + strconv.FormatFloat(secs, 'f', -1, 64))

}

func doUnaryPull(c sftppb.SFTPClient) {
	fmt.Println("Starting 2nd do uniary rpc client: ")
	startPulltime := time.Now()
	req := &sftppb.CopyFromRemoteRequest{
		Sftp: &sftppb.Sftp{
			Username: "foo",
			SystemId: "192.168.1.14",
			FileName: "test2.txt",
			HostKey:  "",
			PassWord: "123",
			HostPort: ":2225",
		},
	}
	res, err := c.CopyFromRemoteService(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	secs := time.Since(startPulltime).Seconds()
	log.Println("Response from SFTP: " + res.Result)
	fmt.Println("Function Time: " + strconv.FormatFloat(secs, 'f', -1, 64))

}
