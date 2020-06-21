package main

import (
	"flag"
	"fmt"
	"net"
	"os/exec"
	"time"
	
	"github.com/james-daniels/sshtunneler/env"
)

var host string
var envir string
var db string
var coll string

func init() {
	flag.StringVar(&host, "h", "", "Enter database host to connect to")
	flag.StringVar(&envir, "e", "", "Enter environment to connect to")
	flag.StringVar(&db, "db", "", "Enter database to connect to")
	flag.StringVar(&coll, "c", "", "Enter collection to connect to")
}

func runSSH() {

	cbis := env.DB(envir, host, db, coll)

	conn, _ := net.Dial("tcp", ":" + cbis.LocalPort)
	if conn != nil {
		fmt.Println("Port: " + cbis.LocalPort + " is aleady open.")

		openURL(cbis.LocalPort, cbis.Path)
		return
	}

	sshArgs := []string{"-N", "-L", cbis.LocalPort + ":" + cbis.IP + ":" + cbis.RemotePort, cbis.JumpServer}

	err := exec.Command("ssh", sshArgs...).Start()
	if err != nil {
		fmt.Println(err)
	}

	time.Sleep(time.Second * cbis.Pause)

	openURL(cbis.LocalPort, cbis.Path)
}

func openURL(port string, path string) {

	err := exec.Command("open", "https://localhost:" + port + path).Start()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	flag.Parse()
	runSSH()
}
