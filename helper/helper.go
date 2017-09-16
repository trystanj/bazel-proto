package helper

import "log"
import "github.com/trystanj/bazel-proto/proto/demo"

func Help() string {
	d := &demo.Demo{very: "hi", useful: 1}
	log.Infof("demo.very, %v", d.very)
	return d.very
}
