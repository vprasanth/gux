package cmds

import (
	"fmt"
	"github.com/Workiva/go-datastructures/queue"
	"github.com/vprasanth/gux/spec"
	"log"
	"os/exec"
)

// @todo fix splitType
func addPane(
	sessionName string,
	window spec.Window,
	windowIndex int,
	paneIndex int,
	pane spec.Pane,
	splitType string) *exec.Cmd {
	tmux, _ := exec.LookPath("tmux")
	target := fmt.Sprintf("%s:%d.%d", sessionName, windowIndex, paneIndex)
	args := []string{"split-window", "-d", splitType, "-t", target, "-c", window.WorkingDir, pane.Command}
	cmd := exec.Command(tmux, args...)
	return cmd
}

func killPane(sessionName string, windowIndex int, paneIndex int) *exec.Cmd {
	tmux, _ := exec.LookPath("tmux")
	target := fmt.Sprintf("%s:%d.%d", sessionName, windowIndex, paneIndex)
	args := []string{"kill-pane", "-t", target}
	cmd := exec.Command(tmux, args...)
	return cmd
}

func createSession(sessionName string) *exec.Cmd {
	// @todo check for tmux
	tmux, _ := exec.LookPath("tmux")
	args := []string{"new-session", "-d", "-s", sessionName}
	return exec.Command(tmux, args...)
}

func createWindow(window spec.Window) *exec.Cmd {
	tmux, _ := exec.LookPath("tmux")
	args := []string{"new-window", "-d", "-n", window.Name, "-c", window.WorkingDir}
	return exec.Command(tmux, args...)
}

func attach() {}

func Init(config spec.GuxConfig) {
	commandQueue := queue.New(10)
	for i := 0; i < len(config.Session); i++ {
		commandQueue.Put(createSession(config.Session[i].Name))
		for y := 0; y < len(config.Session[i].Window); y++ {
			commandQueue.Put(createWindow(config.Session[i].Window[y]))
			for z := 0; z < len(config.Session[i].Window[y].Panes); z++ {
				commandQueue.Put(addPane(
					config.Session[i].Name,
					config.Session[i].Window[y],
					y,
					z,
					config.Session[i].Window[i].Panes[z],
					"-v"))
			}
			killPane(config.Session[i].Name, y, 0)
		}
	}

	disposed := commandQueue.Dispose()
	run(disposed)
}

func run(commands []interface{}) {
	for i := 0; i < len(commands); i++ {
		fmt.Printf("cmd: %+v\n", commands[i])
		cmd := commands[i].(*exec.Cmd)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("I suck!: %s", out)
			log.Fatal(err)
		}
	}
}
