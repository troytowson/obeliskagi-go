package obeliskagi

import (
	"strings"
	"testing"
)

func TestReadData(t *testing.T) {
	// Arrange
	id := "LINE1\nLINE2\nLINE3\nLINE4\nLINE5"
	r := strings.NewReader(id)

	// Act
	d, err := readData(r)

	// Assert
	if err != nil || id != d {
		t.Fail()
	}
}

func BenchmarkReadData(b *testing.B) {
	r := strings.NewReader("LINE1\nLINE2\nLINE3\nLINE4\nLINE5")
	for i := 0; i < b.N; i++ {
		readData(r)
	}
}
