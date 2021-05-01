package taskfile

import "strings"

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

// UnmarshalYAML implements yaml.Unmarshaller interface
func (c *Cmd) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var cmd string
	if err := unmarshal(&cmd); err == nil {
		if strings.HasPrefix(cmd, "^") {
			c.Task = strings.TrimPrefix(cmd, "^")
		} else {
			c.Command = cmd
		}
		return nil
	}

	var cmdStruct struct {
		Command     string
		Silent      bool
		IgnoreError bool `yaml:"ignore_error"`
	}

	if err := unmarshal(&cmdStruct); err == nil && cmdStruct.Command != "" {
		c.Command = cmdStruct.Command
		c.Silent = cmdStruct.Silent
		c.IgnoreError = cmdStruct.IgnoreError
		return nil
	}

	var taskCall struct {
		Task string
		Vars *Vars
	}
	if err := unmarshal(&taskCall); err != nil {
		return err
	}
	c.Task = taskCall.Task
	c.Vars = taskCall.Vars

	return nil
}

// UnmarshalYAML implements yaml.Unmarshaller interface
func (d *Dep) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var task string
	if err := unmarshal(&task); err == nil {
		d.Task = task
		return nil
	}
	var taskCall struct {
		Task string
		Vars *Vars
	}
	if err := unmarshal(&taskCall); err != nil {
		return err
	}
	d.Task = taskCall.Task
	d.Vars = taskCall.Vars
	return nil
}
