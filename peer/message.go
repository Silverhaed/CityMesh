package peer

type Message struct {
	Sender    *Node
	Recipient *Node
	Content   string
}

func NewMessage(sender *Node, recipient *Node, content string) *Message {
	return &Message{
		Sender:    sender,
		Recipient: recipient,
		Content:   content,
	}
}
