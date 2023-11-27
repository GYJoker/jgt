package encrypt

import (
	"fmt"
	"testing"
)

func TestGetUuid(t *testing.T) {
	uuid := GetUuid()
	fmt.Println(uuid)
}
