package main

import (
	"bufio"
	"embed"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

//go:embed ascii_art/*
var gameOverFile embed.FS

func temp() {
	timer := time.NewTimer(time.Second * 5)
	<-timer.C
	shutdown()
}

func main() {
	go temp()
	const key = "K GRANDE"
	explota := true

	sc := bufio.NewReader(os.Stdin)

	for nIntents := 3; nIntents > 0; nIntents-- {
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
	} else {
		win()
	}
}

func shutdown() {
	var fileData, err = gameOverFile.ReadFile("ascii_art/game_over.txt")
	if err != nil {
		log.Fatal("Error getting file")
	}
	//command := exec.Command("shutdown", "now")

	//os := runtime.GOOS
	//if os == "windows" {
	//command = exec.Command("shutodwn", "/s")
	//}

	//command.Run()
	fmt.Printf("\n%s\n", string(fileData))
}

func win() {
	var fileData, err = gameOverFile.ReadFile("ascii_art/win.txt")
	if err != nil {
		log.Fatal("Error getting file")
	}

	fmt.Printf("\n%s\n", string(fileData))
}
