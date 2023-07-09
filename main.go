package main

import (
	"fmt"
	"github.com/Silverhaed/CityMesh/peer"
)

func main() {
	fmt.Println("Welcome to CityMesh!")
	router := peer.NewRouter()
	node1 := peer.NewNode(router)
	node2 := peer.NewNode(router)

	node1.SendMessage("Hello from Node 1!", node2)
	node2.SendMessage("Hi there! This is Node 2.", node1)
}
