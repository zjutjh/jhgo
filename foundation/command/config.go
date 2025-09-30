package command

var DefaultConfig = Config{
	Logger:      "",
	Output:      true,
	PprofSwitch: false,
	PprofOutput: "./",
	PprofType:   []string{"cpu", "heap"},
}

type Config struct {
	Logger      string   `mapstructure:"logger"`
	Output      bool     `mapstructure:"output"`
	PprofSwitch bool     `mapstructure:"pprof_switch"`
	PprofOutput string   `mapstructure:"pprof_output"`
	PprofType   []string `mapstructure:"pprof_type"`
}
