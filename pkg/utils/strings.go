package utils

import (
	"bufio"
	"log"
	"strings"
)

// This function gets user input in the terminal
func GetUserInput(reader *bufio.Reader) string {
	val, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Malformed input. Please start over.")
	}
	// Because ReadString returns a string with the delimiter attached,
	// we need to remove the newline character from the end of the input
	// and below covers Unix-like and Windows
	val = strings.TrimRight(val, "\r\n")
	return val
}
