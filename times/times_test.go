package times

import (
	"fmt"
	"testing"
	"time"
)

func TestBeginOfYear(t *testing.T) {
	fmt.Println(OffsetMonthTime(time.Now(), -1))
}
