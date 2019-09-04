### Architecture

It uses protobuf to communicate between the Client and Server. The server then commuicates via SFTP to the remote server.
The intention is to have a java api call the server to push or pull files as needed.  The server services is pretty fast.  Something on the order of 22-28 MB/sec (your speeds may vary).

### **Server**

To run the service you will need to cd into the agaveSFTP directory.  Then run "go run Server/server.go"  In the future it will be compiled to an executable exe.

this is an SFTP server that will connect to an SFTP service and download (or upload) the file specified.  For now that is all it does. 

This inputs are a user name, password, host and port on the host to connect to.

In the future it will use a key as an alternative to the user name and password for security.

### Client

There is a client to test the SFTP server service.  It is located in the Cleint directory.  To run it "go run Cleint/cleint.go"
this will push a 350mb file to and pull from the server.

### Docker

It uses a docker image to run the SFTP service on a different computer that can be used for testing pruposes.  
To run this change the username to the name that you want to use on your systems.   It is usually your login id.

"docker run \
    -v /Users/username/host/users.conf:/etc/sftp/users.conf:ro \
    -v /Users/username/share:/home \
    -p 2225:22 -d atmoz/sftp"

#### TODO

Modify for the use of keys.  May also want to expand to do more functions such as checking for the presence of a file or getting the file statistics, etc..