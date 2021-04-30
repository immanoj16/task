package task

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/immanoj16/task/pkg/errors"
)

const defaultTaskFile = `# https://taskfile.dev

version: '3'

vars:
	GREETING: Hello, World!

tasks:
	default:
		cmds:
			- echo "{{.GREETING}}"
		silent: true
`

// TaskFile creates a new taskfile
func TaskFile(w io.Writer, dir string) error {
	f := filepath.Join(dir, "Taskfile.yml")

	if _, err := os.Stat(f); err == nil {
		return errors.ErrTaskfileAlreadyExists
	}

	if err := ioutil.WriteFile(f, []byte(defaultTaskFile), 0644); err != nil {
		return err
	}

	fmt.Fprintf(w, "Taskfile.yml created in the current directory\n")
	return nil
}
