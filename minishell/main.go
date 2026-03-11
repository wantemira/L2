package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Ctrl+C
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		for range sigChan {
			fmt.Println()
		}
	}()

	for {
		fmt.Print("minishell> ")

		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("\nexit")
			return
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		runCommand(line)
	}
}

func runCommand(line string) {
	pipes := strings.Split(line, "|")

	if len(pipes) == 1 {
		args := strings.Fields(pipes[0])

		if runBuiltin(args) {
			return
		}
	}

	runPipeline(pipes)
}

func runBuiltin(args []string) bool {

	switch args[0] {

	case "cd":
		if len(args) < 2 {
			fmt.Println("cd: missing path")
			return true
		}
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Println("cd error:", err)
		}
		return true

	case "pwd":
		dir, _ := os.Getwd()
		fmt.Println(dir)
		return true

	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
		return true

	case "kill":
		if len(args) < 2 {
			fmt.Println("kill: missing pid")
			return true
		}

		pid, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("invalid pid")
			return true
		}

		proc, err := os.FindProcess(pid)
		if err == nil {
			proc.Signal(syscall.SIGKILL)
		}

		return true

	case "ps":
		cmd := exec.Command("ps")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
		return true
	}

	return false
}

func runPipeline(commands []string) {

	var cmds []*exec.Cmd

	for _, c := range commands {
		args := strings.Fields(strings.TrimSpace(c))
		cmds = append(cmds, exec.Command(args[0], args[1:]...))
	}

	for i := 0; i < len(cmds)-1; i++ {
		stdout, _ := cmds[i].StdoutPipe()
		cmds[i+1].Stdin = stdout
	}

	cmds[len(cmds)-1].Stdout = os.Stdout
	cmds[len(cmds)-1].Stderr = os.Stderr

	for _, cmd := range cmds {
		cmd.Start()
	}

	for _, cmd := range cmds {
		cmd.Wait()
	}
}
