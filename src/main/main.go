package main

import (
	"fmt"
	"math/rand"
)

var TAILLE_MATRICES int = 30

func remplirMatrices(matList [][][]int, nombre int) {
	for k := 0; k < nombre; k++ {
		mat := make([][]int, TAILLE_MATRICES)
		for i := 0; i < len(mat); i++ {
			mat[i] = make([]int, TAILLE_MATRICES)
			for j := 0; j < len(mat[0]); j++ {
				mat[i][j] = rand.Intn(10000)
			}
		}
		matList[k] = mat
	}
}

func printMat(matList [][][]int, nombre int) {
	res := ""
	if nombre > len(matList) {
		fmt.Printf("Pas assez de matrices dans la liste de matrices pour en afficher %d", nombre)
		return
	}
	for k := 0; k < nombre; k++ {
		res += fmt.Sprintf("\n\nMatrice %d\n", k+1)
		currentMat := matList[k]
		for i := 0; i < len(currentMat); i++ {
			for j := 0; j < len(currentMat[0]); j++ {
				res += fmt.Sprintf("%v ", currentMat[i][j])
			}
			res += "\n"
		}
	}

	fmt.Printf("%s", res)
}

func main() {
	matList := make([][][]int, 2)

	remplirMatrices(matList, 2)

	printMat(matList, 2)

	//fonction pour multiplier les matrices, basique
	// après, la même avec des coroutines
}
