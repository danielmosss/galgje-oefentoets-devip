package main

import (
	"bufio"
	"os"
	"strings"
)

func ScanLetter() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSuffix(strings.TrimSuffix(text, "\n"), "\r")[:1]
}
