package main

import (
	"fmt"
	"sync"

	emitter "github.com/emitter-io/go"
)

const channelKey = ""

func main() {
	var wg sync.WaitGroup
	opts := emitter.NewClientOptions()
	opts.AddBroker("tcp://127.0.0.1:8080")
	opts.SetOnMessageHandler(onMessage)

	client := emitter.NewClient(opts)
	// t := client.Connect()
	wg.Add(1)
	wait(client.Connect())
	wait(client.Subscribe(channelKey, "demo/"))
	wait(client.Publish(channelKey, "demo/", "hello, emitter!"))

	// for {

	// }
	wg.Done()
}

func wait(t emitter.Token) {
	t.Wait()
	if t.Error() != nil {
		panic(t.Error())
	}
}

func onMessage(client emitter.Emitter, msg emitter.Message) {
	fmt.Printf("message: %v %v\n", msg.Topic(), msg.Payload())

}
