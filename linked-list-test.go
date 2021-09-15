package main

import (
	"testing"
	// 	"fmt"
	"os"
	// 	"strconv"
)

func TestLinkedListBasic(t *testing.T) {
	input := []byte("5")
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	_, err = w.Write(input)
	if err != nil {
		t.Error(err)
	}
	w.Close()

	stdin := os.Stdin
	// Restore stdin right after the test
	defer func() { os.Stdin = stdin }()
	os.Stdin = r

	// if err = main(); err != nil {
	// 	t.Fatalf("userInput: %v\n", err)
	// }
}
