package utils

import (
	"fmt"
	"testing"
)

func TestGetPrefixPath(t *testing.T) {
	for i := 0; i <= 10; i++ {
		fmt.Println(GetPrefixPath())
	}

}
