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
	command := flag.String("c", "", "command")

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

	if err := session.Run(*command); err != nil {
		log.Fatalln(err)
	}
}
