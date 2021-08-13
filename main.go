package main

import (
	"os"
	server "github.com/0187773933/FireStickC2Server/v1/server"
)

// https://github.com/48723247842/MediaBox/blob/master/C2Server/main.go
func main() {
	server := server.NewServer( os.Args[ 1 ] )
	server.Listen()
}