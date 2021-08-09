# go-chat

This is a simple chat application build using golang net package. TCP connection is used for server and client communication.
Only two members can join the server and communciate with each other. The purpose of this application is to get familiar with some important libraries of golang for the future
projects.

# Usage


## cd server && go run server.go to run the server

Let's call it server terminal

Open two new terminals and type following commands:
## cd client && go run client.go [yourName]


You must enter name as arguments to client.go file, so that the system could identify you while broadcasting your message to the peers in a network. 



