# Algorithme de comptage des nombres pairs et impairs

## Nom de l'algorithme

### Algorithme de comptage par parcours linéaire

- Classification binaire (pair/impair)
- Complexité : O(n) en temps, O(1) en espace
- Un seul passage dans le tableau

## Idée

Parcourir un tableau d'entier et afficher à la fin du parcours le nombre d'entiers pairs et le nombre d'entiers impairs contenus dans le tableau.

## Exemples d'utilisation

- **Télécommunications** : Vérifier si une transmission de données est correcte (bit de parité)
- **Cryptographie** : Analyse de la distribution des données
- **Statistiques** : Classification de données numériques
- **Contrôle qualité** : Validation de séquences de données

### Ressources utiles

- [Bit de parité - Wikipedia](https://fr.wikipedia.org/wiki/Bit_de_parit%C3%A9)

## Pseudocode - Algorithme de comptage des nombres pairs et impairs

FONCTION compterPairsImpairs(Tableau)
    countPair = 0        // Compteur de nombres pairs initialisé à 0
    countImpair = 0      // Compteur de nombres impairs initialisé à 0

    pour counter allant de 0 à (exclu) taille(Tableau) par pas de 1   
        si Tableau[counter] modulo 2 == 0 alors                         
            incrémenter le compteur de nombres pairs de 1                      
        sinon                                                      
            incrémenter le compteur de nombres impairs de 1
        fin si
    fin pour
    retourner compteur de nombres pairs, compteur de nombres impairs
FIN FONCTION

PROGRAMME PRINCIPAL
    // Création et remplissage du tableau
    donnees = créer un tableau de 20 entiers
    pour i allant de 0 à longueur(donnees) par pas de 1
        donnees[i] = nombre aléatoire entre 0 et 99
    fin pour

    // Affichage et traitement
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

## Complexité

- **Temps** : O(n) - parcours linéaire unique
- **Espace** : O(1) - utilisation de compteurs constants

## Utilisation

bash
go run main.go

## Exemple de sortie

=======★☆☆ : Compter les nombres pair et impairs====
Tableau : [42 17 88 3 56 91 12 7 34 65 20 83 48 1 76 29 14 95 60 37]
Nombre total de Pairs =  10
Nombre total d'Impairs =  10
Temps d'exécution de l'algorithme : 1.234µs
