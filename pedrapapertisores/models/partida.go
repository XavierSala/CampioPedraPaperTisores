package models

import "errors"

var tirada = []string{"pedra", "paper", "tisores"}

const tirades = 3

// Partida són les partides que s'han fet
type Partida struct {
	ID         int
	Idjugador1 string
	Jugador1   string
	Idjugador2 string
	Jugador2   string
	Jugades    []Jugada
}

// Guanyador determina qui és el que guanya
func (p Partida) Guanyador() (string, error) {
	max := len(p.Jugades) - 1
	if p.Jugades[max].Juga1 == p.Jugades[max].Juga2 {
		return "", errors.New("Partida incorrecta: és un empat")
	}
	if p.guanya(p.Jugades[max].Juga1, p.Jugades[max].Juga2) {
		return p.Jugador1, nil
	}
	return p.Jugador2, nil
}

// guanya determina si ha guanyat un jugador o un altre
func (p Partida) guanya(jugador1 string, jugador2 string) bool {

	if jugador1 == jugador2 {
		return false
	}

	hajugat1 := indexOf(jugador1, tirada)
	haJugat2 := indexOf(jugador2, tirada)

	if haJugat2 == (hajugat1+1)%tirades {
		return false
	}
	return true
}

func indexOf(word string, data []string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}
