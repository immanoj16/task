package task

import (
	"io"
	"sync"

	"github.com/immanoj16/task/pkg/taskfile"
)

// Executer executes a taskfile
type Executer struct {
	TaskFile *taskfile.TaskFile

	Dir         string
	Entrypoint  string
	Force       bool
	Watch       bool
	Verbose     bool
	Silent      bool
	Dry         bool
	Summary     bool
	Parallel    bool
	Color       bool
	Concurrency bool

	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer

	Logger      interface{}
	Compiler    interface{}
	Output      interface{}
	OutputStyle string

	taskVars interface{}

	concurrencySemaphone chan struct{}
	taskCallCount        map[string]*int32
	mkdirMutexMap        map[string]*sync.Mutex
}
