package main

import (
	"fmt"
	"math/rand"
	"time"
)

func compterPairsImpairs(Tableau []int) (int, int) {
	countpair := 0   // Compteur de nombres pairs intialisé à 0
	countimpair := 0 // Compteur de nombres impairs intialisé à 0

	for i := 0; i < len(Tableau); i++ { // pour counter allant de 0 à (exclu) taille (tableau) par pas de 1
		if Tableau[i]%2 == 0 { //si tableau [counter] % 2 == 0 alors
			countpair++ //incrémenter le compteur de nombres pairs de 1
		} else { //sinon
			countimpair++ //incrémenter le compteur de nombres impairs de 1
		}
	}
	return countpair, countimpair //retourner compteur de nombres pairs, compteur de nombres impairs
}
func main() {
	tailles := []int{10,100, 1000, 10000, 100000, 1000000, 10000000, 100000000}

	for _, n := range tailles {
		donnees := make([]int, n)
		for i := 0; i < n; i++ {
			donnees[i] = rand.Intn(10000)
		}

		debut := time.Now()
		nbrpairs, nbrimpairs := compterPairsImpairs(donnees)
		duree := time.Since(debut)

		fmt.Printf("Taille: %d  → Temps: %v\n", n, duree)
		fmt.Println("Nombre total de Pairs = ", nbrpairs)
		fmt.Println("Nombre total d'Impairs = ", nbrimpairs)
		fmt.Println("--------------------------------------")
	}
}
