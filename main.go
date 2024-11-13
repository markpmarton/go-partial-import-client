package main

import (
	"fmt"

	"github.com/markpmarton/go-partial-import-poc/pkg/proto/sampleproto"
)

func main() {
	event := sampleproto.SampleEvent{}
	event.Content = []byte{byte(42)}
	fmt.Println(event)
}
