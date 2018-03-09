package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("mdsh -c \"RODDP:DEST=ALL;\"")
	cmd.Stdin = strings.NewReader("some input")

	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("in all caps: %q\n", out.String())
}
