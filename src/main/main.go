package main

import (
	"fmt"
	"math/rand"
)

func remplirMatrices(matList [][][]int, nombre int, sizes [][]int) {
	for k := 0; k < nombre; k++ {
		mat := make([][]int, sizes[k][0])
		for i := 0; i < len(mat); i++ {
			mat[i] = make([]int, sizes[k][1])
			for j := 0; j < len(mat[0]); j++ {
				mat[i][j] = rand.Intn(10000)
			}
		}
		matList[k] = mat
	}
}

func printMat(mat [][]int) {
	res := ""
	res += "\n\nMatrice\n"
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			res += fmt.Sprintf("%v ", mat[i][j])
		}
		res += "\n"
	}

	fmt.Printf("%s", res)
}

func printMatList(matList [][][]int, nombre int) {
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
	sizes := make([][]int, len(matList))
	for i := 0; i < len(matList); i++ {
		sizes[i] = make([]int, 2)
	}
	sizes[0][0] = 30
	sizes[0][1] = 30
	sizes[1][0] = 30
	sizes[1][1] = 30

	remplirMatrices(matList, 2, sizes)
	printMatList(matList, 2)

	printMat(multiplication(matList))
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

func multiplication(matList [][][]int) [][]int {
	var matMult [][]int

	for k := 1; k < len(matList); k++ {
		matMult = make([][]int, len(matList[k]))
		for i := 0; i < len(matList[k]); i++ {
			matMult[i] = make([]int, len(matList[k][i]))
			for j := 0; j < len(matList[k][0]); j++ {
				matMult[i][j] = matList[k-1][i][j] * matList[k][i][j]
			}
		}
	}

	return matMult
}
