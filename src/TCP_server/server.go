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
	"time"
)

type matrix_line struct {
	id          int
	line_string string
}

var inc = 1

var wg_slice []sync.WaitGroup
var lock sync.Mutex

func getArgs() int { // test the number of arguments and return port number, we need 2 arguments : file name, port number
	if len(os.Args) != 2 {
		fmt.Printf("Usage: go run server.go <portnumber>\n")
		os.Exit(1) // end
	} else {
		fmt.Printf("#DEBUG ARGS Port Number : %s\n", os.Args[1])
		portNumber, err := strconv.Atoi(os.Args[1]) // retrieve the port number and convert it to int
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

	connum := 0

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

func possibleProduct(rA int, cA int, rB int, cB int) bool { // test if the matrices product is possible
	return cA == rB
}

func handleConnection(connection net.Conn, connum int) { // handle the connection server-client

	defer connection.Close()                  // at the end of the handleConnection fonction, close the connection
	connReader := bufio.NewReader(connection) // we wait for client's request
	var wg sync.WaitGroup                     // initialization of the token group
	lock.Lock()
	wg_slice = append(wg_slice, wg) // add wg to the wg_slice list
	lock.Unlock()
	for {
		inputLine, err := connReader.ReadString('\n') // read the string sent by the client
		if err != nil {
			fmt.Printf("#DEBUG %d RCV ERROR no panic, just a client\n", connum)
			fmt.Printf("Error :|%s|\n", err.Error())
			break // leave the for and close the connection
		}

		start_time := time.Now() // save the date

		inputLine = strings.TrimSuffix(inputLine, "\n")
		fmt.Printf("#DEBUG %d RCV |%s|\n", connum, inputLine)
		//Check each int and see if it's real ints
		splitLine := strings.Split(inputLine, " ") // make a table of string based on a string, the separator is " "
		str_hauteur_mat1 := splitLine[0]
		str_largeur_mat1 := splitLine[1]
		str_hauteur_mat2 := splitLine[2]
		str_largeur_mat2 := splitLine[3]
		str_int_max_value := splitLine[4]

		hauteur_mat1, err1 := strconv.Atoi(str_hauteur_mat1) // conversion from string to int
		largeur_mat1, err2 := strconv.Atoi(str_largeur_mat1)
		hauteur_mat2, err3 := strconv.Atoi(str_hauteur_mat2)
		largeur_mat2, err4 := strconv.Atoi(str_largeur_mat2)
		int_max_value, err5 := strconv.Atoi(str_int_max_value)

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil || hauteur_mat1 <= 0 || largeur_mat2 <= 0 || hauteur_mat2 <= 0 || largeur_mat1 <= 0 || int_max_value <= 0 { // checking
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

		io.WriteString(connection, fmt.Sprintf("%s\n", "Matrices can be multiplied"))

		//Matrices generation
		matA, matB := remplirMatrices(hauteur_mat1, largeur_mat1, hauteur_mat2, largeur_mat2, int_max_value)

		//Prints the 2 mat to client?
		fmt.Printf("#DEBUG %d START MAT1 PRINTING\n", connum)
		inc = hauteur_mat1 / 10
		if inc == 0 {
			inc = 1
		}
		for i := 0; i < hauteur_mat1; i += inc {
			lock.Lock()
			wg_slice[connum].Add(1)
			lock.Unlock()
			if hauteur_mat1-i < inc {
				inc = 1
			}
			go printMat(i, i+inc-1, matA, connum, connection) // start the goroutine who prints matA
		}
		fmt.Printf("#DEBUG %d WAITING FOR PRINTING TO END\n", connum)
		lock.Lock()
		wg_slice[connum].Wait() // wait for the token group to be empty
		lock.Unlock()
		fmt.Printf("#DEBUG %d START MAT2 PRINTING\n", connum)
		inc = hauteur_mat2 / 10
		if inc == 0 {
			inc = 1
		}
		for i := 0; i < hauteur_mat2; i += inc {
			lock.Lock()
			wg_slice[connum].Add(1)
			lock.Unlock()
			if hauteur_mat2-i < inc { // if the size of the matrix is not a multiple of inc, last lines are processed one by one
				inc = 1
			}
			go printMat(i, i+inc-1, matB, connum, connection) // start the goroutine who prints matB
		}
		fmt.Printf("#DEBUG %d WAITING FOR PRINTING TO END\n", connum)
		lock.Lock()
		wg_slice[connum].Wait()
		lock.Unlock()
		//Do the calculation of mat multiplication
		result := make([][]int, hauteur_mat1) // initialization of the result matrix

		fmt.Printf("#DEBUG %d START MAT MULTIPLICATION\n", connum)
		inc = hauteur_mat1 / 10
		if inc == 0 {
			inc = 1
		}
		for i := 0; i < hauteur_mat1; i += inc {
			lock.Lock()
			wg_slice[connum].Add(1)
			lock.Unlock()
			if hauteur_mat1-i < inc {
				inc = 1
			}
			go multiplicationByLine(i, i+inc-1, matA, matB, result, connum, connection) // Launching calculation goroutines
		}

		lock.Lock()
		wg_slice[connum].Wait()
		lock.Unlock()
		fmt.Printf("#DEBUG %d END MULTIPLICAION\n", connum)

		//Say DONE to the client with the elapsed time
		elapsed_time := time.Since(start_time)
		returnedString := fmt.Sprintf("Done in %s", elapsed_time)
		fmt.Printf("#DEBUG %d RCV Returned value |%s|\n", connum, returnedString)
		//And the client will reassemble them to print the whole mat
	}
}

func printMat(from int, to int, mat [][]int, connum int, connection net.Conn) {
	for line_number := from; line_number <= to; line_number++ {
		//Creation struct
		var matrix_line matrix_line
		matrix_line.id = line_number
		matrix_line.line_string = ""
		for j := 0; j < len(mat[line_number]); j++ {
			matrix_line.line_string += strconv.Itoa(mat[line_number][j]) + " " // convert int to string
		}
		matrix_line.line_string = strings.TrimSuffix(matrix_line.line_string, " ") // trim the suffix

		envoiStruct(matrix_line, connum, connection) // send struct to client
	}
	wg_slice[connum].Done()
}

func multiplicationByLine(from int, to int, matA [][]int, matB [][]int, result [][]int, connum int, connection net.Conn) {
	for line_number := from; line_number <= to; line_number++ {
		result[line_number] = make([]int, len(matB[0]))
		for j := 0; j < len(matB[0]); j++ {
			for l := 0; l < len(matB); l++ {
				result[line_number][j] = result[line_number][j] + matA[line_number][l]*matB[l][j]
			}
		}
	}
	go printMat(from, to, result, connum, connection)
}

func envoiStruct(matrix_line matrix_line, connum int, connection net.Conn) {
	io.WriteString(connection, fmt.Sprintf("%d %s\n", matrix_line.id, matrix_line.line_string))
}

func remplirMatrices(hauteur_matA int, largeur_matA int, hauteur_matB int, largeur_matB int, max_value int) ([][]int, [][]int) {
	fmt.Printf("#DEBUG START remplirMatrices\n")
	matA := make([][]int, hauteur_matA)
	for i := 0; i < hauteur_matA; i++ {
		matA[i] = make([]int, largeur_matA)
		for j := 0; j < largeur_matA; j++ {
			matA[i][j] = rand.Intn(max_value)
		}
	}

	matB := make([][]int, hauteur_matB)
	for i := 0; i < hauteur_matB; i++ {
		matB[i] = make([]int, largeur_matB)
		for j := 0; j < largeur_matB; j++ {
			matB[i][j] = rand.Intn(max_value)
		}
	}

	fmt.Printf("\n#DEBUG END remplirMatrices\n")

	return matA, matB
}
