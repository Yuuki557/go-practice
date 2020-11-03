package main

import (
	"testing"
)

func TestPractice(t *testing.T) {
	if Practice() != "5" {
		t.Fatalf("error")
	}
	t.Log("success")
}
