package taskfile

// Cmd is a task command
type Cmd struct {
	Command     string
	Silent      bool
	Task        string
	Vars        *Vars
	IgnoreError bool
}

// Dep is a task dependency
type Dep struct {
	Task string
	Vars *Vars
}
