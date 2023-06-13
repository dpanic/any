package convert

import (
	"testing"
)

func TestComposite(t *testing.T) {
	test1 := []int{1, 2, 3, 4, 5}
	res1 := ToSliceOfFloat(test1)
	if len(res1) == 0 {
		t.Errorf("Result was incorrect, expected length of slice bigger than 0")
	}

	test2 := []int{1, 2, 3, 4, 5}
	res2 := ToSliceOfString(test2)
	if len(res2) == 0 {
		t.Errorf("Result was incorrect, expected length of slice bigger than 0")
	}

	test8 := map[string]string{
		"key": "val",
	}
	res8 := ToMapOfStrings(test8)
	if _, ok := res8["key"]; ok == false {
		t.Errorf("Result was incorrect, existance of key 'key'")
	}

	type TestStructure struct{
		property1 string
		property2 string
	}

	test9 := map[string]TestStructure{
		"key": {
			property1: "val1",
			property2: "val2",
		},
	}
	res9 := ToMapOfStruct[TestStructure](test9)
	if _, ok := res9["key"]; ok == false {
		t.Errorf("Testing ToMapOfStruct result was incorrect, existance of key 'key'")
	}
}
