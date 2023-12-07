package logger

type Config struct {
	Level   string `yaml:"log_level"`
	JSON    *bool  `yaml:"log_json"`
	NoColor bool   `yaml:"log_no_color"`
}
