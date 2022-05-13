package main

import "testing"

func TestOla(t *testing.T) {
	resultado := Ola("João")
	esperado := "Olá, João"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s'",resultado, esperado)
	}
}