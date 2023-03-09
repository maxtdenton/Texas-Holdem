package lib

import (
	"fmt"
	"strings"
	"os"
	"os/exec"
	"time"
)

func Split(str string) (string, string) {
	var str1, str2 string
	s := strings.Split(str, ",")
	str1 = Up(s[0])
	str2 = Up(s[1])
	return str1, str2
}

func Up(word string) string{
	upcaseWord := strings.ToUpper(word)
	return upcaseWord
}

func ClearTerm() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	fmt.Println("Clearing Screen...")
	duration, _ := time.ParseDuration("1s")
	time.Sleep(duration)
	clear.Run()
}