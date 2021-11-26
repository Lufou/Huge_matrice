package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
)

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
		fmt.Print("Multiplication de matrices impossible")
		return false
	} else {
		fmt.Print("Multiplication de matrices possible")
		return true
	}
}

func handleConnection(connection net.Conn, connum int) {

	defer connection.Close()
	connReader := bufio.NewReader(connection)

	for {
		io.WriteString(connection, fmt.Sprintf("%s\n", "Hello, please provide matrix sizes"))
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

		hauteur_mat1, err1 := strconv.ParseUint(str_hauteur_mat1, 10, 32)
		largeur_mat1, err2 := strconv.ParseUint(str_largeur_mat1, 10, 32)
		hauteur_mat2, err3 := strconv.ParseUint(str_hauteur_mat2, 10, 32)
		largeur_mat2, err4 := strconv.ParseUint(str_largeur_mat2, 10, 32)

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil || hauteur_mat1 <= 0 || largeur_mat2 <= 0 || hauteur_mat2 <= 0 || largeur_mat1 <= 0 {
			io.WriteString(connection, fmt.Sprintf("%send\n", "Wrong matrix sizes provided."))
			fmt.Printf("#DEBUG %d RCV ERROR : wrong mat sizes, no panic, just a client\n", connum)
			break
		}
		//Then return OK to client
		io.WriteString(connection, fmt.Sprintf("%s\n", "Good matrix size received."))
		//Prints the 2 mat to client?

		//Do the calculation of mat multiplication
		//Say DONE to the client with the elapsed time
		//Then send to the client each lines with id (in tuples)?
		//And the client will reassemble them (pretty quick I think) to print the whole mat

		returnedString := "OK"
		fmt.Printf("#DEBUG %d RCV Returned value |%s|\n", connum, returnedString)
		io.WriteString(connection, fmt.Sprintf("%s\n", returnedString))
	}

}
