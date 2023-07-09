package peer

import "fmt"

func NewNode(router *Router) *Node {
	node := &Node{
		ID:     len(router.nodes) + 1,
		router: router,
	}
	router.AddNode(node)
	return node
}

type Node struct {
	ID     int
	router *Router
}

func (n *Node) SendMessage(content string, recipient *Node) {
	message := NewMessage(n, recipient, content)
	n.router.RouteMessage(message, n, recipient)
}

func (n *Node) ReceiveMessage(message *Message) {
	fmt.Printf("[Node %d] Received message: %s\n", n.ID, message.Content)
}
