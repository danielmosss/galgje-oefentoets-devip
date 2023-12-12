package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	antwoord := strings.ToLower(RandomWord())
	voortgang := ""
	fouteLetters := ""
	aantalFouten := 9

	fmt.Println("Welkom bij galgje!")
	fmt.Println("Antwoord: " + antwoord)
	voortgang = strings.Repeat("_", len(antwoord))
	var letter string
	for true {
		letter = printGalgje(voortgang, fouteLetters, aantalFouten)
		if len(letter) > 1 {
			letter = string(letter[0])
		}

		if strings.Contains(antwoord, letter) {
			for i := 0; i < len(antwoord); i++ {
				if string(antwoord[i]) == letter {
					voortgang = ReplaceLetter(voortgang, i, letter)
				}
			}
		} else {
			fouteLetters += letter
			aantalFouten--
		}

		if voortgang == antwoord {
			fmt.Println("--------------------")
			fmt.Println("Het woord was " + antwoord + ".")
			fmt.Println("Lekker man, je hebt gewonnen.")
			fmt.Println("--------------------")
		}

		if aantalFouten == 0 {
			fmt.Println("--------------------")
			fmt.Println("JE BEN ZO EXTREEM NOOB LOZER. IMAGINE TRYING TO PLAY THIS GAME AND LOSE.")
		}

		if aantalFouten == 0 || voortgang == antwoord {
			fmt.Println("Wil je nog een keer spelen? (y/n)")
			antwoord = strings.ToLower(ReadLn())
			if antwoord == "y" {
				antwoord = strings.ToLower(RandomWord())
				voortgang = ""
				fouteLetters = ""
				aantalFouten = 9
				fmt.Println("Antwoord: " + antwoord)
				voortgang = strings.Repeat("_", len(antwoord))
			} else {
				fmt.Println("Ok, doei.")
				break
			}
		}
	}

}

func printGalgje(voortgang string, fouteLetters string, aantalFouten int) string {
	fmt.Println("--------------------")
	fmt.Println("Je woord is " + voortgang)
	fmt.Println("Je hebt nog " + fmt.Sprint(aantalFouten) + " fouten over")
	fmt.Println("Je hebt de letters [" + fouteLetters + "] al geprobeerd")
	fmt.Print("Geef een letter: ")
	volgendeLetter := ReadLn()

	return volgendeLetter
}

func ReplaceLetter(input string, index int, replace string) string {
	return input[:index] + replace + input[index+1:]
}

func RandomWord() string {
	file, err := os.Open("wordlist.txt")
	if err != nil {
		fmt.Errorf("%s\r\n", err)
		return ""
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines[rand.Intn(len(lines))]
}

func ReadLn() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\r\n")
	return text
}
