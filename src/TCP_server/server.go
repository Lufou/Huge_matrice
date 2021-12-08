package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

type matrix_line struct {
	id          int
	line_string string
}

const inc = 5
var numGoroutine int
var wg sync.WaitGroup

func getArgs() int {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: go run server.go <portnumber>\n")
		os.Exit(1)
	} else {
		fmt.Printf("#DEBUG ARGS Port Number : %s\n", os.Args[1])
		portNumber, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Usage: go run server.go <portnumber>\n")
			os.Exit(1)
		} else {
			return portNumber
		}

	}
	return -1
}

func main() {
	port := getArgs()
	fmt.Printf("#DEBUG MAIN Creating TCP Server on port %d\n", port)
	portString := fmt.Sprintf(":%s", strconv.Itoa(port))
	fmt.Printf("#DEBUG MAIN PORT STRING |%s|\n", portString)

	ln, err := net.Listen("tcp", portString)
	if err != nil {
		fmt.Printf("#DEBUG MAIN Could not create listener\n")
		panic(err)
	}

	//If we're here, we did not panic and ln is a valid listener

	connum := 1

	for {
		fmt.Printf("#DEBUG MAIN Accepting next connection\n")
		conn, errconn := ln.Accept()

		if errconn != nil {
			fmt.Printf("DEBUG MAIN Error when accepting next connection\n")
			panic(errconn)

		}

		//If we're here, we did not panic and conn is a valid handler to the new connection

		go handleConnection(conn, connum)
		connum += 1

	}

}

func possibleProduct(rA int, cA int, rB int, cB int) bool {
	if cA != rB {
		return false
	} else {
		return true
	}
}

func handleConnection(connection net.Conn, connum int) {

	defer connection.Close()
	connReader := bufio.NewReader(connection)

	for {
		inputLine, err := connReader.ReadString('\n')
		if err != nil {
			fmt.Printf("#DEBUG %d RCV ERROR no panic, just a client\n", connum)
			fmt.Printf("Error :|%s|\n", err.Error())
			break
		}

		inputLine = strings.TrimSuffix(inputLine, "\n")
		fmt.Printf("#DEBUG %d RCV |%s|\n", connum, inputLine)
		//Check each int and see if it's real ints
		splitLine := strings.Split(inputLine, " ")
		str_hauteur_mat1 := splitLine[0]
		str_largeur_mat1 := splitLine[1]
		str_hauteur_mat2 := splitLine[2]
		str_largeur_mat2 := splitLine[3]
		str_int_max_value := splitLine[4]

		hauteur_mat1, err1 := strconv.Atoi(str_hauteur_mat1)
		largeur_mat1, err2 := strconv.Atoi(str_largeur_mat1)
		hauteur_mat2, err3 := strconv.Atoi(str_hauteur_mat2)
		largeur_mat2, err4 := strconv.Atoi(str_largeur_mat2)
		int_max_value, err5 := strconv.Atoi(str_int_max_value)

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || hauteur_mat1 <= 0 || largeur_mat2 <= 0 || hauteur_mat2 <= 0 || largeur_mat1 <= 0 || int_max_value <= 0 {
			io.WriteString(connection, fmt.Sprintf("%send\n", "Wrong argument provided."))
			fmt.Printf("#DEBUG %d RCV ERROR : wrong arguments, no panic, just a client\n", connum)
			break
		}

		//Check if mat can multiplied (with possibleProduct function)
		if !possibleProduct(hauteur_mat1, largeur_mat1, hauteur_mat2, largeur_mat2) {
			io.WriteString(connection, fmt.Sprintf("%send\n", "Matrix cannot be multiplied."))
			fmt.Printf("#DEBUG %d RCV ERROR : matrix cannot be multiplied, no panic, just a client\n", connum)
			break
		}

		io.WriteString(connection, fmt.Sprintf("%s\n", "Matrix can be multiplied"))
		//Matrix generation
		matA, matB := remplirMatrices(hauteur_mat1, largeur_mat1, hauteur_mat2, largeur_mat2, int_max_value)
		//Prints the 2 mat to client?

		//Do the calculation of mat multiplication
		
		result := make([][]int, hauteur_mat1)
		//var result [][]int
		fmt.Printf("#DEBUG %d START GOROUTINES\n", connum) // debug
		numGoroutine = 0
		for i := 0; i < hauteur_mat1; i += int_max_value {
			fmt.Println(i)
			numGoroutine+=1
			wg.Add(1)                                                   // ajout d'un token
			go multiplicationByLine(i, i+int_max_value-1, matA, matB, result, numGoroutine, connum, connection) // lancement des goroutines qui effectuent le calcul
		}

		wg.Wait()                                                     // on attend ici que le nombre de tokens soit nul
		fmt.Printf("#DEBUG %d END GOROUTINES\n", connum)              // debug
		fmt.Printf("#DEBUG %d START GOROUTINES PRINTLINES\n", connum) // debug

		//Say DONE to the client with the elapsed time
		//Then send to the client each lines with id (in struct)

		//And the client will reassemble them (pretty quick I think) to print the whole mat

		returnedString := "OK"
		fmt.Printf("#DEBUG %d RCV Returned value |%s|\n", connum, returnedString)
		io.WriteString(connection, fmt.Sprintf("%s\n", returnedString))
	}
}

