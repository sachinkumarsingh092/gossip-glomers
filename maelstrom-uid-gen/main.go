package main

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"log"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "generate_ok"
		body["id"] = randomUUIDGen()

		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}

func randomUUIDGen() uint64 {
	timestamp := uint64(time.Now().UnixNano())

	var randomNum uint64
	binary.Read(rand.Reader, binary.NativeEndian, &randomNum)

	return timestamp ^ randomNum
}
