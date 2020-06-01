package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	OptionHelp = "help"
	OptionHelpShort = "?"
	OptionExit = "exit"
)

func main() {
	var exit bool
	var reader *bufio.Reader
	reader = bufio.NewReader(os.Stdin)
	showCredits()
	for ok := true; ok; ok = !exit {
		showMenu()
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		fmt.Println("----------------------------------------------------------------------")
		switch text {
		case OptionHelp, OptionHelpShort:
			fmt.Println("[HELP]\tA basic orchestrator program")
			fmt.Println("\thelp\tGet the list of commands")
			fmt.Println("\texit\tQuit the program")
			break
		case OptionExit:
			fmt.Println("[EXIT] Leaving the current session...")
			exit = true
		default:
			fmt.Println("Invalid option")
		}

	}
}

func showMenu() {
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println(fmt.Sprintf("Use [%s] or [%s] to receive more information.", OptionHelp, OptionHelpShort))
}

func showCredits() {
	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("\tGo distributed - A basic orchestrator program")
	fmt.Println("\tCreated by Marcos Huck")
}