func multiplicationByLine(from int, to int, matA [][]int, matB [][]int, result [][]int, numGoroutine int, connum int, connection net.Conn) {
	//Creation structure
	var matrix matrix_line 
	matrix.id=numGoroutine
	matrix.line_string=""

	for line := from; line <= to; line++ { // parcours des lignes de la matrice résultat
		matrix.line_string += "\n" + " Line " + strconv.Itoa(line) + " : " + "\n"
		result[line] = make([]int, len(matB[line])) // déclaration du tableau stockant une ligne de résultats
		for j := 0; j < len(matB[line]); j++ {      // parcours des colonnes de la matrice
			for l := 0; l < len(matB); l++ { // parcours des lignes de la matrice
				result[line][j] = result[line][j] + matA[line][l]*matB[l][j] // calcul du coefficient à la j-eme colonne de la ligne en cours de calcul
			}
			matrix.line_string +=strconv.Itoa(result[line][j]) + " "
		}
	}
	wg.Done()	
	
	//envoi vers une méthode qui permet d'envoyer la struct
	envoiStruct(matrix, connum, connection)
}

func envoiStruct(matrix matrix_line, connum int, connection net.Conn){
	fmt.Printf("#DEBUG %d RCV Returned value |%s|\n", connum, matrix.id, "\n", matrix.line_string, "\n")
	io.WriteString(connection, fmt.Sprintf("%s\n", matrix.id, matrix.line_string))
}

func remplirMatrices(hauteur_matA int, largeur_matA int, hauteur_matB int, largeur_matB int, max_value int) ([][]int, [][]int) {
	fmt.Printf("#DEBUG START remplirMatrices\n") // debug
	matA := make([][]int, hauteur_matA)          // déclaration du tableau contenant le nombre de lignes de la matrice A
	for i := 0; i < hauteur_matA; i++ {          // parcours de la matrice A par lignes
		matA[i] = make([]int, largeur_matA) // pour chaque ligne i, on génère un tableau contenant [largeur_matA] colonnes
		for j := 0; j < largeur_matA; j++ { // parcours de la matrice par colonnes
			matA[i][j] = rand.Intn(max_value) // Default: 10000
			// remplissage de l'élément [i][j] de matA par une valeur aléatoire comprise entre 1 et max_value
		}
	}

	matB := make([][]int, hauteur_matB) // déclaration du tableau contenant le nombre de lignes de la matrice B
	for i := 0; i < hauteur_matB; i++ { // parcours de la matrice B par lignes
		matB[i] = make([]int, largeur_matB) // pour chaque ligne i, on génère un tableau contenant [largeur_matB] colonnes
		for j := 0; j < largeur_matB; j++ { // parcours de la matrice par colonnes
			matB[i][j] = rand.Intn(max_value) // Default: 10000
			// remplissage de l'élément [i][j] de matB par une valeur aléatoire comprise entre 1 et max_value
		}
	}

	fmt.Printf("\n#DEBUG END remplirMatrices\n") // debug

	return matA, matB
}
