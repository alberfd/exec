package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

func main() {
	for {
		//llamar a tu script
		go execCommand("sleep", []string{"5"})
		time.Sleep(5 * time.Second)
	}
}

func execCommand(command string, params []string) {
	cmd := exec.Command(command, params...)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(out.String())
}
