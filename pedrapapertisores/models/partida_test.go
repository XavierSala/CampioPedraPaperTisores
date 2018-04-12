package models

import "testing"

var tiradesTest = []struct {
	jugador1 string // input
	jugador2 string // input
	expected bool   // expected result
}{
	{"pedra", "pedra", false},
	{"pedra", "paper", false},
	{"pedra", "tisores", true},
	{"tisores", "tisores", false},
	{"tisores", "paper", true},
	{"tisores", "pedra", false},
	{"paper", "paper", false},
	{"paper", "pedra", true},
	{"paper", "tisores", false},
}

func TestComprovaSiGuanyaFunciona(t *testing.T) {
	for _, tt := range tiradesTest {
		var p Partida
		resultat := p.guanya(tt.jugador1, tt.jugador2)
		if resultat != tt.expected {
			t.Errorf("guanya (%s, %s): expected %t actual %t", tt.jugador1, tt.jugador2, tt.expected, resultat)
		}
	}
}
