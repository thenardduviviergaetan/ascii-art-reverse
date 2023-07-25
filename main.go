package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var tabascii map[rune][]string

const Police_Ascii = 8

func main() {
	args := os.Args[1:]
	flag := ""
	file := ""

	if len(args) != 0 && len(args[0]) >= 10 {
		flag = args[0][:10]
		file = args[0][10:]
	} else {
		fmt.Println("Usage: go run . [OPTION]")
		fmt.Println()
		fmt.Println("EX: go run . --reverse=<fileName>")
		return
	}

	if flag != "--reverse=" {
		fmt.Println("Usage: go run . [OPTION]")
		fmt.Println()
		fmt.Println("EX: go run . --reverse=<fileName>")
		return
	}

	// initialise le tableau map ascii qui stocke l'ascii par carractere
	InitAscii()
	ReadFile(file)
	Research()
}

func InitAscii() {
	// tableaux des 3 version ascii possible
	// assci := [2]string{"standard.txt", "shadow.txt" /*, "thinkertoy.txt"*/}
	// lis le fichier ascii selectionner
	data, err := ioutil.ReadFile("data/standard.txt")
	CheckError(err)
	tabascii = make(map[rune][]string)
	var tabcont []string
	str := ""
	compt := 0
	index := ' '
	// découpe le fichier lu précédament pour le stocker
	for _, letter := range string(data) {
		if letter == '\n' && str != "" {
			compt++
			tabcont = append(tabcont, str)
			str = ""
		} else if letter != '\n' {
			str += string(letter)
		}
		if compt == Police_Ascii {
			tabascii[index] = tabcont
			var newtabcont []string
			tabcont = newtabcont
			index++
			compt = 0
		}
	}
}

func CheckError(err error) {
	if err != nil {
		// fmt.Println(err)
		log.Fatal(err)
	}
}

var res [][]string

func ReadFile(fille string) {
	data, err := os.ReadFile(fille)
	CheckError(err)
	data = []byte(strings.ReplaceAll(string(data), "\r\n", "\n"))
	tab := strings.Split(string(data), "\n")
	var asciline []string
	compt := 0
	for _, line := range tab {
		if line != "" && compt != 8 {
			asciline = append(asciline, line)
			compt++
		}
		if compt == 8 {
			res = append(res, asciline)
			asciline = []string{}
			compt = 0
		} else if line == "" {
			res = append(res, []string{""})
		}
	}
	// Testaffichage()
}

func Testaffichage() {
	for _, tabline := range res {
		for _, ligne := range tabline {
			fmt.Println(ligne, "   ", len(ligne))
		}
	}
}

func Clear() {
	time.Sleep(time.Millisecond * 10)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Affichage(tabline []string) {
	Clear()
	for _, ligne := range tabline {
		fmt.Println(ligne) //, "   ", len(ligne))
	}
}

func Research() {
	revers := ""
	var next func([]string, []string) bool
	next = func(message, temp []string) bool {
		for i := ' '; i <= '~'; i++ {
			verif := true
			for index, line := range tabascii[i] {
				strtemp := temp[index] + line
				if len(message[index]) < len(strtemp) {
					verif = false
				} else {
					tabrune := []rune(message[index])
					mes := ""
					for i := 0; i < len(strtemp); i++ {
						mes += string(tabrune[i])
					}
					if strtemp != mes {
						verif = false
					}
				}
			}
			if verif {
				revers += string(i)
				var newtemp []string
				for index, line := range tabascii[i] {
					newtemp = append(newtemp, string(temp[index]+line))
				}
				// Affichage(newtemp)
				if len(temp[0]) != len(message[0]) {
					end := next(message, newtemp)
					if end {
						return true
					}
				}
			}
		}
		return false
	}
	for index, message := range res {
		if len(message) == 1 && index != len(res)-1 {
			revers += "\\n"
		} else {
			if index < len(res)-1 {
				next(message, []string{"", "", "", "", "", "", "", ""})
			}
			if index < len(res)-2 {
				revers += "\\n"
			}
		}
	}
	fmt.Println(revers)
}
