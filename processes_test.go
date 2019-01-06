package pwn

import (
	"bytes"
	"testing"
	"time"
)

func TestEcho(t *testing.T) {
	expected := []byte("Hello, world!")
	p, err := Spawn("echo", "Hello, world!")
	if err != nil {
		t.Fatal(err)
	}

	output, err := p.ReadLine(time.Second)
	if !bytes.Equal(output, expected) {
		t.Fatalf("%q != %q", output, expected)
	}
}

func TestSh(t *testing.T) {
	expected := []byte("Hello, world")
	p, err := Spawn("sh")
	if err != nil {
		t.Fatal(err)
	}

	err = p.WriteLine("echo Hello, world")
	if err != nil {
		t.Fatal(err)
	}

	out, err := p.ReadLine(time.Second)
	if !bytes.Equal(out, expected) {
		t.Fatalf("%q != %q", out, expected)
	}
}
