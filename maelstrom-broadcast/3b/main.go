package main

import (
	"encoding/json"
	"log"
	"sync"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

type msServer struct {
	mu     sync.RWMutex
	msNode *maelstrom.Node
	set    map[int]struct{}
}

func main() {
	n := maelstrom.NewNode()
	set := make(map[int]struct{})
	ms := msServer{
		msNode: n,
		set:    set,
	}

	n.Handle("broadcast", ms.broadcastHandler)
	n.Handle("read", ms.readHandler)
	n.Handle("topology", ms.topologyHandler)

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}

func (ms *msServer) broadcastHandler(msg maelstrom.Message) error {
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	ms.mu.Lock()
	message := int(body["message"].(float64))
	if _, exists := ms.set[message]; exists {
		ms.mu.Unlock()
		return nil
	} else {
		// insert into set
		ms.set[message] = struct{}{}
	}
	ms.mu.Unlock()

	for _, node := range ms.msNode.NodeIDs() {
		if node == ms.msNode.ID() || node == msg.Src {
			continue
		}

		node := node
		go func() {
			if err := ms.msNode.Send(node, body); err != nil {
				panic(err)
			}
		}()
	}

	return ms.msNode.Reply(msg, map[string]any{"type": "broadcast_ok"})
}

func (ms *msServer) readHandler(msg maelstrom.Message) error {
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	ms.mu.RLock()
	body["type"] = "read_ok"
	messages := make([]int, 0, len(ms.set))
	for key := range ms.set {
		messages = append(messages, key)
	}
	body["messages"] = messages
	ms.mu.RUnlock()

	return ms.msNode.Reply(msg, body)
}

func (ms *msServer) topologyHandler(msg maelstrom.Message) error {
	return ms.msNode.Reply(msg, map[string]any{"type": "topology_ok"})
}
