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

<<<<<<< HEAD
const LARGEUR_MATRICES = 1000
const HAUTEUR_MATRICES = 1000

const inc = 100
=======
const LARGEUR_MATRICES = 500
const HAUTEUR_MATRICES = 500
const inc = 50
>>>>>>> 8040cb6ab0fd2c8d800076a00795f5fd9ee30fd4

// Créer queue ?

func remplirMatrices() {
	fmt.Printf("#DEBUG START remplirMatrices\n")
	matA = make([][]int, HAUTEUR_MATRICES)
	for i := 0; i < HAUTEUR_MATRICES; i++ {
		matA[i] = make([]int, LARGEUR_MATRICES)
		for j := 0; j < LARGEUR_MATRICES; j++ {
			matA[i][j] = rand.Intn(10000) // Default: 10000
		}
	}

	matB = make([][]int, HAUTEUR_MATRICES)
	for i := 0; i < HAUTEUR_MATRICES; i++ {
		matB[i] = make([]int, LARGEUR_MATRICES)
		for j := 0; j < LARGEUR_MATRICES; j++ {
			matB[i][j] = rand.Intn(10000) // Default: 10000
		}
	}

	fmt.Printf("\n#DEBUG END remplirMatrices\n")
}

func printMat(mat [][]int) {
	res := ""
	res += "\n\nMatrice\n"
	for i := 0; i < HAUTEUR_MATRICES; i++ {
		for j := 0; j < LARGEUR_MATRICES; j++ {
			res += fmt.Sprintf("%d ", mat[i][j])
		}
		res += "\n"
	}
	fmt.Printf("%s", res)
}

func printMatLine(mat [][]int, from int, to int) {
<<<<<<< HEAD
	for line := from; line < to; line++ {
=======
	for line := from; line <= to; line++ {
>>>>>>> 8040cb6ab0fd2c8d800076a00795f5fd9ee30fd4
		res := ""
		if line == 0 {
			res += "\n\nMatrice\n"
		}
		res += fmt.Sprintf("\nLine %d\n", line)
		for j := 0; j < LARGEUR_MATRICES; j++ {
			res += fmt.Sprintf("%d ", mat[line][j])
		}
		res += "\n"
		fmt.Printf("%s", res)

	}
	wg.Done()
	
}

func main() {
	if inc == 0 || inc > 500 {
		fmt.Print(":^)")
		return
	}
	fmt.Printf("#DEBUG START\n")
	timeStart := time.Now()

	result = make([][]int, HAUTEUR_MATRICES)
	remplirMatrices()

	fmt.Printf("#DEBUG START GOROUTINES\n")
	for i := 0; i < HAUTEUR_MATRICES; i += inc {
		wg.Add(1)
		go multiplicationByLine(i, i+inc-1, matA, matB)
	}

	wg.Wait()
	fmt.Printf("#DEBUG END GOROUTINES\n")
	fmt.Printf("#DEBUG START GOROUTINES PRINTLINES\n")

	for i := 0; i < HAUTEUR_MATRICES; i += inc {
		wg.Add(1)
		go printMatLine(result, i, i+inc)
		time.Sleep(time.Millisecond*inc)	//Cette ligne nous permet d'avoir la matrice affichée dans l'ordre
	}

	wg.Wait()
	fmt.Printf("#DEBUG ALL PRINTLINES GOROUTINES ENDED\n")
	timeEnd := time.Now()
	elapsed := timeEnd.Sub(timeStart)
	fmt.Printf("Time elapsed : %d", elapsed.Milliseconds())
}

func multiplicationByLine(from int, to int, matA [][]int, matB [][]int) {
	for line := from; line <= to; line++ {
		result[line] = make([]int, LARGEUR_MATRICES)
		for j := 0; j < LARGEUR_MATRICES; j++ {
			for l := 0; l < HAUTEUR_MATRICES; l++ {
				result[line][j] = result[line][j] + matA[line][l]*matB[l][j]
			}
		}
	}
	wg.Done()
}

func possibleProduct(rA int, cA int, rB int, cB int) (check bool) {
	if cA != rB {
		fmt.Print("Multiplication de matrices impossible")
		return false
	} else {
		fmt.Print("Multiplication de matrices possible")
		return true
	}
}
