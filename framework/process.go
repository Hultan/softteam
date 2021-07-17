package framework

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"syscall"
)

type Process struct {
}

// OpenInTerminal : Opens a new terminal with path as its current directory
func (p *Process) OpenInTerminal(path string) {
	p.Open("gnome-terminal", "--working-directory="+path)
}

// OpenInNemo : Opens a new nemo instance with the specified folder opened
func (p *Process) OpenInNemo(path string) {
	p.Open("nemo", path)
}

// OpenInGoland : Opens the specified folder in Goland
func (p *Process) OpenInGoland(path string) {
	p.Open("goland", path)
}

// Open : Opens a command with the specified arguments
func (p *Process) Open(command string, args ...string) {

	cmd := exec.Command(command, args...)
	// Forces the new process to detach from the GitDiscover process
	// so that it does not die when GitDiscover dies
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	// set var to get the output
	var out bytes.Buffer

	// set the output to our variable
	cmd.Stdout = &out
	err := cmd.Start()
	if err != nil {
		log.Println(err)
	}
	cmd.Process.Release()

	fmt.Println(out.String())
}
