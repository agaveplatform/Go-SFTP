package main

import (
	"../SftpPb"
	"bufio"
	"context"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type server struct{}

func main() {
	fmt.Println("Hello World")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	sftppb.RegisterSFTPServer(s, &server{})

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (*server) CopyLocalToRemoteService(ctx context.Context, req *sftppb.CopyLocalToRemoteRequest) (*sftppb.CopyLocalToRemoteResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	fileName := req.Sftp.FileName
	fmt.Println("\n" + fileName)
	passWord := req.Sftp.PassWord
	fmt.Println(passWord)
	systemId := req.Sftp.SystemId
	fmt.Println(systemId)
	username := req.Sftp.Username
	fmt.Println(username)
	hostKey := req.Sftp.HostKey
	fmt.Println(hostKey)
	port := ":22"
	// get host public key
	HostKey := getHostKey(systemId)

	config := ssh.ClientConfig{
		User: username,
		Auth: []ssh.AuthMethod{
			ssh.Password(passWord),
		},
		// HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		HostKeyCallback: ssh.FixedHostKey(HostKey),
	}

	result := ""
	// connect
	conn, err := ssh.Dial("tcp", systemId+port, &config)
	if err != nil {
		log.Println(err)
		result = err.Error()
	}
	defer conn.Close()

	// create new SFTP client
	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Println(err)
		result = err.Error()
	}
	defer client.Close()

	// create destination file
	dstFile, err := client.Create("./file.txt")
	if err != nil {
		log.Println(err)
		result = err.Error()
	}
	defer dstFile.Close()

	// create source file
	srcFile, err := os.Open("./file.txt")
	if err != nil {
		log.Println(err)
		result = err.Error()
	}

	// copy source file to destination file
	bytes, err := io.Copy(dstFile, srcFile)
	if err != nil {
		log.Println(err)
		result = err.Error()
	}
	fmt.Printf("%d bytes copied\n", bytes)

	res := &sftppb.CopyLocalToRemoteResponse{
		Result: result,
	}
	fmt.Println(res.String())
	return res, nil
}

func getHostKey(host string) ssh.PublicKey {
	// parse OpenSSH known_hosts file
	// ssh or use ssh-keyscan to get initial key
	fmt.Println(host)
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				log.Fatalf("error parsing %q: %v", fields[2], err)
			}
			break
		}
	}

	if hostKey == nil {
		log.Fatalf("no hostkey found for %s", host)
	}

	return hostKey
}
