package test

import (
	"flag"
	"fmt"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	n := strings.SplitN("<basepath>/<groupname>/<key>", "/", 3)
	for _, v := range n {
		fmt.Println(v)
	}
	var post int
	flag.IntVar(&post, "post", 8080, "My-cache server port")
	fmt.Println(post)
}
