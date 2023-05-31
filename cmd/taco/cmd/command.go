package cmd

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func runAndWatchCommand(command *exec.Cmd) {
	stdout, err := command.StdoutPipe()
	if err != nil {
		log.Panic(err)
	}

	stderr, err := command.StderrPipe()
	if err != nil {
		log.Panic(err)
	}

	err = command.Start()
	if err != nil {
		log.Panic(err)
	}

	scanner := bufio.NewScanner(io.MultiReader(stdout, stderr))
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}

	err = command.Wait()
	if err != nil {
		log.Panic(err)
	}
}
