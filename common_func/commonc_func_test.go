package common_func

import "testing"

func TestSplitStringToArray(t *testing.T) {
	str := "1,2,3"

	array, err := SplitStringToUint64Array(str)

	if err != nil {
		t.Error(err)
	}

	t.Log(array)
}
