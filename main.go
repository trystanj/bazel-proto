package main

import "log"
import "github.com/trystanj/bazel-proto/helper"

func main() {
	log.Info("OH HAI")

	helper.Help()
	d := &demo.Demo{very: "hi", useful: 1}
	log.Infof("demo.very, %v", d.very)
}
