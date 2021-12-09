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

func getArgs() (int, int, int, int, int, int) {
	if len(os.Args) != 7 {
		fmt.Printf("Usage: go run client.go <portnumber> <mat1_height> <mat1_width> <mat2_height> <mat2_width> <max_int_value>\n")
		os.Exit(1)
	} else {
		portNumber, err := strconv.Atoi(os.Args[1])
		mat1_height, err1 := strconv.Atoi(os.Args[2])
		mat1_width, err2 := strconv.Atoi(os.Args[3])
		mat2_height, err3 := strconv.Atoi(os.Args[4])
		mat2_width, err4 := strconv.Atoi(os.Args[5])
		max_int_value, err5 := strconv.Atoi(os.Args[6])
		if err != nil || err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
			fmt.Printf("Usage: go run client.go <portnumber> <mat1_height> <mat1_width> <mat2_height> <mat2_width> <max_int_value>\n")
			os.Exit(1)
		} else {
			return portNumber, mat1_height, mat1_width, mat2_height, mat2_width, max_int_value
		}
	}
	return -1, -1, -1, -1, -1, -1
}

func main() {
	port, hauteur_mat1, largeur_mat1, hauteur_mat2, largeur_mat2, max_int_value := getArgs()
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

		io.WriteString(conn, fmt.Sprintf("%d %d %d %d %d\n", hauteur_mat1, largeur_mat1, hauteur_mat2, largeur_mat2, max_int_value))
		resultString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("#DEBUG MAIN could not read from server")
			os.Exit(1)
		}
		resultString = strings.TrimSuffix(resultString, "\n")
		fmt.Printf("#DEBUG server replied : |%s|\n", strings.Replace(resultString, "end", "", 1))
		if strings.Contains(resultString, "end") {
			fmt.Printf("#DEBUG server decided to end the connection.")
			return
		}

		resultString, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("DEBUG MAIN could not read from server")
			os.Exit(1)
		}
		resultString = strings.TrimSuffix(resultString, "\n")
		fmt.Printf("#DEBUG server replied : |%s|\n", strings.Replace(resultString, "end", "", 1))
		if strings.Contains(resultString, "end") {
			fmt.Printf("#DEBUG server decided to end the connection.")
			return
		}
		
		nbRecept, err := strconv.Atoi(resultString)
		resultString2, err := reader.ReadString('\n')
		
		for i :=0; i < nbRecept+1; i++{
			resultString2, err = reader.ReadString('\n')
			if err != nil {
				fmt.Printf("DEBUG MAIN could not read from server")
				os.Exit(1)
			}
			resultString2 = strings.TrimSuffix(resultString2, "\n")
			fmt.Printf("#DEBUG server replied : |%s|\n", strings.Replace(resultString2, "end", "", 1))
			if strings.Contains(resultString2, "end") {
				fmt.Printf("#DEBUG server decided to end the connection.")
				return
			}
		}
	}
}
