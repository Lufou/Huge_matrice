package main

import "fmt"

var matA[][] int 
var matB[][] int 

var taille int = 5

func multiplication(matA[][] int, matB[][] int){
	var matMult[][] int = make([][]int, taille)

	for i := 0; i < len(matA); i++ {
		matMult[i] = make([]int, len(matA))
		for j := 0; j < len(matA); j++{
			matMult[i][j] = matA[i][j] * matB[i][j]
		}
	}

	fmt.Println("Mult :", matMult)
}

func main() {
	matA := make([][]int, taille)
	for i := 0; i < taille; i++ {
        matA[i] = make([]int, taille)
        for j := 0; j < taille; j++ {
            matA[i][j] = i + j
        }
    }
	matB := make([][]int, taille)
	for i := 0; i < taille; i++ {
        matB[i] = make([]int, taille)
        for j := 0; j < taille; j++ {
            matB[i][j] = i + j
        }
    }

	fmt.Println("matA:", matA)
	fmt.Println("matB:", matB)

	multiplication(matA, matB)
}

