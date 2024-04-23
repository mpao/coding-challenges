package main

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var file []byte

func init() {
	var err error
	if file, err = os.ReadFile("test.txt"); err != nil {
		log.Fatal(err)
	}
}

func TestMessage(t *testing.T) {

}

func TestByteCount(t *testing.T) {
	want := 342190
	got := byteCount(file)
	assert.Equal(t, want, got)
}

func TestLineCount(t *testing.T) {
	want := 7145
	got := lineCount(file)
	assert.Equal(t, want, got)
}

func TestWordCount(t *testing.T) {
	want := 58164
	got := wordCount(file)
	assert.Equal(t, want, got)
}

func TestMultibyteCount(t *testing.T) {
	want := 339292
	got := multibyteCount(file)
	assert.Equal(t, want, got)
}
