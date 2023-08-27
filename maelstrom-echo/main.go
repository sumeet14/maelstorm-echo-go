package main

import (
	"encoding/json"
	"log"
	"os"

	maelstorm "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstorm.NewNode()

	//register a handler for the "echo" message that responds with a "echo_ok"
	n.Handle("echo", func(msg maelstorm.Message) error {
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		body["type"] = "echo_ok"

		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Printf("ERROR: %s", err)
		os.Exit(1)
	}
}
