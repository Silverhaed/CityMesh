package peer

import (
	"fmt"
)

type Router struct {
	// List of connected nodes
	nodes []*Node
}

func NewRouter() *Router {
	return &Router{
		nodes: []*Node{},
	}
}

func (r *Router) AddNode(node *Node) {
	r.nodes = append(r.nodes, node)
}

func (r *Router) RouteMessage(message *Message, sender *Node, recipient *Node) {
	fmt.Printf("[Router] Routing message from Node %d to Node %d\n", sender.ID, recipient.ID)
	recipient.ReceiveMessage(message)
}
