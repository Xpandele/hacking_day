package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func temp() {
	timer := time.NewTimer(time.Second * 5)
	<-timer.C
	fmt.Println("BOOOOOOOOOOOOOOOOOOM!")
	os.Exit(0)
}

func main() {
	go temp()
	const key = "1"
	nIntents := 3
	sc := bufio.NewReader(os.Stdin)

	for i := 0; i < nIntents; i++ {
		text, err := sc.ReadString('\n')
		if err != nil {
			log.Fatal("error reading string: ", err)
		}

		if text != key {
			fmt.Println("Contrasenya incorrecte")
			fmt.Printf("Queden %d intents\n", i)
			continue
		}

		break
	}
}
