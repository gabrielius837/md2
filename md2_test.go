package md2

import (
	"bytes"
	"testing"
)

func TestMD2(t *testing.T) {
	str := "Hello, world!"
	input := []byte(str)

	expectedOutput := []byte{140, 202, 14, 150, 94, 221, 14, 34, 59, 116, 79, 156, 237, 248, 225, 65}

	output := MD2(input)

	if !bytes.Equal(output, expectedOutput) {
		t.Errorf("MD2(\"%s\") = %v, expected %v", str, output, expectedOutput)
	}
}
