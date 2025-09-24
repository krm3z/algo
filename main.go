package main

import (
	"fmt"
	"math/rand"
	"time"
)

func compterPairsImpairs(Tableau []int) (int, int) {
    countPair := 0 // Compteur de nombres pairs intialisé à 0
    countImpair := 0 // Compteur de nombres impairs intialisé à 0

    for i := 0; i < len(Tableau); i++ { // pour counter allant de 0 à (exclu) taille (tableau) par pas de 1  
        if Tableau[i]%2 == 0 { //si tableau [counter] % 2 == 0 alors
            countPair++ //incrémenter le compteur de nombres pairs de 1
        } else { //sinon
            countImpair++ //incrémenter le compteur de nombres impairs de 1
        }
    }
    return countPair, countImpair //retourner compteur de nombres pairs, compteur de nombres impairs
}

func main() {
	donnees := make([]int, 20)
	for i := 0; i < len(donnees); i++ {
		donnees[i] = rand.Intn(100)
	}
	fmt.Println("=======★☆☆ : Compter les nombres pair et impairs====")
	fmt.Println("Tableau :", donnees)

	debut := time.Now()
	pairs, impairs := compterPairsImpairs(donnees)
	duree := time.Since(debut)

	fmt.Println("Nombre total de Pairs = ", pairs)
	fmt.Println("Nombre total d'Impairs = ", impairs)
	fmt.Println("Temps d'exécution   de l'algorithme :", duree)
}
