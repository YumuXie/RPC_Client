package main

import (
	"flag"
	"net/rpc"
	"uk.ac.bris.cs/distributed2/secretstrings/stubs"

	//	"bufio"
	//	"os"
	"fmt"
)

func main() {
	server := flag.String("server", "54.92.168.157:8030", "IP:port string to connect to as server")
	flag.Parse()
	fmt.Println("Server: ", *server)
	client, _ := rpc.Dial("tcp", *server)
	defer client.Close()

	request := stubs.Request{Message: "Hello"}
	response := new(stubs.Response)
	client.Call(stubs.PremiumReverseHandler, request, response)
	fmt.Println("Responded: " + response.Message)
	//TODO: connect to the RPC server and send the request(s)
}
