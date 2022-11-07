package test

import (
	"fmt"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	n := strings.SplitN("<basepath>/<groupname>/<key>", "/", 3)
	for _, v := range n {
		fmt.Println(v)
	}

}
