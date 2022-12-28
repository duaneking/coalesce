package coalesce

import "testing"

func TestSimpleCoalesce(t *testing.T) {
	value := int(3)

	actual := Coalesce(nil, nil, &value)

	if *actual != int(3) {
		t.Fatal(actual)
	}
}

func TestSimpleNoNonNilCoalesce(t *testing.T) {

	actual := Coalesce[int](nil, nil)

	if actual != nil {
		t.Fatal(actual)
	}
}
