package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("start.go <number of clients>")
		os.Exit(1)
	}
	string_amount := os.Args[1]
	client_amount, err_conv := strconv.Atoi(string_amount)
	if err_conv != nil {
		fmt.Printf("Incorrect arg passed, start.go <number of clients>")
		os.Exit(1)
	}
	s := []string{"go", "run", "../TCP_server/server.go"} // LINUX PART TO DO
	if runtime.GOOS == "windows" {
		s = []string{"cmd.exe", "/k", "cmd", "/C", "start", "./server.exe", "6000"}
	}
	cmd := exec.Command(s[0], s[1:]...)
	if err := cmd.Run(); err != nil {
		fmt.Println("Error:", err)
	}
	for i := 0; i < client_amount; i++ {
		s := []string{"go", "run", "../TCP_client/client.go"} // LINUX PART TO DO
		if runtime.GOOS == "windows" {
			s = []string{"cmd.exe", "/k", "cmd", "/C", "start", "./client.exe", "6000", "50", "50", "50", "50", "10"}
		}

		cmd := exec.Command(s[0], s[1:]...)
		if err := cmd.Run(); err != nil {
			fmt.Println("Error:", err)
		}
	}
}
