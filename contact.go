package kademlia

import "net"

type Contact struct {
	ID   NodeID
	Addr net.IP
	Port uint16
}

func NewContact(id NodeID, addr net.IP) Contact {
	return Contact{
		ID:   id,
		Addr: addr,
	}
}

type Contacts []Contact

func (c Contacts) Len() int           { return len(c) }
func (c Contacts) Less(i, j int) bool { return c[i].ID.Less(c[j].ID) }
func (c Contacts) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
