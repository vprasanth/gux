package cmds

import (
	"fmt"
	"github.com/Workiva/go-datastructures/queue"
	"github.com/vprasanth/gux/spec"
	"log"
	_ "os"
	"os/exec"
	_ "syscall"
)

func Start(name string) {
	tmux, err := exec.LookPath("tmux")

	if err != nil {
		log.Fatal(err)
	}

	args := []string{"new-session", "-d", "-s" + name}
	//err = syscall.Exec(tmux, args, os.Environ())
	cmd := exec.Command(tmux, args...)
	err = cmd.Run()

	if err != nil {
		fmt.Printf("Listen, I tried okay. But I couldn't do it, I couldn't make that '%s' session you asked for.\n", name)
		log.Fatal(err)
	}
}

func CreateVerticalSplitLayout(sessionName string, window spec.Window) {
	type command struct {
		bin  string
		args []string
	}

	commands := queue.New(10)

	commands.Put(command{"tmux", []string{"new-session", "-d", "-s", window.Name, "-c", window.WorkingDir}})

	for i := 0; i < len(window.Panes); i++ {
		var args []string
		if i > 0 {
			args = []string{
				"split-window",
				"-v",
				"-c",
				window.WorkingDir,
				"-t",
				fmt.Sprintf("%s:0.%d", window.Name, i-1),
				window.Panes[i].Command}
		} else {
			args = []string{
				"respawn-pane",
				"-c",
				window.WorkingDir,
				"-k",
				"-t",
				fmt.Sprintf("%s:0.%d", window.Name, i),
				window.Panes[i].Command}
		}

		commands.Put(command{"tmux", args})
	}

	commandsToRun := commands.Dispose()

	for i := 0; i < len(commandsToRun); i++ {
		ugly := commandsToRun[i].(command)
		fmt.Printf("Do this: %+v\n", ugly)
		out, err := exec.Command(ugly.bin, ugly.args...).CombinedOutput()
		if err != nil {
			fmt.Printf("I suck: %s\n", out)
			fmt.Printf("Shit I tried to process now (%d): %+v\n", i, ugly)
			log.Fatal(err)
		}
	}
}
