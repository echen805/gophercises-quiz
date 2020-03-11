package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	quiz := make(map[string]int)
	pwd, _ := os.Getwd()
	f, _ := os.Open(filepath.Join(pwd,"problems.csv"))
	r := csv.NewReader(bufio.NewReader(f))
	fmt.Println("Preparing the quiz, get ready!")
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		i, err := strconv.Atoi(record[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		quiz[record[0]] = i
	}
	fmt.Println("It's quiz time! Input your answer and press Enter when you are ready. " +
		"Note: Once you press Enter, you cannot go back to a previous question.")
	reader := bufio.NewReader(os.Stdin)
	correct := 0
	for prob, ans := range quiz {
		fmt.Println(prob)
		text, _,_ := reader.ReadLine()
		userInput, _ := strconv.Atoi(strings.TrimSpace(string(text)))
		if ans == userInput {
			correct++
		}
	}
	fmt.Printf("There were %d questions and you got %d correct.", len(quiz), correct)
}
