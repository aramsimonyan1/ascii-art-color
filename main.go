package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// read standard.txt and convert to array of lines
	readFile, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string
	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	if len(os.Args) < 3 || len(os.Args) > 4 {
		printUsage()
		os.Exit(1)
	}

	colorFlag := os.Args[1]
	if !strings.HasPrefix(colorFlag, "--color=") {
		fmt.Println("Error: Invalid format for --color flag. Please use --color=<color>")
		printUsage()
		os.Exit(1)
	}
	colorFlag = strings.TrimPrefix(colorFlag, "--color=")

	lettersToColor := os.Args[2]
	text := ""

	if len(os.Args) == 4 {
		text = os.Args[3]
	} else {
		text = lettersToColor
	}

	// works with:     --color=red "hello world" "hello world"    or    --color=red "hello world"    or    --color=red "world"
	if lettersToColor == text {
		fmt.Println("processOneString")
		fmt.Println(lettersToColor)
		fmt.Println(text)
		processOneString(text, lettersToColor, colorFlag, fileLines)
	}

	// works with:     --color=orange GuYs "HeY GuYs"      but not    --color=red hello "hello world"    and not with 3 words 'text'
	textSlice := strings.Split(text, " ")
	var matchingWord string
	for i := 0; i < len(textSlice); i++ {
		if lettersToColor == textSlice[i] {
			matchingWord := textSlice[i]
			lenOfPrevWord := len(textSlice[i-1])
			fmt.Println("processMatchingWord")
			fmt.Println("lettersToColor: ", lettersToColor)
			fmt.Println("matchingWord: ", matchingWord)
			fmt.Println("text: ", text)
			processMatchingWord(text, lettersToColor, colorFlag, fileLines, lenOfPrevWord)
		}
	}

	/* work with different variables i.e.:
	   Try specifying set of letters to be colored (the second until the last letter). --color=blue ram "Aram"
	   Try specifying letter to be colored (the second letter).             --color=blue r "Aram"
	   Try specifying a set of letters to be colored (just two letters).	--color=blue rm "Aram"  */
	if lettersToColor != text && lettersToColor != matchingWord {
		fmt.Println("processNotEqualVariables")
		fmt.Println(lettersToColor)
		fmt.Println(text)
		processNotEqualVariables(text, lettersToColor, colorFlag, fileLines)
	}
}

// to handle single string input
func processOneString(text string, lettersToColor string, colorFlag string, fileLines []string) {
	word := []rune(text)
	for k := 1; k < 9; k++ {
		for i := 0; i < len(word); i++ {
			asciiFetch := ((word[i] - 32) * 9) + rune(k)
			fmt.Printf("%s", colorize(fileLines[asciiFetch], colorFlag))
		}
		fmt.Println()
	}
}

// to handle matching word in 'lettersToColor' with word in 'text'
func processMatchingWord(text string, lettersToColor string, colorFlag string, fileLines []string, lenOfPrevWord int) {
	textSlice := []rune(text)
	for j := 1; j < 9; j++ {
		for k := 0; k < len(textSlice); k++ {
			asciiFetch := ((textSlice[k] - 32) * 9) + rune(j)
			letters := lenOfPrevWord + 1
			if k == letters || (k >= lenOfPrevWord && k <= letters+lenOfPrevWord+1) {
				fmt.Printf("%s", colorize(fileLines[asciiFetch], colorFlag))
				letters++
			} else {
				fmt.Print(fileLines[asciiFetch])
			}
		}
		fmt.Println()
	}
}

// to match the letters in 'lettersToColor' with letters in 'text'
func processNotEqualVariables(text string, lettersToColor string, colorFlag string, fileLines []string) {
	word := []rune(text)
	for j := 1; j < 9; j++ {
		for k := 0; k < len(word); k++ {
			asciiFetch := ((word[k] - 32) * 9) + rune(j)
			if strings.ContainsRune(lettersToColor, word[k]) {
				fmt.Printf("%s", colorize(fileLines[asciiFetch], colorFlag))
			} else {
				fmt.Print(fileLines[asciiFetch])
			}
		}
		fmt.Println()
	}
}

func colorize(text string, colorFlag string) string {
	colorMapping := map[string]string{
		"black":   "\033[30m%s\033[0m",
		"red":     "\033[31m%s\033[0m",
		"green":   "\033[32m%s\033[0m",
		"yellow":  "\033[33m%s\033[0m",
		"blue":    "\033[34m%s\033[0m",
		"purple":  "\033[35m%s\033[0m",
		"magenta": "\033[35m%s\033[0m",
		"cyan":    "\033[36m%s\033[0m",
		"white":   "\033[37m%s\033[0m",
		"orange":  "\033[38;5;208m%s\033[0m",
		"gray":    "\033[90m%s\033[0m",
	}

	format, found := colorMapping[colorFlag]
	if !found {
		return text
	}
	return fmt.Sprintf(format, text)
}

func printUsage() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println("EX: go run . --color=<color> <letters to be colored> 'something'")
}
