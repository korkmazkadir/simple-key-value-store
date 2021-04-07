package main

import (
	"log"
	"net"
	"net/rpc"
	"os"

	simplekeyvalue "github.com/korkmazkadir/simple-key-value-store"
)

const sockAddr = "/tmp/keyvalue-store.sock"

func main() {

	log.Printf("Keyvalue server started...\n")

	server := simplekeyvalue.NewServer()
	rpc.Register(server)

	if err := os.RemoveAll(sockAddr); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("unix", sockAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	log.Printf("Listening on unix domain socked %s\n", sockAddr)

	rpc.Accept(l)

	/*

		for {

			conn, err := l.Accept()
			log.Println("connection accepted")
			if err == nil {
				rpc.ServeConn(conn)
			} else {
				log.Println(err)
			}
		}
	*/

}
