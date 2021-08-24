package framework

import (
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
func (p *Process) Open(command string, args ...string) string {

	cmd := exec.Command(command, args...)
	// Forces the new process to detach from the GitDiscover process
	// so that it does not die when GitDiscover dies
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	cmd.Process.Release()

	return string(output)
}
