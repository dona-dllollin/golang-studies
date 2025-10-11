package main

import (
	"fmt"
	"os/exec"
)

func main() {
	var output1, _ = exec.Command("ls").Output()
	fmt.Printf(" -> ls\n%s\n", string(output1))
	var output2, _ = exec.Command("pwd").Output()
	fmt.Printf(" -> pwd\n%s\n", string(output2))

}
