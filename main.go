package main

import "log"
import "github.com/trystanj/bazel-proto/proto/demo"

func main() {
	log.Info("OH HAI")

	d := &demo.Demo{very: "hi", useful: 1}
	log.Infof("demo.very, %v", d.very)
}
