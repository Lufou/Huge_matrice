package main

import (
	"bufio"
	"fmt"
	"os"
)

var TAILLE_MATRICE_A int = 10000
var TAILLE_MATRICE_B int = TAILLE_MATRICE_A
var matA [][]int
var matB [][]int

func main() {
	matA := make([][]int, TAILLE_MATRICE_A)
	matB := make([][]int, TAILLE_MATRICE_B)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Nombre de lignes de la matrice A")
	scanner.Scan()
	fmt.Print("Nombre de colonnes de la matrice A")
	scanner.Scan()
	fmt.Print("Nombre de lignes de la matrice B")
	scanner.Scan()
	fmt.Print("Nombre de colonnes de la matrice B")
	scanner.Scan()

	//fonction pour multiplier les matrices, basique
	// après, la même avec des coroutines

}

//var rowsmatA, columnsmatA int = 2,3
//var rowsmatB, columnsmatB int = 4,5
func possibleProduct(rA int, cA int, rB int, cB int) (check bool) {
	if cA != rB {
		fmt.Print("Multiplication de matrices impossible")
		return false
	} else {
		fmt.Print("Multiplication de matrices possible")
		return true
	}
}
