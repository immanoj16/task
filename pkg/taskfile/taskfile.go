package taskfile

// TaskFile represents a Taskfile.yml
type TaskFile struct {
	Version    string
	Expansions int
	Output     string
	Method     string
	Includes   *IncludedTaskfiles
	Vars       *Vars
	Env        *Vars
	Tasks      Tasks
	Silent     bool
	Dotenv     []string
}

// UnmarshalYAML implements yaml.Unmarshaler interface
func (tf *TaskFile) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var taskFile struct {
		Version    string
		Expansions int
		Output     string
		Method     string
		Includes   *IncludedTaskfiles
		Vars       *Vars
		Env        *Vars
		Tasks      Tasks
		Silent     bool
		Dotenv     []string
	}
	if err := unmarshal(&taskFile); err != nil {
		return err
	}
	tf.Version = taskFile.Version
	tf.Expansions = taskFile.Expansions
	tf.Output = taskFile.Output
	tf.Method = taskFile.Method
	tf.Includes = taskFile.Includes
	tf.Vars = taskFile.Vars
	tf.Env = taskFile.Env
	tf.Tasks = taskFile.Tasks
	tf.Silent = taskFile.Silent
	tf.Dotenv = taskFile.Dotenv
	if tf.Expansions <= 0 {
		tf.Expansions = 2
	}
	if tf.Vars == nil {
		tf.Vars = &Vars{}
	}
	if tf.Env == nil {
		tf.Env = &Vars{}
	}
	return nil
}
