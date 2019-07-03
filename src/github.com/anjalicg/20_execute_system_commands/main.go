package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("date")
	fmt.Println(cmd.Path)
	fmt.Println(cmd.Args)
	stdout, _ := cmd.Output()
	fmt.Println(string(stdout))

	cmd = exec.Command("ifconfig", "en0")
	fmt.Println(cmd.Path)
	fmt.Println(cmd.Args)
	stdout1, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(stdout1))
}
