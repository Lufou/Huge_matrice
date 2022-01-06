package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func main() {
	if runtime.GOOS != "windows" {
		fmt.Printf("Sorry but this script is only made for Windows!")
		os.Exit(1)
	}
	if len(os.Args) < 4 {
		fmt.Printf("start.exe <number of clients> <max_size> <max_values>")
		os.Exit(1)
	}
	string_amount := os.Args[1]
	string_sizes := os.Args[2]
	string_max_values := os.Args[3]
	client_amount, err_conv := strconv.Atoi(string_amount)
	sizes, err_conv2 := strconv.Atoi(string_sizes)
	max_values, err_conv3 := strconv.Atoi(string_max_values)
	if err_conv != nil || err_conv2 != nil || err_conv3 != nil {
		fmt.Printf("Incorrect arg passed, start.exe <number of clients> <max_sizes> <max_values>")
		os.Exit(1)
	}

	s := []string{"cmd", "/C", "start", "./server.exe", "6000"}
	cmd := exec.Command(s[0], s[1:]...)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
	time.Sleep(2 * time.Second)
	for i := 0; i < client_amount; i++ {
		rand.Seed(time.Now().UnixNano())
		ra := fmt.Sprintf("%d", rand.Intn(sizes-1)+2)
		ca := fmt.Sprintf("%d", rand.Intn(sizes-1)+2)
		cb := fmt.Sprintf("%d", rand.Intn(sizes-1)+2)
		max_value := fmt.Sprintf("%d", rand.Intn(max_values-9)+10)

		s = []string{"cmd.exe", "/k", "cmd", "/C", "start", "./client.exe", "6000", ra, ca, ca, cb, max_value}

		cmd := exec.Command(s[0], s[1:]...)
		if err := cmd.Run(); err != nil {
			fmt.Println("Error:", err)
		}
		time.Sleep(1 * time.Millisecond)
	}
}
