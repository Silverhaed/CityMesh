// package main

// import (
// 	"fmt"
// 	"github.com/Silverhaed/CityMesh/peer"
// )

// func main() {
// 	fmt.Println("Welcome to CityMesh!")
// 	router := peer.NewRouter()
// 	node1 := peer.NewNode(router)
// 	node2 := peer.NewNode(router)

// 	node1.SendMessage("Hello from Node 1!", node2)
// 	node2.SendMessage("Hi there! This is Node 2.", node1)
// }

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"
)

type Node struct {
	Connections map[string]bool
	Address     Address
}

type Address struct {
	IPv4 string
	Port string
}

type Package struct {
	To   string
	From string
	Data string
}

func init() {
	if len(os.Args) != 2 {
		panic("len args != 2")
	}
}

func main() {
	NewNode(os.Args[1]).Run(handleServer, handleClient)
}

// ipv4:port
// 127.0.0.1:8080
func NewNode(address string) *Node {
	splited := strings.Split(address, ":")

	if len(splited) != 2 {
		return nil
	}

	return &Node{
		Connections: make(map[string]bool),
		Address: Address{
			IPv4: splited[0],
			Port: ":" + splited[1],
		},
	}

}

func (node *Node) Run(handleServer, handleClient func(*Node)) {
	go handleServer(node)
	handleClient(node)
}

func handleServer(node *Node) {
	listen, err := net.Listen("tcp", "0.0.0.0"+node.Address.Port)
	if err != nil {
		panic("Server: listen error")
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			break
		}
		go handleConnection(node, conn)
	}
}

func handleConnection(node *Node, conn net.Conn) {
	defer conn.Close()
	var (
		buffer  = make([]byte, 512)
		message string
		pack    Package
	)
	for {
		length, err := conn.Read(buffer)
		if err != nil {
			break
		}
		message += string(buffer[:length])
	}
	err := json.Unmarshal([]byte(message), &pack)
	if err != nil {
		return
	}
	node.ConnectTo([]string{pack.From})
	fmt.Println(pack.Data)
}

func handleClient(node *Node) {
	for {
		message := InputString()
		splited := strings.Split(message, " ")
		switch splited[0] {
		case "/exit":
			fmt.Println("handle exit")
			os.Exit(0)
		case "/connect":
			node.ConnectTo(splited[1:])
		case "/nerwork":
			node.PrintNerwork()
		default:
			node.SendMessageToAll(message)
		}
	}
}

func (node *Node) PrintNerwork() {
	for conn := range node.Connections {
		fmt.Println("|", conn)
	}
}

func (node *Node) ConnectTo(addresses []string) {
	for _, addr := range addresses {
		node.Connections[addr] = true
	}
}

func (node *Node) SendMessageToAll(message string) {
	var new_pack = &Package{
		From: node.Address.IPv4 + node.Address.Port,
		Data: message,
	}
	for addr := range node.Connections {
		new_pack.To = addr
		node.Send(new_pack)
	}
}

func (node *Node) Send(pack *Package) {
	conn, err := net.Dial("tcp", pack.To)
	if err != nil {
		delete(node.Connections, pack.To)
		return
	}
	defer conn.Close()
	json_pack, _ := json.Marshal(*pack)
	conn.Write(json_pack)
}

func InputString() string {
	msg, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.Replace(msg, "\n", "", -1)
}
