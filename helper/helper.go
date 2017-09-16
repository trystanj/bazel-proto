package helper

import (
	"log"

	"github.com/trystanj/bazel-proto/proto"
)

func Help() string {
	d := &proto.Demo{Very: "hi", Useful: 1}
	log.Printf("demo.Very, %v", d.Very)

	return d.Very
}
