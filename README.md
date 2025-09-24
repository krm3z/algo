# algo

## Pseudocode - Algorithme de comptage des nombres pairs et impairs

FONCTION compterPairsImpairs(Tableau)
    countPair = 0
    countImpair = 0

    pour counter allant de 0 à (exclu) taille(Tableau) par pas de 1   
        si Tableau[i] modulo 2 == 0 alors                         
            incrémenter le compteur de nombres pairs de 1                      
        sinon                                                      
            incrémenter le compteur de nombres impairs de 1
        fin si
    fin pour
    retourner compteur de nombres pairs, compteur de nombres impairs
FIN FONCTION

PROGRAMME PRINCIPAL
    donnees = tableau de 20 entiers
    pour i allant de 0 à longueur(donnees) par pas de 1
        donnees[i] = nombre aléatoire entre 0 et 99
    fin pour

    afficher "=======★☆☆ : Compter les nombres pair et impairs===="
    afficher "Tableau :", donnees
    
    debut = temps actuel
    pairs, impairs = compterPairsImpairs(donnees)
    duree = temps actuel - debut
    
    afficher "Nombre total de Pairs = ", pairs
    afficher "Nombre total d'Impairs = ", impairs
    afficher "Temps d'exécution de l'algorithme :", duree
FIN PROGRAMME

## Description

Cet algorithme compte le nombre d'entiers pairs et impairs dans un tableau de 20 nombres aléatoires et mesure le temps d'exécution.
