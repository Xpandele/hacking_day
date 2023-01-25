package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func temp() {
	timer := time.NewTimer(time.Second * 5)
	<-timer.C
	fmt.Println("\nBOOOOOOOOOOOOOOOOOOM!")
	shutdown()
}

func main() {
	go temp()
	const key = "1"
	nIntents := 3
	explota := true

	sc := bufio.NewReader(os.Stdin)

	for ;nIntents > 0; nIntents-- {
		fmt.Print("Introdueix la contrasenya: ")
		text, err := sc.ReadString('\n')
		if err != nil {
			log.Fatal("error reading string: ", err)
		}

		text = strings.Trim(text, "\n\t")
		if text == key {
			explota = false
			break
		}

		fmt.Println("Contrasenya incorrecte")
		fmt.Printf("Queden %d intents\n", nIntents)
	}

	if explota {
		shutdown()
	}
}

func shutdown() {
	//command := exec.Command("shutdown", "now")

	//os := runtime.GOOS
	//if os == "windows" {
		//command = exec.Command("shutodwn", "/s")
	//}

	//command.Run()
	fmt.Println("Apaga ordinador")
}
