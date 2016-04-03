package obeliskagi

import (
	"fmt"
	"testing"
)

func TestNewReplyWithResponseCodeAndResult(t *testing.T) {
	// Arrange
	rc := 200
	res := 0
	s := fmt.Sprintf("%d result=%d", rc, res)

	// Act
	r := newReply(s)

	// Assert
	if r.Line != s || r.Status != rc || r.Result != res {
		t.Fail()
	}
}

func TestNewReplyWithResponseCodeResultAndData(t *testing.T) {
	// Arrange
	rc := 200
	res := 0
	d := "Test Data"
	s := fmt.Sprintf("%d result=%d (%s)", rc, res, d)

	// Act
	r := newReply(s)

	// Assert
	if r.Line != s || r.Status != rc || r.Result != res || r.Data != d {
		t.Fail()
	}
}

func BenchmarkNewReply(b *testing.B) {
	for i := 0; i < b.N; i++ {
		newReply("200 result=0")
	}
}
