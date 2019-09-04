
Server
To run the service you will need to cd into the agaveSFTP directory.  Then run "go run Server/server.go"  In the future it will be compiled to an executable exe.

this is an SFTP server that will connect to an SFTP service and download (or upload) the file specified.  For now that is all it does. 

This inputs are a user name, password, host and port on the host to connect to.

In the future it will use a key as an alternative to the user name and password for security.

Client
There is a client to test the SFTP server.  It is located in the Cleint directory.  To run it "go run Cleint/cleint.go"
this will push a 350mb file to and pull from the server.

It uses a docker image.  To run this change the username to the name that you want to use on your systems.   It is usually your login id.

"docker run \
    -v /Users/username/host/users.conf:/etc/sftp/users.conf:ro \
    -v /Users/username/share:/home \
    -p 2225:22 -d atmoz/sftp"

