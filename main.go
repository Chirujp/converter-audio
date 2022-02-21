package main

import (
	// "os"
	// "os/exec"
	// "bufio"
	// "io"
	// "bytes"
	"crypto/tls"
	"fmt"
	"log"
	"os"

	"github.com/valyala/fasthttp"
)

var Connexion = 0

func getCa() ([]byte, []byte) {
	caFile, err := os.ReadFile("./security/cert.pem")

	if err != nil {
		panic(err)
	}

	keyFile, err := os.ReadFile("./security/priv.pem")

	if err != nil {
		panic(err)
	}

	return caFile, keyFile
}

func main() {
	addr := "127.0.0.1:3000"

	caFile, keyFile := getCa()

	log.Println("Listening on " + addr)

	cert, err := tls.X509KeyPair(caFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cert)

	fasthttp.ListenAndServeTLSEmbed(addr, caFile, keyFile, Handler)
}
