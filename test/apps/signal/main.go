package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/flynn/flynn/discoverd/client"
)

const service = "signal-service"

func main() {
	log.SetFlags(log.Lmicroseconds)
	ch := make(chan os.Signal, 1)
	log.Println("setting signal handler")
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println("registering service")
	if _, err := discoverd.AddServiceAndRegister(service, ":12345"); err != nil {
		log.Fatal(err)
	}
	log.Println("waiting for signal")
	sig := <-ch
	fmt.Printf("got signal: %s", sig)
}
