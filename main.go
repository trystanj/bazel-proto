package main

import "log"
import "github.com/trystanj/bazel-proto/helper"

func main() {
	log.Println("OH HAI")

	help := helper.Help()
	log.Printf("Helper: %v", help)
}
