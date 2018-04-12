package db

import (
	"database/sql"
	"errors"
	"fmt"

	m "../models"
	// Driver Postgresql
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "ies2010"
	dbname   = "pedra"
)

const cercaPrimer = `SELECT p.id_partida,
p.id_jugador1, j1.nom as nom1,
p.id_jugador2, j2.nom as nom2
FROM partida p
INNER JOIN jugador j1 ON p.id_jugador1 = j1.id_jugador
INNER JOIN jugador j2 ON p.id_jugador2 = j2.id_jugador
WHERE p.id_partida = 1`

const cercaAltres = `SELECT p.id_partida,
p.id_jugador1, j1.nom as nom1,
p.id_jugador2, j2.nom as nom2
FROM partida p
INNER JOIN jugador j1 ON p.id_jugador1 = j1.id_jugador
INNER JOIN jugador j2 ON p.id_jugador2 = j2.id_jugador
WHERE p.id_partida > $1
AND (j1.nom = $2 OR j2.nom = $2)
ORDER BY p.id_partida
LIMIT 1`

const cercaJugades = `SELECT num_jugada, tira_jugador1, tira_jugador2
FROM jugada
WHERE id_partida=$1`

// BaseDeDades és la interfície bàsica
type BaseDeDades struct {
	db *sql.DB
}

// Connecta amb la base de dades
func (b *BaseDeDades) Connecta() (bool, error) {
	var err error

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	b.db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return false, err
	}

	// Es fa servir Ping perquè Open no crea la connexió, només en
	// comprova els paràmetres
	err = b.db.Ping()
	if err != nil {
		panic(err)
	}

	return true, nil
}

// Desconnecta de la base de dades
func (b *BaseDeDades) Desconnecta() {
	b.db.Close()
}

// CercaPrimer busca el primer registre de la base de dades
func (b *BaseDeDades) CercaPrimer() (*m.Partida, error) {
	var partida m.Partida
	row := b.db.QueryRow(cercaPrimer)
	err := row.Scan(&partida.ID, &partida.Idjugador1, &partida.Jugador1,
		&partida.Idjugador2, &partida.Jugador2)

	if err == sql.ErrNoRows {
		return nil, errors.New("no rows were returned")
	} else if err != nil {
		return nil, err
	}

	partida.Jugades, err = b.recoverJugades(1)
	if err != nil {
		return nil, err
	}
	return &partida, nil
}

// CercaSeguentPartidaDelCampio busca el primer registre de la base de dades
func (b *BaseDeDades) CercaSeguentPartidaDelCampio(vella *m.Partida) (*m.Partida, error) {
	var partida m.Partida
	nomDelCampio, err := vella.Guanyador()
	if err != nil {
		panic(err)
	}
	row := b.db.QueryRow(cercaAltres, vella.ID, nomDelCampio)
	err = row.Scan(&partida.ID, &partida.Idjugador1, &partida.Jugador1,
		&partida.Idjugador2, &partida.Jugador2)

	if err == sql.ErrNoRows {
		return nil, errors.New("no rows were returned")
	} else if err != nil {
		return nil, err
	}

	partida.Jugades, err = b.recoverJugades(partida.ID)
	if err != nil {
		return nil, err
	}
	return &partida, nil
}

func (b *BaseDeDades) recoverJugades(id int) ([]m.Jugada, error) {
	var jugades []m.Jugada
	rows, err := b.db.Query(cercaJugades, id)
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var jugada m.Jugada
		err = rows.Scan(&jugada.Numero, &jugada.Juga1, &jugada.Juga2)
		if err != nil {
			// handle this error
			return nil, err
		}
		jugades = append(jugades, jugada)
	}
	return jugades, nil
}
