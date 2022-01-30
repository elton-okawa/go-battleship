package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("> Type a command: ")
	for scanner.Scan() {
		input := scanner.Text()
		if len(input) <= 0 {
			fmt.Println("> You must type a command")
		} else if input == "quit" || input == "exit" {
			fmt.Println("> Bye")
			break
		} else {
			splitted := strings.Split(input, " ")

			cmd := splitted[0]
			args := splitted[1:]

			executeCommand(cmd, args)
		}

		fmt.Printf("\n> Type a command: ")
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func executeCommand(cmd string, args []string) bool {
	if _, exist := Commands[cmd]; exist {
		if err := Commands[cmd].Parse(args); err == nil {
			end, err := Commands[cmd].Execute()
			if err != nil {
				fmt.Println(err)
			}

			return end
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("> Command '%s' not found\n", cmd)
	}

	return false
}
