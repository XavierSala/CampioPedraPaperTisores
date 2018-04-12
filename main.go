package main

import (
	"fmt"
	"sort"

	db "./pedrapapertisores/db"

	_ "github.com/lib/pq"
)

func main() {
	var connexio db.BaseDeDades

	resultat := make(map[string]int)

	_, err := connexio.Connecta()
	checkErr(err)

	p, err := connexio.CercaPrimer()
	checkErr(err)

	guanyador, err := p.Guanyador()
	checkErr(err)

	resultat[guanyador] = 1
	fmt.Println(p, " -> ", guanyador)

	var campio string

	for err == nil {
		p, err = connexio.CercaSeguentPartidaDelCampio(p)
		if p != nil {
			campio, err = p.Guanyador()
			checkErr(err)
			fmt.Println(p, " -> ", campio)

			valor, ok := resultat[campio]
			if !ok {
				resultat[campio] = 1
			} else {
				resultat[campio] = valor + 1
			}
		}
	}

	connexio.Desconnecta()

	ordena(resultat)
	fmt.Println("---------------------------")
	fmt.Println("CAMPIÓ: " + campio)
	fmt.Println("---------------------------")

}

type kv struct {
	Key   string
	Value int
}

func ordena(resultat map[string]int) {
	fmt.Println("---- Llista de campions ----- ")
	var ss []kv
	for k, v := range resultat {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	for _, kv := range ss {
		fmt.Printf("%s -> %d\n", kv.Key, kv.Value)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
