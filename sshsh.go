package main

import (
	"golang.org/x/crypto/ssh"
	"flag"
	"net"
	"log"
	"os"
)

func main() {
	username := flag.String("u","", "username")
	password := flag.String("p", "", "password")
	host	:= flag.String("h", "", "host")
	port	:= flag.String("P", "22", "port")

	flag.Parse()

	config := &ssh.ClientConfig {
			User: *username,
			Auth: []ssh.AuthMethod {
				ssh.Password(*password),
			},

			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}

	addr := net.JoinHostPort(*host, *port)

	client, err := ssh.Dial("tcp", addr, config)
	if err != nil {
		log.Fatalln(err)
	}

	session, err := client.NewSession()
	if err != nil {
		log.Fatalln(err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
	session.Stdin = os.Stdin

	modes := ssh.TerminalModes {
		ssh.ECHO:	0,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}

	if err := session.RequestPty("xterm", 40, 80, modes); err != nil {
		log.Fatal("request for pseudo terminal failed: ", err)
	}

	if err := session.Shell(); err != nil {
		log.Fatal("failed to start shell: ", err)
	}

	if err := session.Wait(); err != nil {
		log.Fatal("failed to wait shell: ", err)
	}
}
