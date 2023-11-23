package gen

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	var id uint64 = 3
	code := Encode(id)
	fmt.Println(code)
	fmt.Println(Decode(code))
}
