package base62

import (
	"testing"
)

func TestEncode(t *testing.T) {
	if encoded := Encode(89569285645); encoded != "1ZlfarV" {
		t.Fail()
	}
}