package mypkgs_test

import (
	"testing"

	"mypkgs"
)

func TestHello(t *testing.T) {
	got := mypkgs.Hello()
	if got != "Hello" {
		t.Errorf("Hello() want Hello, got = %s", got)
	}
}

func TestNG(t *testing.T) {
	if mypkgs.NG {
		t.Errorf("NG want false, got = true")
	}
}

func TestDeny(t *testing.T) {
	if mypkgs.Deny() {
		t.Errorf("Deny() want false, got = true")
	}
}
