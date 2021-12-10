package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var result [][]int
var matA [][]int
var matB [][]int

const LARGEUR_MATRICES = 2000 // initialisation de la largeur de la matrice
const HAUTEUR_MATRICES = 2000 // initialisationd de la hauteur de la matrice
const inc = 200

// Créer queue ?

func remplirMatrices() {
	fmt.Printf("#DEBUG START remplirMatrices\n") // debug
	matA = make([][]int, HAUTEUR_MATRICES) // déclaration du tableau contenant le nombre de lignes de la matrice A
	for i := 0; i < HAUTEUR_MATRICES; i++ { // parcours de la matrice A par lignes
		matA[i] = make([]int, LARGEUR_MATRICES) // pour chaque ligne i, on génère un tableau contenant [LARGEUR_MATRICES] colonnes
		for j := 0; j < LARGEUR_MATRICES; j++ { // parcours de la matrice par colonnes
			matA[i][j] = rand.Intn(10000) // Default: 10000
			// remplissage de l'élément [i][j] de matA par une valeur aléatoire comprise entre 0 et 10000
		}
	}

	matB = make([][]int, HAUTEUR_MATRICES) // declaration du tableau contenant le nombre de lignes de la matrice B
	for i := 0; i < HAUTEUR_MATRICES; i++ { // parcours de matB par lignes
		matB[i] = make([]int, LARGEUR_MATRICES) // pour chaque ligne i, on génère un tableau contenant [LARGEUR_MATRICES] colonnes
		for j := 0; j < LARGEUR_MATRICES; j++ { // parcours de matB par colonnes
			matB[i][j] = rand.Intn(10000) // Default: 10000
			// remplissage de l'élément [i][j] de matB par une valeur aléatoire comprise entre 0 et 10000
		}
	}

	fmt.Printf("\n#DEBUG END remplirMatrices\n") // debug
}

func printMat(mat [][]int) {
	res := ""
	res += "\n\nMatrice\n"
	for i := 0; i < HAUTEUR_MATRICES; i++ { // parcours de la matrice à afficher par lignes
		for j := 0; j < LARGEUR_MATRICES; j++ { // parcours de la matrice à afficher par colonnes
			res += fmt.Sprintf("%d ", mat[i][j]) // stocke les valeurs des coefs de la matrice dans la variable res
		}
		res += "\n" // retour à la ligne 
	}
	fmt.Printf("%s", res) // affichage de la matrice sous forme de string
}

func printMatLine(mat [][]int, from int, to int) { // affiche les lignes de la from-ième à la to-ième
	for line := from; line < to; line++ { // parcours de la ligne for à la ligne to
		res := "" 
		if line == 0 {
			res += "\n\nMatrice\n" // si c'est la première ligne de la matrice, on affiche juste "Matrice"
		}
		res += fmt.Sprintf("\nLine %d\n", line) // stocke toutes les valeurs d'une ligne dans la variable res
		for j := 0; j < LARGEUR_MATRICES; j++ { // parcours de la ligne courante
			res += fmt.Sprintf("%d ", mat[line][j]) // stocke les valeurs de la ligne dans la variable res
		}
		res += "\n"
		fmt.Printf("%s", res) // affiche toutes les lignes sous forme de string

	}
	wg.Done() // retire un token de la file d'attente

}

func main() {
	if inc == 0 || inc > 500 {
		fmt.Print(":^)")
		return
	}
	fmt.Printf("#DEBUG START\n") // debug
	timeStart := time.Now() // stockage de l'heure à laquelle le programme se lance dans la variable timeStart

	result = make([][]int, HAUTEUR_MATRICES) // déclaration de la matrice résultat
	remplirMatrices()

	fmt.Printf("#DEBUG START GOROUTINES\n") // debug
	for i := 0; i < HAUTEUR_MATRICES; i += inc { 
		wg.Add(1) // ajout d'un token
		go multiplicationByLine(i, i+inc-1, matA, matB) // lancement des goroutines qui effectuent le calcul
	}

	wg.Wait() // on attend ici que le nombre de tokens soit nul
	fmt.Printf("#DEBUG END GOROUTINES\n") // debug
	fmt.Printf("#DEBUG START GOROUTINES PRINTLINES\n") // debug

	for i := 0; i < HAUTEUR_MATRICES; i += inc { // parcours des lignes de la matrice de 0 à inc 
		//si on a initialisé inc à 200, la goroutine va afficher les lignes de la matrice 200 par 200
		wg.Add(1) // ajout d'un token
		go printMatLine(result, i, i+inc) // lancement des goroutines qui affichent les lignes
		time.Sleep(time.Millisecond * inc) //Cette ligne nous permet d'avoir la matrice affichée dans l'ordre
	}

	wg.Wait() // on attend ici que le nombre de tokens soit nul
	fmt.Printf("#DEBUG ALL PRINTLINES GOROUTINES ENDED\n") // debug
	timeEnd := time.Now() // stockage de l'heure à laquelle le programme se termine dans la variable timeEnd
	elapsed := timeEnd.Sub(timeStart) // calcul du temps d'exécution du programme
	fmt.Printf("Time elapsed : %d", elapsed.Milliseconds()) // affichage du temps d'exécution du programme
}

func multiplicationByLine(from int, to int, matA [][]int, matB [][]int) {
	for line := from; line <= to; line++ { // parcours des lignes de la matrice résultat
		result[line] = make([]int, LARGEUR_MATRICES) // déclaration du tableau stockant une ligne de résultats
		for j := 0; j < LARGEUR_MATRICES; j++ { // parcours des colonnes de la matrice
			for l := 0; l < HAUTEUR_MATRICES; l++ { // parcours des lignes de la matrice
				result[line][j] = result[line][j] + matA[line][l]*matB[l][j] // calcul du coefficient à la j-eme colonne de la ligne en cours de calcul
			}
		}
	}
	wg.Done()
}

}
