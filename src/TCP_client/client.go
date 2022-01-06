package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func getArgs() (int, int, int, int, int, int) {
	//Retrieve the arguments
	//<Port Number> <Height Matrix 1> <Width Matrix 1> <Height Matrix 2> <Width Matrix 2> <Maximum int value in the matrices 1 and 2>

	if len(os.Args) != 7 { //Test if we have the good number of arguments
		fmt.Printf("Usage: go run client.go <portnumber> <mat1_height> <mat1_width> <mat2_height> <mat2_width> <max_int_value>\n")
		os.Exit(1)
	}
	portNumber, err := strconv.Atoi(os.Args[1]) // Try convert strings to ints
	mat1_height, err1 := strconv.Atoi(os.Args[2])
	mat1_width, err2 := strconv.Atoi(os.Args[3])
	mat2_height, err3 := strconv.Atoi(os.Args[4])
	mat2_width, err4 := strconv.Atoi(os.Args[5])
	max_int_value, err5 := strconv.Atoi(os.Args[6])

	if err != nil || err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil { //Check conversion errors
		fmt.Printf("Usage: go run client.go <portnumber> <mat1_height> <mat1_width> <mat2_height> <mat2_width> <max_int_value>\n")
		os.Exit(1)
	}

	return portNumber, mat1_height, mat1_width, mat2_height, mat2_width, max_int_value
}

func main() {
	port, hauteur_mat1, largeur_mat1, hauteur_mat2, largeur_mat2, max_int_value := getArgs() //We retrieve the arguments in variables
	fmt.Printf("#DEBUG DIALING TCP Server on port %d\n", port)
	portString := fmt.Sprintf("127.0.0.1:%s", strconv.Itoa(port))
	fmt.Printf("#DEBUG MAIN PORT STRING |%s|\n", portString)

	conn, err := net.Dial("tcp", portString)
	if err != nil {
		fmt.Printf("#DEBUG MAIN could not connect\n")
		fmt.Scanln()
		os.Exit(1)
	}

	defer conn.Close()
	reader := bufio.NewReader(conn)
	fmt.Printf("#DEBUG MAIN connected\n")

	start_time := time.Now()
	fmt.Printf("Product of %dx%d mat with %dx%d mat, max values = %d\n", hauteur_mat1, largeur_mat1, hauteur_mat2, largeur_mat2, max_int_value)
	io.WriteString(conn, fmt.Sprintf("%d %d %d %d %d\n", hauteur_mat1, largeur_mat1, hauteur_mat2, largeur_mat2, max_int_value))

	resultString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("#DEBUG MAIN could not read from server")
		fmt.Scanln()
		os.Exit(1)
	}
	resultString = strings.TrimSuffix(resultString, "\n")
	fmt.Printf("#DEBUG server replied : |%s|\n", strings.Replace(resultString, "end", "", 1))
	if strings.Contains(resultString, "end") {
		fmt.Printf("#DEBUG server decided to end the connection.")
		fmt.Scanln()
		return
	}

	fmt.Printf("Awaiting mat1\n")
	// Receiving matrix 1
	mat1_string := make([]string, hauteur_mat1)

	for i := 0; i < hauteur_mat1; i++ {
		resultString, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("DEBUG MAIN could not read from server")
			fmt.Scanln()
			os.Exit(1)
		}
		resultString = strings.TrimSuffix(resultString, "\n")
		mat1_string[i] = resultString
	}

	fmt.Printf("Awaiting mat2\n")
	// Receiving matrix 2
	mat2_string := make([]string, hauteur_mat2)

	for i := 0; i < hauteur_mat2; i++ {
		resultString, err = reader.ReadString('\n')
		if err != nil {
			fmt.Printf("DEBUG MAIN could not read from server")
			fmt.Scanln()
			os.Exit(1)
		}
		resultString = strings.TrimSuffix(resultString, "\n")
		mat2_string[i] = resultString
	}

	fmt.Printf("Awaiting results\n")
	// Receiving result
	result_string := make([]string, hauteur_mat1)
	fmt.Printf("Wesh poto %d\n", hauteur_mat1)
	for i := 0; i < hauteur_mat1; i++ {
		fmt.Printf(" I'm here %d\n", i)
		resultString, err = reader.ReadString('\n')
		fmt.Printf(" I'm here now %d\n", i)
		if err != nil {
			fmt.Printf("DEBUG MAIN could not read from server")
			fmt.Scanln()
			os.Exit(1)
		}
		resultString = strings.TrimSuffix(resultString, "\n")
		result_string[i] = resultString
		fmt.Printf("And now i'm at the end %d\n", i)
	}

	// Printing Matrix 1
	fmt.Printf("\nMatrix 1\n")
	matA := make([][]int, hauteur_mat1)

	for i := 0; i < hauteur_mat1; i++ {
		split := strings.Split(mat1_string[i], " ")
		line_number, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Printf("ERROR Incorrect string format received from server. FATAL ERROR")
			fmt.Scanln()
			return
		}

		matA[line_number] = make([]int, largeur_mat1)

		for j := 0; j < len(split)-1; j++ {
			value, err := strconv.Atoi(split[j+1])
			if err != nil {
				fmt.Printf("ERROR Incorrect string format received from server. FATAL ERROR")
				fmt.Scanln()
				return
			}
			matA[line_number][j] = value
		}
	}

	printMat(matA)

	// Printing Matrix 2
	fmt.Printf("\nMatrix 2\n")
	matB := make([][]int, hauteur_mat2)

	for i := 0; i < hauteur_mat2; i++ {
		split := strings.Split(mat2_string[i], " ")
		line_number, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Printf("ERROR Incorrect string format received from server. FATAL ERROR")
			fmt.Scanln()
			return
		}

		matB[line_number] = make([]int, largeur_mat2)

		for j := 0; j < len(split)-1; j++ {
			value, err := strconv.Atoi(split[j+1])
			if err != nil {
				fmt.Printf("ERROR Incorrect string format received from server. FATAL ERROR")
				fmt.Scanln()
				return
			}
			matB[line_number][j] = value
		}
	}

	printMat(matB)

	// Printing Matrix Result
	fmt.Printf("\nResult\n")
	result := make([][]int, hauteur_mat1)

	for i := 0; i < hauteur_mat1; i++ {
		split := strings.Split(result_string[i], " ")
		line_number, err := strconv.Atoi(split[0])
		if err != nil {
			fmt.Printf("ERROR Incorrect string format received from server. FATAL ERROR")
			fmt.Scanln()
			return
		}

		result[line_number] = make([]int, largeur_mat2)

		for j := 0; j < len(split)-1; j++ {
			value, err := strconv.Atoi(split[j+1])
			if err != nil {
				fmt.Printf("ERROR Incorrect string format received from server. FATAL ERROR")
				fmt.Scanln()
				return
			}
			result[line_number][j] = value
		}
	}

	printMat(result)

	elapsed_time := time.Since(start_time)               //Calculating elapsed time
	fmt.Printf("\n\n\nFinished in %s\n\n", elapsed_time) //Printing elapsed time
	fmt.Println("Press the Enter Key to terminate the console screen!")
	fmt.Scanln() //Wait for Enter Key, useful to prevent the terminal from closing

}

func printMat(mat [][]int) {
	for i := 0; i < len(mat); i++ {
		fmt.Printf("|")
		for j := 0; j < len(mat[0]); j++ {
			fmt.Print(mat[i][j])
			if j != len(mat[0])-1 {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("|\n")
	}
}
