package main

import "log"
import _ "github.com/trystanj/bazel-proto/proto/demo"

func main() {
	log.Info("OH HAI")

	demo := &demo.VeryUseful{very: "hi", useful: 1}
	log.Infof("demo.very, %v", demo.very)
}
