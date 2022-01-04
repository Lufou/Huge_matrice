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
	if len(os.Args) < 2 {
		fmt.Printf("start.exe <number of clients>")
		os.Exit(1)
	}
	string_amount := os.Args[1]
	client_amount, err_conv := strconv.Atoi(string_amount)
	if err_conv != nil {
		fmt.Printf("Incorrect arg passed, start.exe <number of clients>")
		os.Exit(1)
	}

	s := []string{"cmd", "/C", "start", "./server.exe", "6000"}
	cmd := exec.Command(s[0], s[1:]...)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
	time.Sleep(2000)
	for i := 0; i < client_amount; i++ {
		rand.Seed(time.Now().UnixNano())
		ra := fmt.Sprintf("%d", rand.Intn(99)+2)
		ca := fmt.Sprintf("%d", rand.Intn(99)+2)
		cb := fmt.Sprintf("%d", rand.Intn(99)+2)
		max_value := fmt.Sprintf("%d", rand.Intn(91)+10)

		s = []string{"cmd.exe", "/k", "cmd", "/C", "start", "./client.exe", "6000", ra, ca, ca, cb, max_value}

		cmd := exec.Command(s[0], s[1:]...)
		if err := cmd.Run(); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
