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
		fmt.Printf("Usage: go run client.go <portnumber>\n")
		os.Exit(1)
	} else {
		fmt.Printf("#DEBUG ARGS Port Number : %s\n", os.Args[1])
		portNumber, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Printf("Usage: go run client.go <portnumber>\n")
			os.Exit(1)
		} else {
			return portNumber
		}

	}
	return -1
}

func main() {
	port := getArgs()
	fmt.Printf("#DEBUG DIALING TCP Server on port %d\n", port)
	portString := fmt.Sprintf("127.0.0.1:%s", strconv.Itoa(port))
	fmt.Printf("#DEBUG MAIN PORT STRING |%s|\n", portString)

	conn, err := net.Dial("tcp", portString)
	if err != nil {
		fmt.Printf("#DEBUG MAIN could not connect\n")
		os.Exit(1)
	} else {
		defer conn.Close()
		reader := bufio.NewReader(conn)
		fmt.Printf("#DEBUG MAIN connected\n")
		resultString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("#DEBUG MAIN could not read from server")
			os.Exit(1)
		}
		resultString = strings.TrimSuffix(resultString, "\n")
		fmt.Printf("#DEBUG server input : |%s|\n", resultString)
		var hauteur_mat1 string
		var largeur_mat1 string
		var hauteur_mat2 string
		var largeur_mat2 string
		fmt.Scanf("%s", &hauteur_mat1)
		fmt.Scanf("%s", &largeur_mat1)
		fmt.Scanf("%s", &hauteur_mat2)
		fmt.Scanf("%s", &largeur_mat2)

		io.WriteString(conn, fmt.Sprintf("%s %s %s %s\n", hauteur_mat1, largeur_mat1, hauteur_mat2, largeur_mat2))
		resultString, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("#DEBUG MAIN could not read from server")
			os.Exit(1)
		}
		resultString = strings.TrimSuffix(resultString, "\n")
		fmt.Printf("#DEBUG server replied : |%s|\n", resultString)
		if strings.Contains(resultString, "end") {
			fmt.Printf("#DEBUG server decided to end the connection.")
			return
		}

		var max_int_value string
		fmt.Scanf("%s\n", &max_int_value)
		io.WriteString(conn, fmt.Sprintf("%s\n", max_int_value))
		resultString, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("DEBUG MAIN could not read from server")
			os.Exit(1)
		}
		resultString = strings.TrimSuffix(resultString, "\n")
		fmt.Printf("#DEBUG server replied : |%s|\n", resultString)
		if strings.Contains(resultString, "end") {
			fmt.Printf("#DEBUG server decided to end the connection.")
			return
		}
	}
}
