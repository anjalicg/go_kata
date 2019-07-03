package main

//https://tutorialedge.net/golang/reading-console-input-golang/

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
		fmt.Printf("type=%T\n", text)
	}

}
func method2() {
	// read one unicode from command line
	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(char)
	fmt.Printf("Type = %T\n", char)

	switch char {
	case 'A':
		fmt.Println("A pressed")
	case 'a':
		fmt.Println("a pressed")
	default:
		fmt.Println("Neither A nor a pressed")
	}

}
func method3() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		x := scanner.Text()
		fmt.Println("User entered:-", x)
		fmt.Printf("type=%T\n", x)
	}

}
func main() {
	// method1()
	method2()
	// method3()

}
