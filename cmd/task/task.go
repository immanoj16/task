package main

import (
	"fmt"
	"log"
	"os"

	"github.com/immanoj16/task/pkg/task"
	"github.com/spf13/pflag"
)

var (
	version = ""
)

const usage = `Usage: task [-ih] [--init] [--version] [--help]
Options:
`

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stderr)

	pflag.Usage = func() {
		log.Print(usage)
		pflag.PrintDefaults()
	}

	var (
		versionFlag bool
		helpFlag    bool
		init        bool
	)

	pflag.BoolVar(&versionFlag, "version", false, "show Task version")
	pflag.BoolVarP(&helpFlag, "help", "h", false, "show Task Usage")
	pflag.BoolVarP(&init, "init", "i", false, "creates a new Taskfile.yml in the current folder")
	pflag.Parse()

	if versionFlag {
		fmt.Printf("Task version: %s\n", getVersion())
		return
	}

	if helpFlag {
		pflag.Usage()
		return
	}

	if init {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		if err := task.TaskFile(os.Stdout, wd); err != nil {
			log.Fatal(err)
		}
		return
	}
}

func getVersion() string {
	return "2.3.3"
}
