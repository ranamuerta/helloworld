package main

import (
	"testing"
)

func TestMsg(t *testing.T){
	if hw() != "Hello world"{
		t.Fatal("Wrong hello world message")
	}
}