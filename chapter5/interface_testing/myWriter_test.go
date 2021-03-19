package myWriter

import (
	"io"
	"testing"
)

/*
Test: Check does MyWriter struct satisfy implemetation io.Writer interface
*/
func TestWriter(t *testing.T) {
	var _ io.Writer = &MyWriter{}
}

/*
Output:
*MyWriter does not implement io.Writer (wrong type for Write method)
		have Write([]byte) error
		want Write([]byte) (int, error)
*/
