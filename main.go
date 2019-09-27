package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func register(c chan os.Signal) {
	counter := 0
	for s := range c {
		counter++
		if counter > 3 {
			os.Exit(1)
		} else {
			fmt.Printf("signal: %s %d\n", s, counter)
		}
	}
}

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go register(c)
	for {
		fmt.Println("waiting 5s")
		time.Sleep(5000 * time.Millisecond)
	}
}
