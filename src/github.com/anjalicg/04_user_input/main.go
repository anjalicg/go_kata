package main

import (
	"bufio"
	"fmt"
	"os"
)

func method1() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome")
	for {
		fmt.Print("user>")
		text, _ := reader.ReadString('\n')
		fmt.Println("Hello", text)
	}

}
func method2() {
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(char)

	switch char {
	case 'A':
		fmt.Println("A pressed")
	case 'a':
		fmt.Println("a pressed")
	}

}
func method3() {

}
func main() {
	method1()
	// method2()
	// method3()

}
